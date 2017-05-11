package testsimpleapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
)

type TestSimpleApi struct {
	Url     string            //url router
	Func    rest.HandlerFunc  //handler function for api
	Method  string            //GET,POST,PUT,DELETE
	ReqUrl  string            //url client request
	Payload interface{}       //Payload of request
	Header  map[string]string //Header of request
}

func (ts *TestSimpleApi) RunRequest(t *testing.T) *test.Recorded {
	app := rest.NewApi()
	app.Use(rest.DefaultDevStack...)
	routes := []*rest.Route{}
	routes = append(routes,
		&rest.Route{
			HttpMethod: ts.Method,
			PathExp:    ts.Url,
			Func:       ts.Func,
		},
	)
	router, err := rest.MakeRouter(routes...)
	if err != nil {
		log.Fatal(err)
	}
	app.SetApp(router)

	recorded := test.RunRequest(t, app.MakeHandler(),
		makeCustomRequest(ts.Method, "http://1.2.3.4"+ts.ReqUrl, ts.Payload, ts.Header))
	return recorded
}

func makeCustomRequest(method string, urlStr string, payload interface{}, header map[string]string) *http.Request {
	var s string

	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			panic(err)
		}
		s = fmt.Sprintf("%s", b)
	}

	r, err := http.NewRequest(method, urlStr, strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	r.Header.Set("Accept-Encoding", "gzip")

	for k, v := range header {
		r.Header.Set(k, v)
	}
	if payload != nil {
		r.Header.Set("Content-Type", "application/json")
	}

	return r
}

package main

import (
	"testing"
	"testsimpleapi"
)

func TestCallIndex(t *testing.T) {
	routing := Routing{}
	testApi := testsimpleapi.TestSimpleApi{
		Url:    "/",
		Func:   routing.Index,
		Method: "GET",
		ReqUrl: "/",
		Header: map[string]string{},
	}
	recorded := testApi.RunRequest(t)

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`{
  "id": "1",
  "name": "Golang"
}`)
}

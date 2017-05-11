package main

import (
	"testing"
	"testsimpleapi"
)

func TestCallIndex(t *testing.T) {
	routing := Routing{}
	testApi := testsimpleapi.TestSimpleApi{
		Url:    "/",
		Func:   routing.GetStudentInfo,
		Method: "GET",
		ReqUrl: "/",
		Header: map[string]string{},
	}
	recorded := testApi.RunRequest(t)

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`{
  "studentList": [
    {
      "id": "6004101330",
      "name": "phatcharaphan",
      "lastName": "ananpreechakun",
      "basicEnglish1": "3.5",
      "calculus1": "4.0",
      "computerGraphics": "2.5",
      "softwareArchitecture": "1.0",
      "databaseSystem": "3.0"
    },
    {
      "id": "6004101332",
      "name": "angkarn",
      "lastName": "janjuang",
      "basicEnglish1": "4.0",
      "calculus1": "3.0",
      "computerGraphics": "3.5",
      "softwareArchitecture": "2.0",
      "databaseSystem": "3.5"
    }
  ]
}`)
}

package jsonutils

import (
	"testing"
	"bytes"
	"net/http"
	"net/http/httptest"
)


var jsonTests = []struct {
	name string
	json string
	maxSize int64
	allowUnknown bool
	errorExpected bool
}{
	{name:"good json", json: `{"field" : "value"}`, maxSize: 1024, allowUnknown: false, errorExpected: false},
	{name:"badly-formated json", json: `{"field" :}`, maxSize: 1024, allowUnknown: false, errorExpected: true},
	{name:"incorrect type", json: `{"field" : 1}`, maxSize: 1024, allowUnknown: false, errorExpected: true},
	{name:"two json files", json: `{"field" : "val1"}{"anothe_one": "val2"}`, maxSize: 1024, allowUnknown: false, errorExpected: true},
	{name: "empty body", json: ``, errorExpected: true, maxSize: 1024, allowUnknown: false},
	{name: "syntax error in json", json: `{"field": 1"`, errorExpected: true, maxSize: 1024, allowUnknown: false},
	{name: "unknown field in json", json: `{"fooo": "1"}`, errorExpected: true, maxSize: 1024, allowUnknown: false},
	{name: "allow unknown fields in json", json: `{"fooo": "1"}`, errorExpected: false, maxSize: 1024, allowUnknown: true},
	{name: "missing field name", json: `{jack: "1"}`, errorExpected: true, maxSize: 1024, allowUnknown: true},
	{name: "file too large", json: `{"field": "value"}`, errorExpected: true, maxSize: 5, allowUnknown: true},
	{name: "not json", json: `Hello, world!`, errorExpected: true, maxSize: 1024, allowUnknown: true},
}

func TestJsonUtils_ReadJSON(t *testing.T) {
	var testUtils JSONUtils
	for _, e := range jsonTests {
		testUtils.MaxSize = e.maxSize
		testUtils.AllowUnknownFields = e.allowUnknown

		var decodedJSON struct {
			Field string `json:"field"`
		}

		req, err := http.NewRequest("POST", "/", bytes.NewReader([]byte(e.json)))
		if err != nil {
			t.Log("Error:", err)
		}

		recorder := httptest.NewRecorder()

		err = testUtils.ReadJSON(recorder, req, &decodedJSON)

		if e.errorExpected && err == nil {
			t.Errorf("%s: error expected, but non received", e.name)
		}

		if !e.errorExpected && err != nil {
			t.Errorf("%s: error not expected, but one received: %s", e.name, err.Error())
		}

		req.Body.Close()
	}
}
package ezhttp

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestData struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	t := TestData{
		Field1: "hello1",
		Field2: "hello2",
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(t)
}

func TestJsonGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handlerFunc))
	defer ts.Close()

	c := NewClient()
	var data TestData
	status, err := c.JsonGet(ts.URL, &data)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}

	if status != 200 {
		t.Fatalf("JsonGet status should be 200, got %d\n", status)
	}

	if data.Field1 != "hello1" || data.Field2 != "hello2" {
		t.Fatalf("Incorrect decode, expected fields to be 'hello1' and 'hello2', got %s/%s",
			data.Field1, data.Field2)
	}
}

func TestJsonPost(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			decoder := json.NewDecoder(r.Body)
			var data TestData
			err := decoder.Decode(&data)
			if err != nil {
				t.Fatalf("Unable to decode post data\n")
			}

			if data.Field1 != "hello1" || data.Field2 != "hello2" {
				t.Fatalf("Incorrect decode, expected fields to be 'hello1' and 'hello2', got %s/%s\n",
					data.Field1, data.Field2)
			}

			encoder := json.NewEncoder(w)
			encoder.Encode(data)
		}))
	defer ts.Close()
	td := TestData{
		Field1: "hello1",
		Field2: "hello2",
	}

	var td2 TestData

	c := NewClient()
	status, err := c.Json(&td).JsonPost(ts.URL, &td2)
	if err != nil {
		t.Fatalf("JsonPost returned unexpected error %s\n", err.Error())
	}

	if status != 200 {
		t.Fatalf("Expected status == 200, got %d\n", status)
	}

	if td2.Field1 != "hello1" || td2.Field2 != "hello2" {
		t.Fatalf("Incorrect decode, expected fields to be 'hello1' and 'hello2', got %s/%s\n",
			td2.Field1, td2.Field2)
	}

	status, err = c.JsonStr(`{"field1": "hello1", "field2": "hello2"}`).JsonPost(ts.URL, &td2)
	if err != nil {
		t.Fatalf("JsonPost returned unexpected error %s\n", err.Error())
	}

	if status != 200 {
		t.Fatalf("Expected status == 200, got %d\n", status)
	}

	if td2.Field1 != "hello1" || td2.Field2 != "hello2" {
		t.Fatalf("Incorrect decode, expected fields to be 'hello1' and 'hello2', got %s/%s\n",
			td2.Field1, td2.Field2)
	}
}

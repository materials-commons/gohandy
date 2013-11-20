package gohandy

import (
	"encoding/json"
	"fmt"
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

	t.Field1 = "hello field1"

	b, _ := json.Marshal(&t)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, b)
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

	fmt.Printf("status = %d\n", status)
	fmt.Println(data)
}

package earthquake

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GoodResponse(t *testing.T) {
	c := NewClient()

	resp, err := c.Get("20190222")
	if err != nil {
		t.Error(err)
	}

	if len(resp) == 0 {
		t.Error("should return more than 0")
	}

	if resp[0].Profundidad != 20 {
		t.Errorf("should return profundidad: %d", 20)
	}
}

func TestClient_BadResponse(t *testing.T) {
	c := NewClient()

	resp, err := c.Get("20200222")
	if len(resp) > 0 {
		t.Error("Should return 0 records")
	}

	if err == nil {
		t.Error("Should return an error")
	}

	if err != errInvalidDate {
		t.Errorf("should return err: %s", errInvalidDate.Error())
	}
}

func TestHTTPClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL.String())
		// Send response to be tested
		// resp := []byte(`{"horaConsulta":"14:16","id":"NULL","descripcion":"Paradero no v\u00e1lido","servicios":[]}`)
		rw.WriteHeader(500)
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{
		client:  server.Client(),
		baseURL: server.URL,
	}
	_, err := api.Get("pa674")
	if err != nil {
		if err != errInternalServerError {
			t.Error("should return internal server error")
		}
	}
}

func TestHTTPClient_BadMessage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL.String())
		// Send response to be tested
		resp := []byte(`"horaConsulta":"14:16","id":"NULL","descripcion":"Paradero no v`)
		rw.Write(resp)
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{
		client:  server.Client(),
		baseURL: server.URL,
	}
	_, err := api.Get("pa674")
	if err == nil {
		t.Error("should return an error: 'invalid character'")
	}
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestHTTPClient_errorOnReading(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL.String())
		// Send response to be tested
		resp := []byte(`"horaConsulta":"14:16","id":"NULL","descripcion":"Paradero no v`)
		rw.Write(resp)
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := Client{
		client:  server.Client(),
		baseURL: server.URL,
	}
	testRequest := httptest.NewRequest(http.MethodPost, "/something", errReader(0))
	_, err := api.Do(testRequest)
	if err == nil {
		t.Error("should return an error: 'invalid character'")
	}
}

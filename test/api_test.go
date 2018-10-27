package test

import (
	"encoding/json"
	"github.com/eexe1/supermaid-go/model"
	"github.com/eexe1/supermaid-go/network/routes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateBooking(t *testing.T) {
	r := routes.Routes()
	ts := httptest.NewServer(r)
	defer ts.Close()

	body := "{ \"orderid\": \"SAC22000811281\", \"sid\":\"C7725168\", \"name\": \"TestName\", \"rd1\": \"A2\", \"createDate\": \"2010/12/10\", \"phone\": \"09876123123\", \"phone2\": \"123123907\", \"email\": \"123@ase.com\", \"rd3\": \"day\", \"rd4\": \"ATM\", \"address\": \"test.comalsd\", \"note\": \"test Note\" }"
	if resp, resultBody := testRequest(t, ts, "PUT", "/v1/booking", strings.NewReader(body)); resultBody != nil {
		var booking model.Booking
		err := json.Unmarshal(resultBody, &booking)
		if err != nil {
			t.Fatal(err)
		}
	} else {
		if resp.StatusCode != 201 {
			t.Fatalf("invalid status code")
		} else {
			t.Fatalf("unknown error")
		}
	}

}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, []byte) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, nil
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, nil
	}
	defer resp.Body.Close()

	return resp, respBody
}

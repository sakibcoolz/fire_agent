package apimodel

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Request struct {
	Url      string   `json:"url"`
	Body     any      `json:"body"`
	Method   string   `json:"method"`
	Header   []Header `json:"header"`
	Response []byte   `json:"response"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func New() *Request {
	return &Request{}
}

func Tostring(v any) string {
	str, _ := json.Marshal(v)
	return string(str)
}

func ToStruct(str []byte, v any) {
	err := json.Unmarshal(str, v)
	if err != nil {
		log.Println(err)
	}
}

func (r *Request) SetRequest() {
	payload := strings.NewReader(Tostring(r.Body))

	client := &http.Client{}

	req, err := http.NewRequest(r.Method, r.Url, payload)
	if err != nil {
		log.Println(err)

		return
	}

	for _, value := range r.Header {
		req.Header.Add(value.Key, value.Value)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)

		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)

		return
	}

	r.Response = body
}

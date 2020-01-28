package modules

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type request struct {
}

func (m *Module) Request() {
	r := new(request)
	Baker.CreateModuleFunction("HTTP_GET", r.HttpGet)
	Baker.CreateModuleFunction("HTTP_POST", r.HttpPost)
	Baker.CreateModuleFunction("HTTP_REQUEST", r.HttpRequest)
}

func (r *request) HttpGet(url string, header map[string]interface{}) string {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("HTTP GET request error: %s\n", err)
		return ""
	}
	for k, v := range header {
		val, ok := v.(string)
		if ok {
			request.Header.Set(k, val)
		}
	}
	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		log.Printf("HTTP GET request error: %s\n", err)
		return ""
	}

	respByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("HTTP GET response error: %s\n", err)
		return ""
	}
	return string(respByte)
}

func (r *request) HttpPost(url string, header map[string]interface{}, payload string) string {
	request, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		log.Printf("HTTP POST request error: %s\n", err)
		return ""
	}
	for k, v := range header {
		val, ok := v.(string)
		if ok {
			request.Header.Set(k, val)
		}
	}
	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		log.Printf("HTTP POST request error: %s\n", err)
		return ""
	}

	respByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("HTTP POST response error: %s\n", err)
		return ""
	}
	return string(respByte)
}

func (r *request) HttpRequest(method string, url string, header map[string]interface{}, payload string) string {
	request, err := http.NewRequest(method, url, strings.NewReader(payload))
	if err != nil {
		log.Printf("HTTP %s request error: %s\n", method, err)
		return ""
	}
	for k, v := range header {
		val, ok := v.(string)
		if ok {
			request.Header.Set(k, val)
		}
	}
	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		log.Printf("HTTP %s request error: %s\n", method, err)
		return ""
	}

	respByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("HTTP %s response error: %s\n", method, err)
		return ""
	}
	return string(respByte)
}

package net

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
)

func GetStruct(url string, target interface{}, headers map[string]interface{}, printerr bool) *http.Response {
	req := gorequest.New().Get(url)
	if headers != nil {
		for k, v := range headers {
			req.Set(k, v.(string))
		}
	}
	resp, _, errs := req.EndStruct(&target)
	if errs != nil && printerr {
		fmt.Println(errs)
	}
	return resp
}

func Get(url string, headers map[string]interface{}, printerr bool) (body string, response *http.Response) {
	req := gorequest.New().Get(url)
	if headers != nil {
		for k, v := range headers {
			req.Set(k, v.(string))
		}
	}
	resp, body, errs := req.End()
	if errs != nil && printerr {
		fmt.Println(errs)
	}
	return body, resp
}

func PostStruct(url string, headers map[string]interface{}, body interface{}, target interface{}, printerr bool) *http.Response {
	req := gorequest.New().Post(url)
	if headers != nil {
		for k, v := range headers {
			req.Set(k, v.(string))
		}
	}
	resp, _, errs := req.Send(body).EndStruct(&target)
	if errs != nil && printerr {
		fmt.Println(errs)
	}
	return resp
}

func Post(url string, headers map[string]interface{}, br interface{}, printerr bool) (body string, response *http.Response) {
	req := gorequest.New().Post(url)
	if headers != nil {
		for k, v := range headers {
			req.Set(k, v.(string))
		}
	}
	resp, b, errs := req.Send(br).End()
	if errs != nil && printerr {
		fmt.Println(errs)
	}
	return b, resp
}

func PutStruct(url string, headers map[string]interface{}, body interface{}, target interface{}, printerr bool) *http.Response {
	req := gorequest.New().Post(url)
	if headers != nil {
		for k, v := range headers {
			req.Set(k, v.(string))
		}
	}
	resp, _, errs := req.Send(body).EndStruct(&target)
	if errs != nil && printerr {
		fmt.Println(errs)
	}
	return resp
}

func Put(url string, headers map[string]interface{}, br interface{}, printerr bool) (body string, response *http.Response) {
	req := gorequest.New().Post(url)
	if headers != nil {
		for k, v := range headers {
			req.Set(k, v.(string))
		}
	}
	resp, b, errs := req.Send(br).End()
	if errs != nil && printerr {
		fmt.Println(errs)
	}
	return b, resp
}

func DeleteStruct(url string, headers map[string]interface{}, body interface{}, target interface{}, printerr bool) *http.Response {
	req := gorequest.New().Post(url)
	if headers != nil {
		for k, v := range headers {
			req.Set(k, v.(string))
		}
	}
	resp, _, errs := req.Send(body).EndStruct(&target)
	if errs != nil && printerr {
		fmt.Println(errs)
	}
	return resp
}

func Delete(url string, headers map[string]interface{}, br interface{}, printerr bool) (body string, response *http.Response) {
	req := gorequest.New().Post(url)
	if headers != nil {
		for k, v := range headers {
			req.Set(k, v.(string))
		}
	}
	resp, b, errs := req.Send(br).End()
	if errs != nil && printerr {
		fmt.Println(errs)
	}
	return b, resp
}

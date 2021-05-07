// Package network
// Author: youngchan
// CreateDate: 2021/5/7 3:30 下午
// Copyright: ©2021 NEW CORE Technology Co. Ltd. All rights reserved.
// Description:  Http Client 封装
//
package network

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/youngchan1988/gocommon"
	"github.com/youngchan1988/gocommon/cast"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

const (
	HttpGet    = "GET"
	HttpPost   = "POST"
	HttpPut    = "PUT"
	HttpDelete = "DELETE"
)

const (
	HttpContentJson     = "application/json"
	HttpContentFormData = "application/x-www-form-urlencoded"
	HttpContentText     = "application/text"
)

type HttpClient struct {
	Host        string
	MaxConnects int
	ConnTimeout time.Duration

	client *http.Client
}

type HttpResponse struct {
	StatusCode int
	Header     http.Header
	Cookies    []*http.Cookie
	JsonBody   map[string]interface{}
	TextBody   string
}

func NewHttpClient(host string, maxConnects int, connTimeout time.Duration) (*HttpClient, error) {

	if gocommon.IsEmpty(host) {
		return nil, errors.New("host can't be empty")
	}
	if maxConnects == 0 {
		maxConnects = 10
	}
	if connTimeout == 0 {
		connTimeout = 30 * time.Second
	}
	httpClient := &HttpClient{
		Host:        host,
		MaxConnects: maxConnects,
		ConnTimeout: connTimeout,
	}
	tr := &http.Transport{
		MaxIdleConns:    httpClient.MaxConnects,
		IdleConnTimeout: httpClient.ConnTimeout,
	}
	httpClient.client = &http.Client{Transport: tr}
	return httpClient, nil
}

func (c *HttpClient) Req(path string, method string, headers map[string]interface{}, cookies []*http.Cookie, queryParams map[string]interface{}, data interface{}, contentType string) (*HttpResponse, error) {
	httpUrl := fmt.Sprintf("%s%s", c.Host, path)
	var body io.Reader

	if data != nil {
		dv := reflect.ValueOf(data)
		if dv.Kind() == reflect.Ptr {
			dv = dv.Elem()
		}
		if contentType == HttpContentJson && (dv.Kind() == reflect.Slice || dv.Kind() == reflect.Map) {
			//转换json
			j, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			body = strings.NewReader(string(j))
		} else if contentType == HttpContentFormData && dv.Kind() == reflect.Map {
			//postform
			m := data.(map[string]interface{})
			form := url.Values{}
			for k, v := range m {
				form.Add(k, cast.InterfaceToStringWithDefault(v))
			}
			body = strings.NewReader(form.Encode())
		} else {
			//text
			body = strings.NewReader(cast.InterfaceToStringWithDefault(data))
		}
	}
	req, err := http.NewRequest(method, httpUrl, body)
	if err != nil {
		return nil, err
	}

	if !gocommon.IsEmpty(queryParams) {
		//拼接query 参数
		for k, v := range queryParams {
			req.URL.Query().Add(k, cast.InterfaceToStringWithDefault(v))
		}
	}
	req.Header.Add("Content-Type", contentType)
	if !gocommon.IsEmpty(headers) {
		//设置header
		for k, v := range headers {
			req.Header.Add(k, cast.InterfaceToStringWithDefault(v))
		}
	}
	if !gocommon.IsEmpty(cookies) {
		for _, v := range cookies {
			req.AddCookie(v)
		}
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &HttpResponse{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Cookies:    res.Cookies(),
	}

	j := make(map[string]interface{})
	err = json.Unmarshal(b, &j)
	if err == nil {
		r.JsonBody = j
	} else {
		r.TextBody = string(b)
	}
	return r, nil
}

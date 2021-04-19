package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)



var jar, _  = cookiejar.New(nil)

type Client struct {
	url         string
	method      string
	header      map[string]string
	cookies     []http.Cookie
	timeout     time.Duration
	request     *http.Request
	queryParams string
	body        io.Reader
}



//
func New(URL, method string) *Client {
	client := &Client{}
	client.url = URL
	client.method = method
	client.timeout = time.Second * 10
	fmt.Println("new--->")

	return client
}

// 设置cookie
func (c *Client) WithCookies (cookies []http.Cookie) *Client{
	for _, cookie := range cookies{
		c.cookies = append(c.cookies, cookie)
	}
	return c
}

// 设置超时时间
func (c *Client) WithTimeout(t time.Duration) *Client{
	c.timeout = t
	return c
}

// 携带查询参数
func (c *Client) WithQueryParams(queryParams map[string]string) *Client {
	if queryParams == nil {
		return c
	}
	if c.queryParams != "" {
		c.queryParams = c.queryParams + "&"
	}
	for k, v := range queryParams {
		if c.queryParams == "" {
			c.queryParams = k + "=" + v
			continue
		}
		c.queryParams = c.queryParams + "&" + k + "=" + v
	}
	return c
}
// 表单数据
func (c *Client) WithFormBody (data url.Values){
	c.request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	c.method = "POST"
	c.body = bytes.NewBuffer([]byte(data.Encode()))
}

// json数据
func (c *Client) WithRawBody(data interface{}) *Client {
	jsonValue, _ := json.Marshal(data)
	c.body = bytes.NewBuffer([]byte(jsonValue))
	return c
}

//

func (c *Client) Request() (interface{}, error) {
	c.request, _ = http.NewRequest(c.method, c.url, c.body)
	if c.request == nil {
		return nil, errors.New("request is nil,should be init first")
	}
	// 设置请求对象的header信息
	for k, v := range c.header {
		c.request.Header.Add(k, v)
	}
	// 设置cookie
	for _, cookie := range c.cookies {
		c.request.AddCookie(&cookie)
	}
	fmt.Println("jar--->", jar)
	// 设置超时时间
	client := &http.Client{
		Jar: jar,
	}
	client.Timeout = c.timeout

	// 执行请求
	resp, err := client.Do(c.request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 获取请求结果
	var res interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		return nil, err
	}
	return res, nil
}


func HTTPGet(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}
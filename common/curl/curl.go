package curl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	Client *http.Client
}

// HTTP请求参数
type RequestParam struct {
	Url     string
	Method  string
	Header  map[string]interface{}
	Query   map[string]interface{}
	Json    map[string]interface{}
	Form    map[string]interface{}
	Body    string
	Context context.Context
}

func DefaultClient() *Request {
	return &Request{
		Client: http.DefaultClient,
	}
}

// 初始化客户端
func NewClient(client *http.Client) *Request {
	return &Request{
		Client: client,
	}
}

// 发送请求
func (r *Request) Send(requestParam *RequestParam) (string, error) {

	var (
		request *http.Request
		err     error
	)

	// 创建一个新的HTTP请求
	request, err = createRequest(requestParam)
	if err != nil {
		return "", err
	}

	// 设置请求头部信息
	if requestParam.Header != nil {
		for key, value := range requestParam.Header {
			request.Header.Set(key, fmt.Sprint(value))
		}
	}

	// 将请求与提供的context相关联
	if requestParam.Context != nil {
		request = request.WithContext(requestParam.Context)
	}

	// 发送HTTP请求
	result, err := r.Client.Do(request)
	if err != nil {
		return "", err
	}

	defer result.Body.Close()

	// 读取响应体内容，并加入缓冲区
	var buffer bytes.Buffer
	if _, err = io.Copy(&buffer, result.Body); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

// 创建请求
func createRequest(requestParam *RequestParam) (*http.Request, error) {

	switch strings.ToLower(requestParam.Method) {
	case "get":
		return getRequest(requestParam)
	case "post":
		return postRequest(requestParam)
	default:
		return getRequest(requestParam)
	}
}

// get请求
func getRequest(requestParam *RequestParam) (*http.Request, error) {

	// 解析URL
	url, err := url.Parse(requestParam.Url)
	if err != nil {
		return nil, err
	}

	query := url.Query()
	for key, value := range requestParam.Query {
		query.Set(key, fmt.Sprint(value))
	}
	url.RawQuery = query.Encode()

	// 更新请求参数的URL
	requestParam.Url = url.String()

	return http.NewRequest("GET", requestParam.Url, nil)
}

// post请求
func postRequest(requestParam *RequestParam) (*http.Request, error) {

	var body io.Reader

	// Json 传参
	if requestParam.Json != nil {
		// 将json序列化为字节数组
		jsonData, _ := json.Marshal(requestParam.Json)
		body = bytes.NewBuffer(jsonData)
	}

	// Form 传参
	if requestParam.Form != nil {
		// 创建表单数据
		formData := url.Values{}
		for key, value := range requestParam.Form {
			formData.Add(key, fmt.Sprint(value))
		}
		body = strings.NewReader(formData.Encode())
	}

	// Body 传参
	if requestParam.Body != "" {
		body = strings.NewReader(requestParam.Body)
	}

	return http.NewRequest("POST", requestParam.Url, body)
}

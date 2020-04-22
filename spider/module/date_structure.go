package module

import (
	"net/http"
)


//数据的接口
type Data interface {
	//判断数据是否有效
	Valid() bool
}

/**
请求深度一旦创建就不可变（小写，包外不可访问）
*/
//数据请求类型
type Request struct {
	httpReq *http.Request
	depth uint32//请求深度
}

//创建一个新的请求实例
func CreateRequest(httpReq *http.Request,depth uint32) *Request {
	return &Request{httpReq,depth}
}

//获取HTTP请求
func (r *Request) HttpReq() *http.Request {
	return r.httpReq
}

//获取请求的深度
func (r *Request) Depth() uint32 {
	return r.depth
}

func (r *Request) Valid() bool {
	return r.httpReq != nil && r.httpReq.URL != nil
}

//数据响应的类型
type Response struct {
	httpResp *http.Response
	//响应深度
	depth uint32
}

//创建一个响应实例
func CreateResponse(httpResp *http.Response,depth uint32) *Response {
	return &Response{httpResp,depth}
}

func (p *Response) HttpResp() *http.Response {
	return p.httpResp
}

func (p *Response) Depth() uint32 {
	return p.depth
}

func (p *Response) Valid() bool {
	return p.httpResp != nil && p.httpResp.Body != nil
}

//条目类型
type Item map[string]interface{}

func (item Item) Valid() bool {
	return item != nil
}


package module

import (
	"net/http"
)

type MID string
type Counts uint64
type Status uint8

//用于计算组件评分的函数类型
type CalculateScore func(counts Counts) uint64

type ParseResponse func(p *http.Response, pDepath uint32) ([]Data, []error)
type ProcessItem func(item Item) (restult Item, err error)

//组件类型
type Type string

const (
	TYPE_DOWNLOADER Type = "downloader"
	TYPE_ANALYZER   Type = "analyzer"
	TYPE_PIPELINE   Type = "pipeline"
)

var legalTypeletterMap = map[Type]string{
	TYPE_DOWNLOADER: "D",
	TYPE_ANALYZER:   "A",
	TYPE_PIPELINE:   "P",
}

//组件摘要类型(可序列化成json)
type SummaryStruct struct {
	ID        MID         `json:"id"`
	Called    uint64      `json:"called"`
	Accepted  uint64      `json:"accepted"`
	Completed uint64      `json:"completed"`
	Handling  uint64      `json:"handling"`
	Extra     interface{} `json:"extra,omitempty"`
}

//序列号生成器接口的类型
type SNGenertor interface {
	//用于获取预设的最小序列号
	Start() uint64
	//用于获取预设的最大序列号
	Max() uint64
	//获取下一个序列号
	CycleCount() uint64
	//用于获取序列号并准备下一个序列号
	Get() uint64
}

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
	depth   uint32 //请求深度
}

//创建一个新的请求实例
func CreateRequest(httpReq *http.Request, depth uint32) *Request {
	return &Request{httpReq, depth}
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
func CreateResponse(httpResp *http.Response, depth uint32) *Response {
	return &Response{httpResp, depth}
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

type Module interface {
	//获取组件ID
	ID() MID
	//获取当前组件的网络地址的字符串形式
	Addr() string
	//获取当前组件的评分
	Score() uint64
	//用于设置当前组件评分
	SetScore(scroe uint64)
	//获取评分计算器
	ScoreCalculator() CalculateScore
	//获取当前组件被调用的计数
	CalledCount() uint64
	//获取当前组件接收的调用的计数
	//组件一般会由于超负荷或参数有误而拒绝调用
	AcceptedCount() uint64
	//用于获取当前组件已成功完成的调用的计数
	CompletedCount() uint64
	//用于获取当前组件正在处理的调用的数量
	HandlingNumber() uint64
	//用于一次性获取所有计数
	Counts() Counts
	//用于获取组件的摘要
	Summary() SummaryStruct
}

//在网络爬虫真正启动前，应先向组件注册器注册足够的组件实例

//组件注册器接口
type Registrar interface {
	//用于注册实例
	Register(module Module) (bool, error)
	//注销实例
	Unregister(mid MID) (bool, error)
	//获取组件实例(基于评分)
	Get(moduleType Type) (Module, error)
	//获取指定组件类型的所有组件实例
	GetAllByType(moduleType Type) (map[MID]Module, error)
	GetAll() map[MID]Module
	Clear()
}

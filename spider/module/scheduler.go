package module

import "errors"

const (
	//未初始化状态
	SCHED_STATUS_UNIITIALIZED Status = 0
	//正在初始化
	SCHED_STATUS_INITIALIZING Status = 1
	//已初始化
	SCHED_STATUS_INITIALIZED Status = 2
	//正在启动
	SCHED_STATUS_STARTING Status = 3
	//已经启动
	SCHED_STATUS_STARTED Status = 4
	//正在停止
	SCHED_STATUS_STOPPING Status = 5
	//已停止
	SCHED_STATUS_STOPPED Status = 6
)

type Args interface {
	Check() error
}


type RequestArgs struct {
	//可接受的URL的主域名列表
	AcceptedDomains []string `json:"accepted_primary_domains"`
	MaxDepth        uint32   `json:"max_depath"`
}

type DataArgs struct {
	//请求缓冲器的容量
	ReqBufferCap uint32 `json:"req_buffer_cap"`
	//请求缓冲器的最大数量
	ReqMaxBufferNumber   uint32 `json:"req_max_buffer_number`
	RespBufferCap        uint32 `json:"resp_buffer_cap"`
	RespMaxBufferNumber  uint32 `json:"resp_max_buffer_number"`
	ItemBufferCap        uint32 `json:"item_buffer_cap"`
	ItemMaxBufferNumber  uint32 `json:"item_max_buffer_number"`
	ErrorBufferCap       uint32 `json:"error_buffer_cap"`
	ErrorMaxBufferNumber uint32 `json:"error_max_buffer_number"`
}

type ModuleArgs struct {
	//下载器列表
	Downloaders []module.Downloader
	//分析器列表
	Analyzers []module.Analyzers
	//条目处理管道列表
	Pipelines []module.Pipeline
}

type Scheduler interface {
	Init(requestArgs, RequestArgs,
		dataArgs DataArgs,
		moduleArgs ModuleArgs) error

	Start(firstHttpReq *http.Request) error
	Stop() error
	Status() Status
	//各个模块的错误都发会发到该通道
	//ErrorChan用于获得错误的通道
	//若结果值为nil，则说明错误通道不可用或调度器已经停止
	ErrorChan() <-chan error
	//用于判断所有处理模块是否都处于空闲状态
	Idle() bool
	//用于获取摘要实例
	Summary() SchedSummary
}

type NewScheduler struct {
}

func (s *NewScheduler) Init(requestArgs, RequestArgs,
	dataArgs DataArgs,
	moduleArgs ModuleArgs) error {
		return errors.New("Init success")
}

func (s *NewScheduler) Start(firstHttpReq *http.Request) error {
	return errors.New("Init success")
}

func (s *NewScheduler) Stop() error {
	return errors.New("Init success")
}

func (s *NewScheduler) Status() Status {
	return 1
}

func (s *NewScheduler) ErrorChan() <-chan error {
	errCh := make(chan error,20)
	return errCh
}

func (s *NewScheduler) Idle() bool {
	return true
}

func (s *NewScheduler) ummary() SchedSummary {
	
}



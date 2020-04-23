package module

//该接口的实现类型必须是并发安全的
type Downloader interface {
	Module
	Download(r *Request) (*Response, error)
}

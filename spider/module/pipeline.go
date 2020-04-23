package module

//该接口的实现的类型必须是并发安全的
type Pipeline interface {
	Module
	//用于返回当前条目处理管道使用的条目处理函数的列表
	ItemProcessors() []ProcessItem
	//向条目处理管道发送条目
	Send(item Item) []error
	//FailFast方法会返回一个bool，该值表示当前条目管道是否是快速失败的
	FailFast() bool
	SetFailFast(failFast bool)
}

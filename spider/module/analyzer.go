package module

type Analyzer interface {
	Module
	//用于返回当前分析器使用的响应解析函数的列表
	RespParsers() []ParseResponse
	Analyze(p *Response) ([]Data, []error)
}

package enums

type JsonResultCode int

const (
	JRCodeSucc JsonResultCode  = iota
	JRCodeFailed 	= -1	// 操作失败
	JRCode302		= 302	// 跳转至地址
	JRCode401		= 401	// 未授权访问
)

const (
	Deleted = iota - 1
	Disabled
	Enabled
)

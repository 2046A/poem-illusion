/**
 * 一些错误信息结构的定义
 */
package message

const (
	Success = 1
	Failure = 0
)

//好像就是这么简单啊
type ErrMessage struct {
	//错误标示嘛
	Code int
	Text string
}

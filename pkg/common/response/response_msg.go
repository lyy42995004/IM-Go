package response

type ResponseMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func SuccessMsg(data any) *ResponseMsg {
	msg := &ResponseMsg{
		Code: 0,
		Msg:  "SUCCESS",
		Data: data,
	}
	return msg
}

func FailMsg(msg string) *ResponseMsg {
	msgObj := &ResponseMsg{
		Code: -1,
		Msg:  msg,
	}
	return msgObj
}

func FailCodeMsg(code int, msg string) *ResponseMsg {
	msgObj := &ResponseMsg{
		Code: code,
		Msg:  msg,
	}
	return msgObj
}

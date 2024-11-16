package common

type Response struct {
	StatusCode int64  `json:"code"`
	StatusMsg  string `json:"msg"`
	Data       any    `json:"data,omitempty"`
}

// Msg returns the message of the resp
func (r *Response) Msg() string {
	if m, ok := Msg[r.StatusCode]; ok {
		return m
	}
	return ""
}

// GetMsg returns the message of the resp without Response type
func GetMsg(code int64) string {
	if msg, ok := Msg[code]; ok {
		return msg
	}
	return ""
}

// SetNoData prepares the resp without data
func (r *Response) SetNoData(code int64) {
	r.StatusCode = code
	r.StatusMsg = r.Msg()
}

// SetWithData prepares the resp with data
func (r *Response) SetWithData(code int64, data interface{}) {
	r.StatusCode = code
	r.StatusMsg = r.Msg()
	r.Data = data
}

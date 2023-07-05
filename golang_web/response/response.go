package response

import (
	"net/http"
	"sync"
)

type resResult struct {
	Data    []interface{} `json:"data"`
	Status  uint32        `json:"status"`
	Message string        `json:"message"`
}

type Response struct {
	HttpStatus int
	R          resResult
}

var pool = sync.Pool{
	New: func() interface{} {
		return &Response{}
	},
}

func NewResponse(status int, code uint32, data ...interface{}) *Response {
	response := pool.Get().(*Response)
	response.HttpStatus = status
	response.R.Status = code
	response.R.Message = MessageForCode[code]
	response.R.Data = data
	if len(data) == 0 {
		response.R.Data = make([]interface{}, 0)
	} else {
		response.R.Data = data
	}

	return response
}

func PutResponse(res *Response) {
	pool.Put(res)
}

func NewResponseOk(code uint32, data ...interface{}) *Response {
	return NewResponse(http.StatusOK, code, data...)
}

func NewResponseOkND(code uint32) *Response {
	return NewResponse(http.StatusOK, code)
}

func RQuerySuccess(data ...interface{}) *Response {
	return NewResponse(http.StatusOK, QuerySuccess, data...)
}

func RQueryFailed() *Response {
	return NewResponse(http.StatusOK, QueryFailed)
}

func ROperateFailed() *Response {
	return NewResponse(http.StatusOK, OperateFailed)
}

func ROperateSuccess() *Response {
	return NewResponse(http.StatusOK, OperateSuccess)
}

func RDeleteFailed() *Response {
	return NewResponse(http.StatusOK, DeleteFailed)
}

func RDeleteSuccess() *Response {
	return NewResponse(http.StatusOK, DeleteSuccess)
}

func RGetResourceFailed() *Response {
	return NewResponse(http.StatusOK, GetResourceFailed)
}

func RGetResourceSuccess(data ...interface{}) *Response {
	return NewResponse(http.StatusOK, GetResourceSuccess, data...)
}

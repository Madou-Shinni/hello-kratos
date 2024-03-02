package common

import (
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	httpnet "net/http"
)

type HttpResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// EncoderResponse encode response to the HTTP response.
func EncoderResponse() http.EncodeResponseFunc {
	return func(w http.ResponseWriter, r *http.Request, i interface{}) (err error) {
		if i == nil {
			return nil
		}

		// 自定义响应
		resp := &HttpResponse{
			Code: httpnet.StatusOK,
			Msg:  "success",
			Data: i,
		}

		codec := encoding.GetCodec("json")
		marshal, err := codec.Marshal(resp)
		if err != nil {
			return errors.InternalServer("json", err.Error())
		}
		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(marshal)
		if err != nil {
			return errors.InternalServer("json", err.Error())
		}

		return nil
	}
}

// EncoderError encode error to the HTTP response.
func EncoderError() http.EncodeErrorFunc {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		se := errors.FromError(err)

		resp := &HttpResponse{
			Code: int(se.Code),
			Msg:  se.Message,
			Data: nil,
		}

		// 自定义响应
		codec := encoding.GetCodec("json")
		marshal, err := codec.Marshal(resp)
		if err != nil {
			w.WriteHeader(httpnet.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(marshal)

		return
	}
}

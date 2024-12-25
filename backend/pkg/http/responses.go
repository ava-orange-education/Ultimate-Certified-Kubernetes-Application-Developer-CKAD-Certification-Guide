package http

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Error(ctx context.Context, w http.ResponseWriter, err error) {
	var httpStatus int
	var apiResp *ErrorResponse

	var apiErr *Error
	if errors.As(err, &apiErr) {
		httpStatus = apiErr.HTTPStatus
	} else {
		httpStatus = http.StatusInternalServerError
		apiResp = &ErrorResponse{
			Message: err.Error(),
		}
	}

	JSON(ctx, w, httpStatus, apiResp)
}

func ErrorWithStatus(ctx context.Context, w http.ResponseWriter, status int, errMsg string) {
	apiResp := &ErrorResponse{
		Message: errMsg,
	}

	JSON(ctx, w, status, apiResp)
}

func JSON(ctx context.Context, w http.ResponseWriter, status int, resp any) {
	b, err := json.Marshal(resp)
	if err != nil {
		log.Default().Printf("unmarshalling response: %v", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if _, err := w.Write(b); err != nil {
		log.Default().Printf("writing response: %v", err)
	}
}

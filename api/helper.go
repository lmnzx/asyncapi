package api

import (
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

type ApiResponse[T any] struct {
	Data    *T     `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

type ErrWithStatus struct {
	status int
	err    error
}

func (e *ErrWithStatus) Error() string {
	return e.err.Error()
}

func NewErrWithStatus(status int, err error) *ErrWithStatus {
	return &ErrWithStatus{status: status, err: err}
}

func handler(f func(ctx *fasthttp.RequestCtx) error) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if err := f(ctx); err != nil {
			status := fasthttp.StatusInternalServerError
			msg := http.StatusText(status)
			if e, ok := err.(*ErrWithStatus); ok {
				status = e.status
				msg = http.StatusText(e.status)
				if status == fasthttp.StatusBadRequest || status == fasthttp.StatusConflict {
					msg = e.err.Error()
				}
			}

			log.Error().Err(err).Int("status", status).Str("message", msg).Msg("error executing handler")

			if err := encode(ApiResponse[struct{}]{
				Message: msg,
			}, status, ctx); err != nil {
				log.Error().Err(err).Msg("encoding error")
			}
		}
	}
}

func encode[T any](v T, status int, ctx *fasthttp.RequestCtx) error {
	ctx.SetContentType("application/json; charset=utf-8")
	ctx.SetStatusCode(status)
	if err := json.NewEncoder(ctx.Response.BodyWriter()).Encode(v); err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}
	return nil
}

type Validator interface {
	Validate() error
}

func decode[T Validator](ctx *fasthttp.RequestCtx) (T, error) {
	var t T
	if err := json.Unmarshal(ctx.Request.Body(), &t); err != nil {
		return t, fmt.Errorf("failed to decode request body: %w", err)
	}
	if err := t.Validate(); err != nil {
		return t, err
	}
	return t, nil
}

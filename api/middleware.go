package api

import (
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
)

func NewLoggerMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		log.Info().Str("method", string(ctx.Method())).
			Str("path", string(ctx.Path())).
			Msg("http request")
		next(ctx)
	}
}

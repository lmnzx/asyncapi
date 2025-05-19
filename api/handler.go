package api

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/lmnzx/asyncapi/store"
	"github.com/valyala/fasthttp"
)

func (s *Server) ping(ctx *fasthttp.RequestCtx) error {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody([]byte("pong"))
	return nil
}

type SignupRequest store.CreateUserParams

func (r SignupRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

func (s *Server) signupHandler(ctx *fasthttp.RequestCtx) error {
	req, err := decode[SignupRequest](ctx)
	if err != nil {
		return NewErrWithStatus(fasthttp.StatusBadRequest, err)
	}

	_, err = s.store.CreateUser(ctx, store.CreateUserParams{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return NewErrWithStatus(fasthttp.StatusConflict, fmt.Errorf("user already exists"))
		} else {
			return NewErrWithStatus(fasthttp.StatusInternalServerError, err)
		}
	}

	if err := encode(ApiResponse[struct{}]{
		Message: "successfully signed up user",
	}, fasthttp.StatusCreated, ctx); err != nil {
		return NewErrWithStatus(fasthttp.StatusInternalServerError, err)
	}

	return nil
}

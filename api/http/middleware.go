package http

import (
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
)

// ValidateAuthTokenMiddleware - middleware to check the auth token.
func (s *Server) ValidateAuthTokenMiddleware(ctx huma.Context, next func(huma.Context)) {
	authHeader := ctx.Header("Authorization")
	if authHeader == "" {
		huma.WriteErr(s.API, ctx, http.StatusUnauthorized, "Missing Authorization Header")
		return
	}

	requestToken := strings.TrimPrefix(authHeader, "Bearer ")
	if requestToken != s.AuthToken {
		huma.WriteErr(s.API, ctx, http.StatusUnauthorized, "Invalid Authorization Header")
		return
	}

	next(ctx)
}

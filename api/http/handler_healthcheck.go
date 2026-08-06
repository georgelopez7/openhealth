package http

import (
	"context"
)

func (s *Server) HealthcheckHandler(ctx context.Context, input *HealthcheckInput) (*HealthcheckResponse, error) {
	return &HealthcheckResponse{
		Body: HealthcheckBody{
			Status: "ok",
		},
	}, nil
}

package http

type HealthcheckInput struct{}

type HealthcheckResponse struct {
	Body HealthcheckBody
}

type HealthcheckBody struct {
	Status string `json:"status" doc:"Service health status"`
}

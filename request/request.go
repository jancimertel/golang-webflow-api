package request

import "net/http"

type HttpMethod string

const (
	MethodGet  HttpMethod = http.MethodGet
	MethodPost HttpMethod = http.MethodPost
)

// Envelope holds request data to webflow api
type Envelope struct {
	Method HttpMethod
	Path   string
	Body   interface{}
}

// This file is safe to edit. Once it exists it will not be overwritten.

package monitoring

import (
	"github.com/go-openapi/runtime/middleware"
	"gitlab.com/kilianpaquier/craft/examples/golang-api/restapi/operations/monitoring"
)

// UnsecuredPing - GET /unsecured/ping.
//
// checks API health and retrieves ping result.
func UnsecuredPing(params monitoring.UnsecuredPingParams) middleware.Responder {
	_ = params.HTTPRequest.Context() // request context
	return middleware.NotImplemented("not implemented")
}

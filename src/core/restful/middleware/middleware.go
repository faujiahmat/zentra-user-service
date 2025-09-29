package middleware

import "github.com/faujiahmat/zentra-user-service/src/core/restful/client"

type Middleware struct {
	restfulClient *client.Restful
}

func New(rc *client.Restful) *Middleware {
	return &Middleware{
		restfulClient: rc,
	}
}

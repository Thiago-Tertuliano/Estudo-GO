package httpmw

import "net/http"

// Chain aplica middlewares em ordem: o primeiro da lista fica mais "por fora".
func Chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

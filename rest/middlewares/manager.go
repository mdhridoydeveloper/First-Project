package middleware

import "net/http"

type middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...middleware) *Manager {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
	return mngr
}

func (mngr *Manager) With(Handler http.Handler, middlewares ...middleware) http.Handler {
	h := Handler
	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h

}
func (mngr *Manager) WrapMux(Handler http.Handler) http.Handler {
	h := Handler
	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h)
	}

	return h

}

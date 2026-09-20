package middleware

import "net/http"

type middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddleware []middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleware: make([]middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...middleware) *Manager {
	mngr.globalMiddleware = append(mngr.globalMiddleware, middlewares...)
	return mngr
}

func (mngr *Manager) With(next http.Handler, middlewares ...middleware) http.Handler {
	n := next
	//middlewares = [hudai, logger] // idx= hudai - 0, logger - 1
	//middleware.hudai(middleware.logger(http.HandlerFunc(handlers.GetProducts)))
	// for i := len(middlewares) - 1; i >= 0; i-- { // i =0
	// 	middleware := middlewares[i] // hudai
	// 	n = middleware(n)
	// middlewares
	for _, middleware := range middlewares {
		n = middleware(n)
	}
	for _, globalMiddleware := range mngr.globalMiddleware {
		n = globalMiddleware(n)
	}

	return n

}

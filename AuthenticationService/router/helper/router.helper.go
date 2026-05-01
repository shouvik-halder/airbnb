package routerhelper

import "github.com/go-chi/chi/v5"

type Router interface {
	Register(r chi.Router)
}

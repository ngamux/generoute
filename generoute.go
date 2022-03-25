package generoute

import (
	"net/http"

	"github.com/ngamux/ngamux"
)

type RestResourceController interface {
	Index(rw http.ResponseWriter, r *http.Request) error
	Show(rw http.ResponseWriter, r *http.Request) error
	Store(rw http.ResponseWriter, r *http.Request) error
	Update(rw http.ResponseWriter, r *http.Request) error
	Destroy(rw http.ResponseWriter, r *http.Request) error
}

func RestResource(mux *ngamux.Ngamux, path string, controller RestResourceController) {
	mux.Get(path, controller.Index)
	mux.Get(path+"/:id", controller.Show)
	mux.Post(path, controller.Store)
	mux.Put(path+"/:id", controller.Update)
	mux.Patch(path+"/:id", controller.Update)
	mux.Delete(path+"/:id", controller.Destroy)
}

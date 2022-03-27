package generoute

import (
	"net/http"

	"github.com/ngamux/ngamux"
)

type ResourceController interface {
	RestResourceController
	PageStore(rw http.ResponseWriter, r *http.Request) error
	PageUpdate(rw http.ResponseWriter, r *http.Request) error
}

func Resource(mux *ngamux.Ngamux, path string, controller ResourceController) {
	mux.Get(path+"/create", controller.PageStore)
	mux.Get(path+"/:id/edit", controller.PageUpdate)
	RestResource(mux, path, controller)
}

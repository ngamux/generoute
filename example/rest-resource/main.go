package main

import (
	"net/http"

	"github.com/ngamux/generoute"
	"github.com/ngamux/ngamux"
)

type UserController struct{}

func (c *UserController) Index(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Index")
}

func (c *UserController) Show(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Show")
}

func (c *UserController) Store(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Store")
}

func (c *UserController) Update(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Update")
}

func (c *UserController) Destroy(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Destroy")
}

func main() {
	mux := ngamux.New()

	generoute.RestResource(mux, "/users", &UserController{})
	http.ListenAndServe(":8080", mux)
}

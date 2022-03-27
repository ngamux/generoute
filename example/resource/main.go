package main

import (
	"log"
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

func (c *UserController) PageStore(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "PageStore")
}

func (c *UserController) Update(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Update")
}

func (c *UserController) PageUpdate(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "PageUpdate")
}

func (c *UserController) Destroy(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Destroy")
}

func main() {
	mux := ngamux.New()

	generoute.Resource(mux, "/users", &UserController{})
	log.Fatal(http.ListenAndServe(":8080", mux))
}

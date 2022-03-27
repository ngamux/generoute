package generoute

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-must/must"
	"github.com/ngamux/ngamux"
)

func TestResource(t *testing.T) {
	routes := []struct {
		name   string
		method string
		path   string
	}{
		{"Index", http.MethodGet, "/users"},
		{"Show", http.MethodGet, "/users/1"},
		{"Store", http.MethodPost, "/users"},
		{"PageStore", http.MethodGet, "/users/create"},
		{"Update", http.MethodPut, "/users/1"},
		{"PageUpdate", http.MethodGet, "/users/1/edit"},
		{"Destroy", http.MethodDelete, "/users/1"},
	}

	t.Run("route should registered", func(t *testing.T) {
		must := must.New(t)

		for _, route := range routes {
			expected := route.name

			rw := httptest.NewRecorder()
			r := httptest.NewRequest(route.method, route.path, nil)

			mux := ngamux.New()
			Resource(mux, "/users", controller)
			mux.ServeHTTP(rw, r)
			must.Equal(expected+"\n", rw.Body.String())
		}

	})
}

var controller = &userController{}

type userController struct{}

func (c *userController) Index(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Index")
}

func (c *userController) Show(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Show")
}

func (c *userController) Store(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Store")
}

func (c *userController) PageStore(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "PageStore")
}

func (c *userController) Update(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Update")
}

func (c *userController) PageUpdate(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "PageUpdate")
}

func (c *userController) Destroy(rw http.ResponseWriter, r *http.Request) error {
	return ngamux.String(rw, "Destroy")
}

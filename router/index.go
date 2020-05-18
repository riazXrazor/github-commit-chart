package router

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Router type for router
type Router struct {
	method  string
	path    string
	handler http.HandlerFunc
}

var r = mux.NewRouter()

// Init type for router
func Init() *mux.Router {

	STATICPATH := os.Getenv("STATICPATH")
	api := r.PathPrefix("/api/").Subrouter()
	for _, route := range Routers {
		api.Path(route.path).
			Methods(route.method).
			HandlerFunc(route.handler)
	}

	// Serve static files
	r.PathPrefix(STATICPATH).Handler(http.StripPrefix(STATICPATH, http.FileServer(http.Dir("./client/public/static/"))))

	// Serve index page on all unhandled routes
	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./client/public/index.html")
	})

	return r
}

/*
 * ServiceDesk
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"servicedesk/middlewares"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middlewares.AuthMiddleware)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello ServiceDesk!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreateRequest",
		strings.ToUpper("Post"),
		"/support",
		CreateRequest,
	},

	Route{
		"DeleteRequest",
		strings.ToUpper("Delete"),
		"/support/{requestId[0-9]+}",
		DeleteRequest,
	},

	Route{
		"GetRequest",
		strings.ToUpper("Get"),
		"/support/{requestId[0-9]+}",
		GetRequest,
	},

	Route{
		"ListRequests",
		strings.ToUpper("Get"),
		"/support",
		ListRequests,
	},

	Route{
		"UpdateRequest",
		strings.ToUpper("Put"),
		"/support/{requestId[0-9]+}",
		UpdateRequest,
	},
}
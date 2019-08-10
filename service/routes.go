
package service
import "net/http"
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
// Defines the type Routes which is just an array of Route structs.
type Routes []Route
// Initialize our routes
var routes = Routes{
	Route{
		"GetAccount",                                     // Name
		"GET",                                            // HTTP method
		"/accounts/{accountId}",                          // Route pattern
		GetAccount,
	},
}
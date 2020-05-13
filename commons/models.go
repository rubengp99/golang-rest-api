package commons

import "net/http"

//NewResponse is what is sended to the requester through API
type NewResponse struct{
	Counter 		int 			`json:"count"`
	DataResponse 	interface{} 	`json:"data"`
	Message 		APIResponse 	`json:"response"`
}

// Route is the struct to define a new route on the router
type Route struct {
    Method      string				`json:"request"`
    Pattern     string				`json:"path"`
    HandlerFunc http.HandlerFunc	`json:"func"`
}

// Routes an array of routes for each EndPoint
type Routes []Route

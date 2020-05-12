package commons

import "net/http"

//NewResponse is what is sended to the requester through API
type NewResponse struct{
	Count 		int 		`json:"count"`
	Data 		interface{} `json:"data"`
	Message 	APIResponse `json:"response"`
}

// Route is the struct to define a new route on the router
type Route struct {
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Routes an array of routes for each EndPoint
type Routes []Route

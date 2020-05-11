package global

// APIResponse is the result of a request
type APIResponse struct{
    message string
    code int
}

// Ok returns HTTP 200 OK response
func Ok() APIResponse{
    return APIResponse{ "Ok", 200 }
}

// Created returns HTTP 201 Created response
func Created() APIResponse {
    return APIResponse{ "Record created", 201 }
}

// Update returns HTTP 201 Created response (Case: Updated)
func Update() APIResponse {
    return APIResponse{ "Record updated", 201 }
}

// Deleted returns HTTP 200 OK response (Case: Deleted)
func Deleted() APIResponse {
    return APIResponse{ "Record deleted", 200 }
}

// Empty returns HTTP 200 OK response (Case: OK but no records found)
func Empty() APIResponse {
    return APIResponse{ "This entity is empty", 200 }
}

// InvalidID returns HTTP 400 BAD REQUEST response (Case: ID with bad format)
func InvalidID() APIResponse {
    return APIResponse{ "The given id is not valid", 400 }
}

// BadRequest returns HTTP 400 BAD REQUEST response 
func BadRequest() APIResponse {
    return APIResponse{ "Bad Request", 400 }
}

// Unauthorized returns HTTP 401 UNAUTHORIZED response 
func Unauthorized() APIResponse {
    return APIResponse{ "The credentials are invalids", 401 }
}

// Forbidden returns HTTP 403 FORBIDDEN response 
func Forbidden() APIResponse {
    return APIResponse{ "You are not allowed to use this route", 403 }
}

// ElementNotFound returns HTTP 404 NOT FOUND response (Case: Element not found)
func ElementNotFound() APIResponse {
    return APIResponse{ "The element does not exist",404 }
}

// RouteNotFound returns HTTP 404 NOT FOUND response (Case: Route not found)
func RouteNotFound() APIResponse {
    return APIResponse{ "Route not found", 404 }
}

// BadFormat returns HTTP 406 NOT ACCEPTABLE response
func BadFormat() APIResponse {
    return APIResponse{ "Bad Format", 406 }
}

// InternalServerError returns HTTP 500 INTERNAL SERVER ERROR
func InternalServerError() APIResponse {
    return APIResponse{ "Internal server error", 500 }
}
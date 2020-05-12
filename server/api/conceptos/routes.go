package conceptos

import (
	"github.com/rubengp99/golang-rest-api/commons"
)

// Routes of the EndPoint conceptos
var Routes = commons.Routes{
    commons.Route{
        "GET",
		"/api/conceptos/",
		GetAll,
    },
}

package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Clik API", func() {
	Title("The Clik API")
	Description("The API for Clik dating app")
	Host("localhost:8080")
	Scheme("http")
})

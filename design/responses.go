package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var Token = MediaType("application/clk.token+json", func() {
	Description("A JWT token")
	Attributes(func() {
		Attribute("token", String, func() {
			Example("{JWT_TOKEN}")
		})
	})
})

var Match = MediaType("application/clk.match+json", func() {
	Description("A match")
	Attributes(func() {
		Attribute("initiated_user", User, "The user who initiated the match")
		Attribute("accepted_user", User, "The user who accepted the match")
	})

	View("default", func() {
		Attribute("initiated_user")
		Attribute("accepted_user")
	})
})

// User is the account resource media type.
var User = MediaType("application/clk.user+json", func() {
	Description("A user")
	Attributes(func() {
		Attribute("id", Integer, "ID of user", func() {
			Example(1)
		})
		Attribute("email", String, "Email of the user", func() {
			Format("email")
			Example("bob@gmail.com")
		})
		Attribute("first_name", String, "First name of the user", func() {
			Example("John")
		})
		Attribute("last_name", String, "Last name of the user", func() {
			Example("Snow")
		})
		Attribute("address", String, "Address of user", func() {
			Example("123 Douglas Road, Victoria, BC")
		})
		Attribute("phone_number", String, "Phone number of the user", func() {
			Example("555-555-5555")
		})
		Attribute("created_at", DateTime, "When the user was created")
	})

	View("default", func() {
		Attribute("id")
		Attribute("email")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("address")
		Attribute("phone_number")
	})
})

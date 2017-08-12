package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var CreateDecisionPayload = Type("CreateDecisionPayload", func() {
	Attribute("user_id", Integer, func() {
		Minimum(1)
		Example(1)
	})

	Attribute("likes", Boolean, func() {
		Example(true)
	})
	Required("user_id", "likes")
})

var CreateMessagePayload = Type("CreateMessagePayload", func() {
	Attribute("message", String, func() {
		MinLength(1)
		MaxLength(5000)
		Example("This is a nice app!")
	})

	Attribute("conversation_id", Integer, func() {
		Minimum(1)
		Example(1)
	})
	Required("message", "conversation_id")
})

var LoginPayload = Type("LoginPayload", func() {
	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(400)
		Format("email")
		Example("jamesbond@gmail.com")
	})

	Attribute("password", String, func() {
		MinLength(5)
		MaxLength(100)
		Example("abcd1234")
	})
	Required("email", "password")
})

var RegisterPayload = Type("RegisterPayload", func() {
	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(150)
		Format("email")
		Example("jamesbond@gmail.com")
	})

	Attribute("first_name", String, func() {
		MinLength(1)
		MaxLength(200)
		Example("John")
	})

	Attribute("last_name", String, func() {
		MinLength(1)
		MaxLength(200)
		Example("Doe")
	})

	Attribute("address", String, func() {
		MinLength(5)
		MaxLength(300)
		Example("123 Dallas Road")
	})

	Attribute("phone_number", String, func() {
		MinLength(7)
		MaxLength(20)
		Example("1-800-123-5555")
	})

	Attribute("password", String, func() {
		MinLength(5)
		MaxLength(100)
		Example("abcd1234")
	})

	Required("email", "password", "first_name", "last_name")
})

var AddImagesPayload = Type("AddImagesPayload", func() {
	Attribute("images", Files, )
})

var UpdateProfilePayload = Type("UpdateProfilePayload", func() {
	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(150)
		Format("email")
		Example("jamesbond@gmail.com")
	})

	Attribute("first_name", String, func() {
		MinLength(1)
		MaxLength(200)
		Example("John")
	})

	Attribute("last_name", String, func() {
		MinLength(1)
		MaxLength(200)
		Example("Doe")
	})

	Attribute("address", String, func() {
		MinLength(5)
		MaxLength(300)
		Example("123 Dallas Road")
	})

	Attribute("phone_number", String, func() {
		MinLength(7)
		MaxLength(20)
		Example("1-800-123-5555")
	})
})
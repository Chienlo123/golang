package routes

import "github.com/gofiber/fiber/v2"

func SetupUserRoutes(app *fiber.App) {

	// Read all data table
	app.Get("/users", getUsers)

	//Specific id
	app.Get("/users/:id", getUser)

	// Create ng table
	app.Post("/users12345", CreateUser)

	//Registration
	app.Post("/users123", RegisterUser)

	// Update ng data sa table
	app.Patch("/users/:id", UpdateUser)

	// Delete ng data sa table
	app.Delete("/users/:id", deleteUser)

	//Cloneuser
	app.Post("/users123/:id", CloneUser)

	//XmlTry create file
	app.Get("/xmltry", XmlTry)

	//XmlTry construct root
	app.Get("/xmltry123", Xml2ndTry)

	app.Post("/xmltry12345", GetXml)

	app.Post("/gg", GG)

}

package routes

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"os"
	"sample/db"
	models "sample/models"
	"strings"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

// create another all data in user and also aditional create table in postgres databased
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	db.DB.Create(&user)
	return c.XML(user)
}

// register account
func RegisterUser(c *fiber.Ctx) error {
	var newUser models.User

	fmt.Println("            ================ REGISTRATION ================")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("          Enter username: ")
	newUser.Name, _ = reader.ReadString('\n')
	newUser.Name = strings.TrimSpace(newUser.Name)

	fmt.Print("          Enter Lastname: ")
	newUser.Lastname, _ = reader.ReadString('\n')
	newUser.Lastname = strings.TrimSpace(newUser.Lastname)

	fmt.Print("          Enter Address: ")
	newUser.Address, _ = reader.ReadString('\n')
	newUser.Address = strings.TrimSpace(newUser.Address)

	fmt.Print("          Enter Contact: ")
	newUser.Contact, _ = reader.ReadString('\n')
	newUser.Contact = strings.TrimSpace(newUser.Contact)

	fmt.Print("          Enter email: ")
	newUser.Email, _ = reader.ReadString('\n')
	newUser.Email = strings.TrimSpace(newUser.Email)

	fmt.Print("          Enter Password: ")
	newUser.Password, _ = reader.ReadString('\n')
	newUser.Password = strings.TrimSpace(newUser.Password)

	db.DB.Create(&newUser)
	return c.Status(fiber.StatusCreated).XML(newUser)
}

// Display all users
func getUsers(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Find(&users)
	return c.XML(users)
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	db.DB.First(&user, id)
	return c.XML(user)
}

// func GetUsers(c *fiber.Ctx) error {
// 	var users []User
// 	db.Find(&users)
// 	return c.JSON(users)
// }

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return err
	}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	db.DB.Save(&user)
	return c.XML(user)
}

// // Update Users
// func UpdateUser(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var user models.User
// 	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
// 		return c.Status(fiber.StatusNotFound).XML(fiber.Map{"error": "User not found"})
// 	}

// 	fmt.Println("            ================ UPDATE DATA ================")

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("          Enter new username: ")
// 	user.Name, _ = reader.ReadString('\n')
// 	user.Name = strings.TrimSpace(user.Name)

// 	fmt.Print("          Enter new Lastname: ")
// 	user.Lastname, _ = reader.ReadString('\n')
// 	user.Lastname = strings.TrimSpace(user.Lastname)

// 	fmt.Print("          Enter new Address: ")
// 	user.Address, _ = reader.ReadString('\n')
// 	user.Address = strings.TrimSpace(user.Address)

// 	fmt.Print("          Enter new Contact: ")
// 	user.Contact, _ = reader.ReadString('\n')
// 	user.Contact = strings.TrimSpace(user.Contact)

// 	fmt.Print("          Enter new email: ")
// 	user.Email, _ = reader.ReadString('\n')
// 	user.Email = strings.TrimSpace(user.Email)

// 	fmt.Print("          Enter new Password: ")
// 	user.Password, _ = reader.ReadString('\n')
// 	user.Password = strings.TrimSpace(user.Password)

// 	db.DB.Save(&user)
// 	return c.XML(user)
// }

//DELETE USER
// func deleteUser(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var user models.User
// 	db.DB.Delete(&user, id)
// 	return nil
// }

// Delete Users
// func deleteUser(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	var user models.User

// 	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
// 	}

// 	db.DB.Delete(&user)
// 	return c.JSON(fiber.Map{"message": "User deleted"})
// }

func deleteUser(c *fiber.Ctx) error {

	userResponse := models.User{}
	id := c.Params("id")
	var user models.User

	result := db.DB.Debug().Raw("DELETE FROM users WHERE id = ?", id).Scan(&userResponse)

	if result != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User Deleted"})
	}

	db.DB.Delete(&user)
	return c.JSON(fiber.Map{"message": "User Not Deleted"})
}

//Clone ng user choose any id

func CloneUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).XML(fiber.Map{"error": "User not found"})
	}

	var clonedUser models.User
	clonedUser.Name = "cloned_" + user.Name
	clonedUser.Lastname = "cloned_" + user.Lastname
	clonedUser.Address = "cloned_" + user.Address
	clonedUser.Contact = "cloned_" + user.Contact
	clonedUser.Email = "cloned_" + user.Email
	clonedUser.Password = "cloned_" + user.Password
	db.DB.Create(&clonedUser)

	return c.XML(clonedUser)
}

// xml Create file
func XmlTry(c *fiber.Ctx) error {

	xmlData := models.UsersInfo{

		CorporateID:      "0434235325",
		BranchID:         "fs43",
		TransactionKey:   "fa0328332",
		RequestRefNo:     "00097323214324300214",
		TransactionType:  "bdfsaf",
		RequestTimeStamp: "gfd32gd",
		TerminalID:       "093424254325",
		Address:          "Laguna",
	}
	file, err := os.Create("xmlData.xml")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return nil
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "\t")

	// Encode the Users struct as XML and write to the file and display it
	err = encoder.Encode(xmlData)
	if err != nil {
		fmt.Println("Error encoding XML:", err)

	}
	return nil
}

// this is the xml construct root
func Xml2ndTry(c *fiber.Ctx) error {

	deped := models.Deped{
		School: "Sto Tomas Integraed higschool",

		PrincipalHead: models.PrincipalHead{
			Principal: "Mr james bond",

			TeacherBody: models.TeacherBody{
				Teacher: "teacher",

				Student: models.Student{

					StudentName:  "chienlo",
					StudentID:    "432532532525325",
					Section:      "Rizal",
					MajorSubject: "BSIS",
				},
			},
		},
	}

	xmlInfo, err := xml.MarshalIndent(deped, "", "")
	if err != nil {

		return c.Status(http.StatusInternalServerError).SendString("error generating XML response")
	}

	c.Response().Header.Set("Content-Type", "application/xml")
	return c.Send(xmlInfo)

	// xmlInfo, err := xml.Marshal(xmlData)
	// if err != nil {

	// 	return c.Status(http.StatusInternalServerError).SendString("error generating XML response")
	// }

	// c.Response().Header.Set("Content-Type", "application/xml")
	// return c.Send(xmlInfo)

}

func GetXml(c *fiber.Ctx) error {

	TryReturn, TryErr := xml.MarshalIndent(models.Student{}, "", "")

	if TryErr != nil {

		return c.SendString(TryErr.Error())
	}

	response, respErr := http.NewRequest(http.MethodGet, "http://127.0.0.1:3000/users", bytes.NewBuffer(TryReturn))

	if respErr != nil {

		return c.SendString(respErr.Error())
	}

	response.Header.Set("Content-Type", "application/xml")
	client := &http.Client{}

	resp, clientErr := client.Do(response)
	if clientErr != nil {

		return c.SendString(clientErr.Error())
	}
	defer resp.Body.Close()

	respBody, readErr := io.ReadAll(resp.Body)
	if readErr != nil {

		return c.SendString(readErr.Error())
	}

	c.Set("Content-Type", "application/xml")

	return c.Send(respBody)

}

func GG(c *fiber.Ctx) error {

	fmt.Println("               ====================================")

	fmt.Println("                        LOGIN ACCOUNT:")
	fmt.Print("                Username: ")
	var name string
	fmt.Scanln(&name)

	fmt.Print("                Password: ")
	var password string
	fmt.Scanln(&password)

	fmt.Println("               ====================================")

	var user models.User
	result := db.DB.First(&user, "name = ?", name)
	if result.Error != nil {
		fmt.Println("Invalid credentials")
		os.Exit(1)
	}

	if password != user.Password {
		fmt.Println("Invalid credentials")
		os.Exit(1)
	}

	response := models.LoginResponse{
		Message: fmt.Sprintf("Welcome, %s!", user.Name),
	}

	xmlResponse, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatal("Failed to generate XML response")
	}

	fmt.Println("Login successful!")
	fmt.Println(string(xmlResponse))

	// db.DB.First(&user)
	// return c.XML(user)

	db.DB.First(&xmlResponse)
	return c.XML(xmlResponse)

}

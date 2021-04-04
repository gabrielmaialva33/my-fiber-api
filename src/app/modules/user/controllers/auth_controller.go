package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-api/src/app/modules/user/models"
	"go-api/src/app/modules/user/repositories"
	"go-api/src/app/utils"
	"time"
)

type UsersServices struct {
	ur repositories.UserRepositoryInterface
}

func UsersController(ur repositories.UserRepositoryInterface) *UsersServices {
	return &UsersServices{
		ur: ur,
	}
}

// -> Register : function
func (s *UsersServices) Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newUser, err := s.ur.Create(user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(newUser)
}

// -> Login : function
func (s *UsersServices) Login(c *fiber.Ctx) error {
	var user *models.User

	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	u, userErr := s.ur.GetUserByEmailAndPassword(user)
	if userErr != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": userErr,
		})
	}

	token, err := utils.GenerateJwt(u.Id)
	if err != nil {
		c.Status(400)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "go-api",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // 1 day
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("go-api")

	id, _ := utils.ParseJwt(cookie)

	//var user models.User

	//database.DB.Where("id = ?", id).First(&user)

	return c.JSON(id)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "go-api",
		Value:    "no-cookie",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

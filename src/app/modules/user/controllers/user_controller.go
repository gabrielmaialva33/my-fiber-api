package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-api/src/app/modules/user/models"
	"go-api/src/app/modules/user/repositories"
	"go-api/src/app/modules/user/validators"
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

func (s *UsersServices) Index(c *fiber.Ctx) error {
	users := models.Users{}
	var err error
	users, err = s.ur.Index()
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(users.PublicUsers())
}

func (s *UsersServices) Show(c *fiber.Ctx) error {
	id := c.Params("user_id")
	if err := validators.UUIDUserValidator(id); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user, err := s.ur.Show(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(user)
}

func (s *UsersServices) Create(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := validators.CreateUserValidator(user); err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newUser, err := s.ur.Create(&user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(newUser)
}

func (s *UsersServices) Login(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := validators.LoginUserValidator(user); err != nil {
		c.Status(fiber.StatusUnprocessableEntity)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	u, userErr := s.ur.GetUserByEmailAndPassword(&user)
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
		"message": "success", "user_id": u.Id,
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "go-api",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

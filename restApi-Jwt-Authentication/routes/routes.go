package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"restApi-Jwt-Authentication/common"
	"restApi-Jwt-Authentication/data"
	"xorm.io/xorm"
)

func SetupRoutes(app *fiber.App, engine *xorm.Engine) {
	app.Post("/signup", func(c *fiber.Ctx) error {
		req := new(common.SignupRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}
		if req.Username == "" || req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid Request")
		}

		hash, err := common.HashPassword(req.Password)
		if err != nil {
			return err
		}
		user := &data.User{
			Username: req.Username,
			Email:    req.Email,
			Password: hash,
		}

		_, err = engine.Insert(user)
		if err != nil {
			return err
		}
		token, exp, err := common.CreateJWTToken(common.User(*user))
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"token": token, "exp": exp})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		req := new(common.LoginRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}
		if req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid login Request")
		}

		user := new(data.User)
		has, err := engine.Where("email = ?", req.Email).Desc("id").Get(user)
		if err != nil {
			return err
		}
		if !has {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}
		if !common.CheckPasswordHash(req.Password, user.Password) {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
		}

		token, exp, err := common.CreateJWTToken(common.User(*user))
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"token": token, "exp": exp})
	})

	private := app.Group("/private")
	private.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
	private.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"path":    "private",
		})
	})

	public := app.Group("/public")
	public.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"path":    "public",
		})
	})
}

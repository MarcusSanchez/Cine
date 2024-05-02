package controller

import (
	"cine/entity/model"
	"cine/entity/schemas"
	"cine/pkg/fault"
	"cine/server/middleware"
	"cine/service"
	"github.com/MarcusSanchez/go-parse"
	"github.com/MarcusSanchez/go-z"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthController struct {
	auth service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{auth: authService}
}

func (ac *AuthController) Routes(router fiber.Router, mw *middleware.Middleware) {
	router.Post("/register", mw.SignedOut, ac.Register)
	router.Post("/login", mw.SignedOut, ac.Login)
	router.Delete("/logout", mw.SignedIn, mw.CSRF, ac.Logout)
	router.Post("/authenticate", mw.SignedIn, mw.CSRF, ac.Authenticate)
}

// Register [POST] /api/register
func (ac *AuthController) Register(c *fiber.Ctx) error {

	type Payload struct {
		DisplayName    string `json:"display_name"    z:"display_name"`
		Email          string `json:"email"           z:"email"`
		Username       string `json:"username"        z:"username"`
		Password       string `json:"password"        z:"password"`
		ProfilePicture string `json:"profile_picture" z:"profile_picture"`
	}

	p, err := parse.JSON[Payload](c.Body())
	if err != nil {
		return fault.BadRequest(err.Error())
	}

	schema := z.Struct{
		"display_name":    schemas.DisplayNameSchema,
		"email":           schemas.EmailSchema,
		"username":        schemas.UsernameSchema,
		"password":        schemas.PasswordSchema,
		"profile_picture": schemas.ProfilePictureSchema,
	}
	if errs := schema.Validate(p); errs != nil {
		return fault.Validation(errs.One())
	}

	user, session, err := ac.auth.Register(c.Context(),
		&service.RegisterInput{
			DisplayName:    p.DisplayName,
			Email:          p.Email,
			Username:       p.Username,
			Password:       p.Password,
			ProfilePicture: p.ProfilePicture,
		},
	)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"user": user, "session": session})
}

// Login [POST] /api/login
func (ac *AuthController) Login(c *fiber.Ctx) error {

	type Payload struct {
		Username string `json:"username" z:"username"`
		Password string `json:"password" z:"password"`
	}

	p, err := parse.JSON[Payload](c.Body())
	if err != nil {
		return fault.BadRequest(err.Error())
	}

	schema := z.Struct{
		"username": schemas.UsernameSchema,
		"password": schemas.PasswordSchema,
	}
	if errs := schema.Validate(p); errs != nil {
		return fault.BadRequest(errs.One())
	}

	user, session, err := ac.auth.Login(c.Context(), p.Username, p.Password)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"session": session, "user": user})
}

// Logout [DELETE] /api/logout
func (ac *AuthController) Logout(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)

	err := ac.auth.Logout(c.Context(), session)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

// Authenticate [POST] /api/authenticate
func (ac *AuthController) Authenticate(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)

	user, session, err := ac.auth.Authenticate(c.Context(), session)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user, "session": session})
}

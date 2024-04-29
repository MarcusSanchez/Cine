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

type UserController struct {
	user service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{user: userService}
}

func (uc *UserController) Routes(router fiber.Router, mw *middleware.Middleware) {
	users := router.Group("/users")
	users.Put("/", mw.SignedIn, mw.CSRF, uc.UpdateUser)
	users.Delete("/", mw.SignedIn, mw.CSRF, uc.DeleteUser)
}

// UpdateUser [PUT] /api/users
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {

	type Payload struct {
		DisplayName    *string `json:"display_name,optional"    z:"display_name" `
		Email          *string `json:"email,optional"           z:"email"`
		Username       *string `json:"username,optional"        z:"username"`
		Password       *string `json:"password,optional"        z:"password"`
		ProfilePicture *string `json:"profile_picture,optional" z:"profile_picture"`
	}

	p, err := parse.JSON[Payload](c.Body())
	if err != nil {
		return fault.BadRequest(err.Error())
	}

	schema := z.Struct{
		"display_name":    schemas.DisplayNameSchema.Optional(),
		"email":           schemas.EmailSchema.Optional(),
		"username":        schemas.UsernameSchema.Optional(),
		"password":        schemas.PasswordSchema.Optional(),
		"profile_picture": schemas.ProfilePictureSchema.Optional(),
	}
	if errs := schema.Validate(p); errs != nil {
		return fault.Validation(errs.One())
	}

	session := c.Locals("session").(*model.Session)

	user, err := uc.user.UpdateUser(c.Context(),
		session.UserID, &model.UserU{
			DisplayName:    p.DisplayName,
			Username:       p.Username,
			Email:          p.Email,
			Password:       p.Password,
			ProfilePicture: p.ProfilePicture,
		},
	)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}

// DeleteUser [DELETE] /api/users
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	session := c.Locals("session").(*model.Session)

	err := uc.user.DeleteUser(c.Context(), session.UserID)
	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

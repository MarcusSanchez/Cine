package schemas

import (
	"github.com/MarcusSanchez/go-z"
	"github.com/go-resty/resty/v2"
	"net/url"
)

var DisplayNameSchema = z.String().
	Min(3, "display name must be at least 3 characters").
	Max(32, "display name must be at most 32 characters")

var UsernameSchema = z.String().
	Min(3, "username must be at least 3 characters").
	Max(32, "username must be at most 32 characters")

var EmailSchema = z.String().
	Min(3, "email must be at least 3 characters").
	Max(254, "email must be at most 254 characters").
	Email("email must be a valid email address")

var PasswordSchema = z.String().
	Min(8, "password must be at least 8 characters").
	Max(50, "password must be at most 50 characters").
	Regex(`[A-Z]`, "password must contain at least one uppercase letter").
	Regex(`[0-9]`, "password must contain at least one number").
	Regex(`[!@#$%^&*()_+{}|:<>?~]`, "password must contain at least one special character")

var ProfilePictureSchema = z.String().Custom(
	func(image string) bool {
		imageURL, err := url.ParseRequestURI(image)
		if err != nil {
			return false
		}

		resp, err := resty.New().R().Head(imageURL.String())
		if err != nil || resp.IsError() {
			return false
		}

		switch resp.Header().Get("Content-Type") {
		case "image/jpeg", "image/png", "image/webp":
			return true
		default:
			return false
		}
	}, "profile picture must be a valid image URL",
)

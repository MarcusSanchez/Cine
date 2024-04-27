package schemas

import "github.com/MarcusSanchez/go-z"

var CommentContentSchema = z.String().
	Min(1, "content must not be empty").
	Max(280, "content must be at most 280 characters")

package schemas

import (
	"github.com/MarcusSanchez/go-z"
	"github.com/google/uuid"
)

var CommentContentSchema = z.String().
	Min(1, "content must not be empty").
	Max(280, "content must be at most 280 characters")

var CommentReplyingToIDSchema = z.String().
	Custom(func(s string) bool {
		if _, err := uuid.Parse(s); err != nil {
			return false
		}
		return true
	}, "replying_to_id must be a valid UUID")

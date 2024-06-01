package schemas

import "github.com/MarcusSanchez/go-z"

var ReviewContentSchema = z.String().
	Min(1, "content must be at least 1 character long").
	Max(140, "content must be at most 140 characters long")

var ReviewRatingSchema = z.Int().
	Gt(0, "rating must be greater than 0").
	Lte(10, "rating must be less than or equal to 10")

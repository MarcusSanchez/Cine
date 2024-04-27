package schemas

import "github.com/MarcusSanchez/go-z"

var MediaTypeSchema = z.String().
	In([]string{"movie", "show"}, "type must be either 'movie' or 'show'")

package schemas

import "github.com/MarcusSanchez/go-z"

var MediaTypeSchema = z.String().
	In([]string{"movie", "show"}, "type must be either 'movie' or 'show'")

var MovieListSchema = z.String().
	In(
		[]string{"popular", "topRated", "nowPlaying", "upcoming"},
		"list must be either 'popular', 'topRated', 'nowPlaying', or 'upcoming",
	)

var ShowListSchema = z.String().
	In(
		[]string{"popular", "topRated", "onTheAir", "airingToday"},
		"list must be either 'popular', 'topRated', 'onTheAir', or 'airingToday",
	)

package tmdb

type Movie struct {
	ID           int     `json:"id"`
	Overview     string  `json:"overview"`
	GenreIDs     []int   `json:"genre_ids"`
	BackdropPath string  `json:"backdrop_path"`
	Language     string  `json:"original_language"`
	Popularity   float64 `json:"popularity"`
	PosterPath   string  `json:"poster_path"`
	ReleaseDate  string  `json:"release_date"`
	Title        string  `json:"title"`
	Video        bool    `json:"video"`
}

type Show struct {
	ID           int     `json:"id"`
	Overview     string  `json:"overview"`
	GenreIDs     []int   `json:"genre_ids"`
	BackdropPath string  `json:"backdrop_path"`
	Language     string  `json:"original_language"`
	Popularity   float64 `json:"popularity"`
	PosterPath   string  `json:"poster_path"`
	FirstAirDate string  `json:"first_air_date"`
	Name         string  `json:"name"`
}

type DetailedGenre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DetailedMovie struct {
	BackdropPath     *string         `json:"backdrop_path,optional"`
	Budget           int             `json:"budget"`
	Genres           []DetailedGenre `json:"genres"`
	Homepage         string          `json:"homepage"`
	ID               int             `json:"id"`
	OriginCountry    []string        `json:"origin_country"`
	OriginalLanguage string          `json:"original_language"`
	OriginalTitle    string          `json:"original_title"`
	Overview         string          `json:"overview"`
	Popularity       float64         `json:"popularity"`
	PosterPath       *string         `json:"poster_path,optional"`
	ReleaseDate      *string         `json:"release_date,optional"`
	Revenue          int             `json:"revenue"`
	Runtime          int             `json:"runtime"`
	Status           string          `json:"status"`
	Tagline          string          `json:"tagline"`
	Title            string          `json:"title"`
	Video            bool            `json:"video"`
	VoteAverage      float64         `json:"vote_average"`
	VoteCount        int             `json:"vote_count"`
}

type CreatedBy struct {
	ID           int     `json:"id"`
	CreditID     string  `json:"credit_id"`
	Name         string  `json:"name"`
	OriginalName string  `json:"original_name"`
	Gender       int     `json:"gender"`
	ProfilePath  *string `json:"profile_path"`
}

type Network struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type Season struct {
	AirDate      string  `json:"air_date"`
	EpisodeCount int     `json:"episode_count"`
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   *string `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
	VoteAverage  float64 `json:"vote_average"`
}

type DetailedShow struct {
	Adult            bool            `json:"adult"`
	BackdropPath     *string         `json:"backdrop_path,optional"`
	CreatedBy        []CreatedBy     `json:"created_by"`
	EpisodeRunTime   []int           `json:"episode_run_time"`
	FirstAirDate     *string         `json:"first_air_date,optional"`
	Genres           []DetailedGenre `json:"genres"`
	Homepage         string          `json:"homepage"`
	ID               int             `json:"id"`
	InProduction     bool            `json:"in_production"`
	Languages        []string        `json:"languages"`
	LastAirDate      *string         `json:"last_air_date,optional"`
	Name             string          `json:"name"`
	Networks         []Network       `json:"networks"`
	NumberOfEpisodes int             `json:"number_of_episodes"`
	NumberOfSeasons  int             `json:"number_of_seasons"`
	OriginCountry    []string        `json:"origin_country"`
	OriginalLanguage string          `json:"original_language"`
	OriginalName     string          `json:"original_name"`
	Overview         string          `json:"overview"`
	Popularity       float64         `json:"popularity"`
	PosterPath       *string         `json:"poster_path,optional"`
	Seasons          []Season        `json:"seasons"`
	Status           string          `json:"status"`
	Tagline          string          `json:"tagline"`
	Type             string          `json:"type"`
	VoteAverage      float64         `json:"vote_average"`
	VoteCount        int             `json:"vote_count"`
}

type MovieCredits struct {
	ID   int         `json:"id"`
	Cast []MovieCast `json:"cast"`
	Crew []MovieCrew `json:"crew"`
}

type MovieCast struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CastID             int     `json:"cast_id"`
	Character          string  `json:"character"`
	CreditID           string  `json:"credit_id"`
	Order              int     `json:"order"`
}

type MovieCrew struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        string  `json:"profile_path"`
	CreditID           string  `json:"credit_id"`
	Department         string  `json:"department"`
	Job                string  `json:"job"`
}

type ShowCredits struct {
	ID   int        `json:"id"`
	Cast []ShowCast `json:"cast"`
	Crew []ShowCrew `json:"crew"`
}

type ShowCast struct {
	Adult              bool       `json:"adult"`
	Gender             int        `json:"gender"`
	ID                 int        `json:"id"`
	KnownForDepartment string     `json:"known_for_department"`
	Name               string     `json:"name"`
	OriginalName       string     `json:"original_name"`
	Popularity         float64    `json:"popularity"`
	ProfilePath        string     `json:"profile_path"`
	Roles              []ShowRole `json:"roles"`
	TotalEpisodeCount  int        `json:"total_episode_count"`
	Order              int        `json:"order"`
}

type ShowRole struct {
	CreditID     string `json:"credit_id"`
	Character    string `json:"character"`
	EpisodeCount int    `json:"episode_count"`
}

type ShowCrew struct {
	Adult              bool      `json:"adult"`
	Gender             int       `json:"gender"`
	ID                 int       `json:"id"`
	KnownForDepartment string    `json:"known_for_department"`
	Name               string    `json:"name"`
	OriginalName       string    `json:"original_name"`
	Popularity         float64   `json:"popularity"`
	ProfilePath        *string   `json:"profile_path"`
	Jobs               []ShowJob `json:"jobs"`
	Department         string    `json:"department"`
	TotalEpisodeCount  int       `json:"total_episode_count"`
}

type ShowJob struct {
	CreditID     string `json:"credit_id"`
	Job          string `json:"job"`
	EpisodeCount int    `json:"episode_count"`
}

type DetailedSeason struct {
	ID           string    `json:"_id"`
	AirDate      string    `json:"air_date"`
	Episodes     []Episode `json:"episodes"`
	Name         string    `json:"name"`
	Overview     string    `json:"overview"`
	PosterPath   string    `json:"poster_path"`
	SeasonNumber int       `json:"season_number"`
	VoteAverage  float64   `json:"vote_average"`
}

type Episode struct {
	AirDate        string  `json:"air_date"`
	EpisodeNumber  int     `json:"episode_number"`
	EpisodeType    string  `json:"episode_type"`
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Overview       string  `json:"overview"`
	ProductionCode string  `json:"production_code"`
	Runtime        int     `json:"runtime"`
	SeasonNumber   int     `json:"season_number"`
	ShowID         int     `json:"show_id"`
	StillPath      string  `json:"still_path"`
	VoteAverage    float64 `json:"vote_average"`
	VoteCount      int     `json:"vote_count"`
}

type MovieList string

const (
	MovieListNowPlaying MovieList = "nowPlaying"
	MovieListPopular    MovieList = "popular"
	MovieListTopRated   MovieList = "topRated"
	MovieListUpcoming   MovieList = "upcoming"
)

type ShowList string

const (
	ShowListAiringToday ShowList = "airingToday"
	ShowListPopular     ShowList = "popular"
	ShowListTopRated    ShowList = "topRated"
	ShowListOnTheAir    ShowList = "onTheAir"
)

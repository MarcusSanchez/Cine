package mocks

import (
	"cine/pkg/tmdb"
)

var _ tmdb.API = (*APIMock)(nil)

type APIMock struct {
	SearchMoviesFn         func(query string, filter ...tmdb.SearchMovieFilter) ([]tmdb.Movie, error)
	SearchShowsFn          func(query string, filter ...tmdb.SearchShowFilter) ([]tmdb.Show, error)
	GetMovieFn             func(ref int) (*tmdb.DetailedMovie, error)
	GetMovieCreditsFn      func(ref int) (*tmdb.MovieCredits, error)
	ListNowPlayingMoviesFn func() ([]tmdb.Movie, error)
	ListPopularMoviesFn    func() ([]tmdb.Movie, error)
	ListTopRatedMoviesFn   func() ([]tmdb.Movie, error)
	ListUpcomingMoviesFn   func() ([]tmdb.Movie, error)
	GetShowFn              func(ref int) (*tmdb.DetailedShow, error)
	GetShowCreditsFn       func(ref int) (*tmdb.ShowCredits, error)
	GetShowSeasonDetailsFn func(ref int, seasonNumber int) (*tmdb.DetailedSeason, error)
	ListAiringTodayShowsFn func() ([]tmdb.Show, error)
	ListPopularShowsFn     func() ([]tmdb.Show, error)
	ListTopRatedShowsFn    func() ([]tmdb.Show, error)
	ListOnTheAirShowsFn    func() ([]tmdb.Show, error)
}

func NewTMDB() *APIMock {
	return &APIMock{}
}

func (m *APIMock) SearchMovies(query string, filter ...tmdb.SearchMovieFilter) ([]tmdb.Movie, error) {
	if m.SearchMoviesFn != nil {
		return m.SearchMoviesFn(query, filter...)
	}
	return []tmdb.Movie{}, nil
}

func (m *APIMock) SearchShows(query string, filter ...tmdb.SearchShowFilter) ([]tmdb.Show, error) {
	if m.SearchShowsFn != nil {
		return m.SearchShowsFn(query, filter...)
	}
	return []tmdb.Show{}, nil
}

func (m *APIMock) GetMovie(ref int) (*tmdb.DetailedMovie, error) {
	if m.GetMovieFn != nil {
		return m.GetMovieFn(ref)
	}
	return &tmdb.DetailedMovie{}, nil
}

func (m *APIMock) GetMovieCredits(ref int) (*tmdb.MovieCredits, error) {
	if m.GetMovieCreditsFn != nil {
		return m.GetMovieCreditsFn(ref)
	}
	return &tmdb.MovieCredits{}, nil
}

func (m *APIMock) ListNowPlayingMovies() ([]tmdb.Movie, error) {
	if m.ListNowPlayingMoviesFn != nil {
		return m.ListNowPlayingMoviesFn()
	}
	return []tmdb.Movie{}, nil
}

func (m *APIMock) ListPopularMovies() ([]tmdb.Movie, error) {
	if m.ListPopularMoviesFn != nil {
		return m.ListPopularMoviesFn()
	}
	return []tmdb.Movie{}, nil
}

func (m *APIMock) ListTopRatedMovies() ([]tmdb.Movie, error) {
	if m.ListTopRatedMoviesFn != nil {
		return m.ListTopRatedMoviesFn()
	}
	return []tmdb.Movie{}, nil
}

func (m *APIMock) ListUpcomingMovies() ([]tmdb.Movie, error) {
	if m.ListUpcomingMoviesFn != nil {
		return m.ListUpcomingMoviesFn()
	}
	return []tmdb.Movie{}, nil
}

func (m *APIMock) GetShow(ref int) (*tmdb.DetailedShow, error) {
	if m.GetShowFn != nil {
		return m.GetShowFn(ref)
	}
	return &tmdb.DetailedShow{}, nil
}

func (m *APIMock) GetShowCredits(ref int) (*tmdb.ShowCredits, error) {
	if m.GetShowCreditsFn != nil {
		return m.GetShowCreditsFn(ref)
	}
	return &tmdb.ShowCredits{}, nil
}

func (m *APIMock) GetShowSeasonDetails(ref int, seasonNumber int) (*tmdb.DetailedSeason, error) {
	if m.GetShowSeasonDetailsFn != nil {
		return m.GetShowSeasonDetailsFn(ref, seasonNumber)
	}
	return &tmdb.DetailedSeason{}, nil
}

func (m *APIMock) ListAiringTodayShows() ([]tmdb.Show, error) {
	if m.ListAiringTodayShowsFn != nil {
		return m.ListAiringTodayShowsFn()
	}
	return []tmdb.Show{}, nil
}

func (m *APIMock) ListPopularShows() ([]tmdb.Show, error) {
	if m.ListPopularShowsFn != nil {
		return m.ListPopularShowsFn()
	}
	return []tmdb.Show{}, nil
}

func (m *APIMock) ListTopRatedShows() ([]tmdb.Show, error) {
	if m.ListTopRatedShowsFn != nil {
		return m.ListTopRatedShowsFn()
	}
	return []tmdb.Show{}, nil
}

func (m *APIMock) ListOnTheAirShows() ([]tmdb.Show, error) {
	if m.ListOnTheAirShowsFn != nil {
		return m.ListOnTheAirShowsFn()
	}
	return []tmdb.Show{}, nil
}

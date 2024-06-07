export type User = {
  id: string,
  username: string,
  display_name: string,
  profile_picture: string
}

export type UserStats = {
  following_count: number,
  followers_count: number,
  likes_count: number,
  comments_count: number,
  reviews_count: number,
  lists_count: number,
  followed: boolean
}

export type Session = {
  id: string,
  user_id: string,
  csrf: string,
  token: string,
  expiration: string
}

export type Review = {
  id: string;
  user_id: string;
  media_id: string;
  content: string;
  rating: number;
  created_at: string;
  updated_at?: string | null;
}

export enum MediaType {
  Movie = "movie",
  Show = "show"
}

export type Movie = {
  id: number;
  overview: string;
  genre_ids: number[];
  backdrop_path: string;
  original_language: string;
  popularity: number;
  poster_path: string;
  release_date: string;
  title: string;
  video: boolean;
};

export type Show = {
  id: number;
  overview: string;
  genre_ids: number[];
  backdrop_path: string;
  original_language: string;
  popularity: number;
  poster_path: string;
  first_air_date: string;
  name: string;
};

export type Comment = {
  id: string;
  user_id: string | null;
  media_id: string;
  replying_to_id: string | null;
  content: string;
  created_at: string;
  updated_at: string | null;
}

export type DetailedComment = {
  user: User;
  comment: Comment;
  replies_count: number;
  likes_count: number;
  liked_by_user: boolean;
}

export type DetailedReview = {
  user: User;
  review: Review;
}

export type DetailedGenre = {
  id: number,
  name: string
}

export type DetailedMovie = {
  backdrop_path: string | null;
  budget: number;
  genres: DetailedGenre[];
  homepage: string;
  id: number;
  imdb_id: string;
  origin_country: string[];
  original_language: string;
  original_title: string;
  overview: string;
  popularity: number;
  poster_path: string | null;
  release_date: string | null;
  revenue: number;
  runtime: number;
  status: string;
  tagline: string;
  title: string;
  video: boolean;
  vote_average: number;
  vote_count: number;
};

export type MovieCredits = {
  id: number;
  cast: Cast[];
  crew: Crew[];
}

export type Cast = {
  adult: boolean;
  gender: number;
  id: number;
  known_for_department: string;
  name: string;
  original_name: string;
  popularity: number;
  profile_path: string | null;
  cast_id: number;
  character: string;
  credit_id: string;
  order: number;
}

export type Crew = {
  adult: boolean;
  gender: number;
  id: number;
  known_for_department: string;
  name: string;
  original_name: string;
  popularity: number;
  profile_path: string | null;
  credit_id: string;
  department: string;
  job: string;
}

export type CreatedBy = {
  id: number;
  credit_id: string;
  name: string;
  original_name: string;
  gender: number;
  profile_path?: string | null;
}

export type Network = {
  id: number;
  logo_path: string;
  name: string;
  origin_country: string;
}

export type Season = {
  air_date: string;
  episode_count: number;
  id: number;
  name: string;
  overview: string;
  poster_path?: string | null;
  season_number: number;
  vote_average: number;
}

export type DetailedShow = {
  backdrop_path?: string | null;
  created_by: CreatedBy[];
  episode_run_time: number[];
  first_air_date?: string | null;
  genres: DetailedGenre[];
  homepage: string;
  id: number;
  in_production: boolean;
  languages: string[];
  last_air_date?: string | null;
  name: string;
  networks: Network[];
  number_of_episodes: number;
  number_of_seasons: number;
  origin_country: string[];
  original_language: string;
  original_name: string;
  overview: string;
  popularity: number;
  poster_path?: string | null;
  seasons: Season[];
  status: string;
  tagline: string;
  type: string;
  vote_average: number;
  vote_count: number;
}

export type ShowCredits = {
  id: number;
  cast: ShowCast[];
  crew: ShowCrew[];
}

export type ShowCast = {
  adult: boolean;
  gender: number;
  id: number;
  known_for_department: string;
  name: string;
  original_name: string;
  popularity: number;
  profile_path: string;
  roles: ShowRole[];
  total_episode_count: number;
  order: number;
}

export type ShowRole = {
  credit_id: string;
  character: string;
  episode_count: number;
}

export type ShowCrew = {
  adult: boolean;
  gender: number;
  id: number;
  known_for_department: string;
  name: string;
  original_name: string;
  popularity: number;
  profile_path: string | null;
  jobs: ShowJob[];
  department: string;
  total_episode_count: number;
}

export type ShowJob = {
  credit_id: string;
  job: string;
  episode_count: number;
}

export type DetailedSeason = {
  _id: string;
  air_date: string;
  episodes: Episode[];
  name: string;
  overview: string;
  poster_path: string;
  season_number: number;
  vote_average: number;
}

export type Episode = {
  air_date: string;
  episode_number: number;
  episode_type: string;
  id: number;
  name: string;
  overview: string;
  production_code: string;
  runtime: number;
  season_number: number;
  show_id: number;
  still_path: string;
  vote_average: number;
  vote_count: number;
}

export enum MovieList {
  NowPlaying = "nowPlaying",
  Popular = "popular",
  TopRated = "topRated",
  Upcoming = "upcoming",
}

export enum ShowList {
  AiringToday = "airingToday",
  Popular = "popular",
  TopRated = "topRated",
  OnTheAir = "onTheAir",
}

export type List = {
  id: string,
  owner_id: string,
  name: string,
  is_public: boolean,
  created_at: string,
  updated_at?: string
}

export type DetailedList = {
  list: List,
  members: User[],
  movies: Media[],
  shows: Media[]
}

export type Media = {
  id: string,
  ref: number,
  media_type: MediaType,
  overview: string,
  backdrop_path?: string,
  language: string,
  poster_path?: string,
  release_date?: string,
  title: string,
  created_at: string,
  updated_at?: string
}


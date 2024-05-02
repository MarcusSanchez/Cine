type User = {
  id: string,
  username: string,
  display_name: string,
  profile_picture: string
}

type Session = {
  id: string,
  user_id: string,
  csrf: string,
  token: string,
  expiration: string
}

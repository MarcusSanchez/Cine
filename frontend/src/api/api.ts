type APIError = {
  error: string;
  message: string;
}

export default class API {
  static url = process.env.API_URL;

  static async authenticate(sessionToken: string, csrf: string) {
    const response = await fetch(`${API.url}/authenticate`, {
      method: "POST",
      headers: {
        "X-Session-Token": sessionToken,
        "X-CSRF-Token": csrf,
      },
    });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { user: User, session: Session };
  }

  static async register(input: {
    email: string,
    password: string,
    username: string,
    display_name: string,
    profile_picture: string
  }) {
    const response = await fetch(`${API.url}/register`, {
      method: "POST",
      body: JSON.stringify(input),
    });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { user: User, session: Session };
  }

  static async login(username: string, password: string) {
    const response = await fetch(`${API.url}/login`, {
      method: "POST",
      body: JSON.stringify({ username, password }),
    });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { user: User, session: Session };
  }

  static async logout(csrf: string, sessionToken: string) {
    const response = await fetch(`${API.url}/logout`, {
      method: "DELETE",
      headers: {
        "X-CSRF-Token": csrf,
        "X-Session-Token": sessionToken,
      },
    });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }
  }
}
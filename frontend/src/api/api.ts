import { FormSchema as ReviewFormSchema } from "@/components/UserReview";
import {
  Comment,
  DetailedComment,
  DetailedMovie,
  DetailedReview, DetailedSeason,
  DetailedShow,
  MediaType,
  MovieCredits,
  Review,
  Session,
  ShowCredits,
  User,
  UserStats
} from "@/models/models";

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

  static async fetchUser(user: string, sessionToken?: string) {
    const headers: HeadersInit = sessionToken ? { "X-Session-Token": sessionToken } : {};
    const response = await fetch(`${API.url}/users/${user}`, { headers: headers });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { user: User };
  }

  static async fetchUserStats(user: string, sessionToken?: string) {
    const headers: HeadersInit = sessionToken ? { "X-Session-Token": sessionToken } : {};
    const response = await fetch(`${API.url}/users/detailed/${user}`, { headers: headers });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    const data = await response.json() as { detailed_user: UserStats };
    return { stats: data.detailed_user };
  }

  static async updateUser(
    csrf: string, sessionToken: string,
    input: {
      username?: string,
      display_name?: string,
      password?: string,
      profile_picture?: string,
    }
  ) {
    for (const key in input) {
      if (!input[key as keyof typeof input]) delete input[key as keyof typeof input];
    }

    const response = await fetch(`${API.url}/users`, {
      method: "PUT",
      headers: {
        "X-CSRF-Token": csrf,
        "X-Session-Token": sessionToken,
      },
      body: JSON.stringify(input),
    });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { user: User };
  }

  static async deleteUser(csrf: string, sessionToken: string) {
    const response = await fetch(`${API.url}/users`, {
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

    return { success: true };
  }

  static async fetchMovie(ref: number, sessionToken?: string) {
    const headers: HeadersInit = sessionToken ? { "X-Session-Token": sessionToken } : {};
    const response = await fetch(`${API.url}/medias/movie/${ref}`, { headers: headers });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { detailed_movie: DetailedMovie };
  }

  static async fetchMovieCredits(ref: number, sessionToken?: string) {
    const headers: HeadersInit = sessionToken ? { "X-Session-Token": sessionToken } : {};
    const response = await fetch(`${API.url}/medias/movie/${ref}/credits`, { headers: headers });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { movie_credits: MovieCredits };
  }

  static async fetchShow(ref: number, sessionToken?: string) {
    const headers: HeadersInit = sessionToken ? { "X-Session-Token": sessionToken } : {};
    const response = await fetch(`${API.url}/medias/show/${ref}`, { headers: headers });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { detailed_show: DetailedShow };
  }

  static async fetchShowCredits(ref: number, sessionToken?: string) {
    const headers: HeadersInit = sessionToken ? { "X-Session-Token": sessionToken } : {};
    const response = await fetch(`${API.url}/medias/show/${ref}/credits`, { headers: headers });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { show_credits: ShowCredits };
  }

  static async fetchShowSeasonDetails(ref: number, season: number, sessionToken?: string) {
    const headers: HeadersInit = sessionToken ? { "X-Session-Token": sessionToken } : {};
    const response = await fetch(`${API.url}/medias/show/${ref}/season/${season}`, { headers: headers });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { detailed_season: DetailedSeason };
  }

  static async fetchReviews(media: MediaType, ref: number, sessionToken?: string) {
    const headers: HeadersInit = sessionToken ? { "X-Session-Token": sessionToken } : {};
    const response = await fetch(`${API.url}/reviews/${media}/${ref}`, { headers: headers });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { detailed_reviews: DetailedReview[] };
  }

  static async createReview(
    csrf: string,
    sessionToken: string,
    data: ReviewFormSchema,
    media: MediaType,
    ref: number,
  ) {
    const response = await fetch(`${API.url}/reviews/${media}/${ref}`, {
      method: "POST",
      headers: {
        "X-CSRF-Token": csrf,
        "X-Session-Token": sessionToken,
      },
      body: JSON.stringify(data),
    });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { review: Review };
  }

  static async editReview(
    csrf: string,
    sessionToken: string,
    data: ReviewFormSchema,
    reviewID: string,
  ) {
    const response = await fetch(`${API.url}/reviews/${reviewID}`, {
      method: "PUT",
      headers: {
        "X-CSRF-Token": csrf,
        "X-Session-Token": sessionToken,
      },
      body: JSON.stringify(data),
    });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { review: Review };
  }

  static async deleteReview(csrf: string, sessionToken: string, reviewID: string) {
    const response = await fetch(`${API.url}/reviews/${reviewID}`, {
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

    return { success: true };
  }

  static async fetchComments(sessionToken: string, media: MediaType, ref: number) {
    const headers: HeadersInit = { "X-Session-Token": sessionToken };
    const response = await fetch(`${API.url}/comments/${media}/${ref}`, { headers });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { detailed_comments: DetailedComment[] };
  }

  static async createComment(
    csrf: string,
    sessionToken: string,
    content: string,
    media: MediaType,
    ref: number,
    replyingTo?: string
  ) {
    const response = await fetch(`${API.url}/comments/${media}/${ref}`, {
      method: "POST",
      headers: {
        "X-CSRF-Token": csrf,
        "X-Session-Token": sessionToken,
      },
      body: JSON.stringify({ content, replying_to_id: replyingTo }),
    });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { comment: Comment };
  }

  static async deleteComment(csrf: string, sessionToken: string, commentID: string) {
    const response = await fetch(`${API.url}/comments/${commentID}`, {
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

    return { success: true };
  }

  static async likeComment(csrf: string, sessionToken: string, commentID: string) {
    const response = await fetch(`${API.url}/comments/like/${commentID}`, {
      method: "POST",
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

  static async unlikeComment(csrf: string, sessionToken: string, commentID: string) {
    const response = await fetch(`${API.url}/comments/like/${commentID}`, {
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

  static async fetchReplies(sessionToken: string, commentID: string) {
    const headers: HeadersInit = { "X-Session-Token": sessionToken };
    const response = await fetch(`${API.url}/comments/${commentID}/replies`, { headers });
    if (!response.ok) {
      const data = await response.json() as APIError;
      throw new Error(data.message);
    }

    return await response.json() as { replies: DetailedComment[] };
  }


}
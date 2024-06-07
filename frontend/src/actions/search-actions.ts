"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { Movie, Show } from "@/models/models";

type MovieResults = {
  success: true, movies: Movie[]
} | {
  success: false, error: string
};

export async function searchMoviesAction(query: string, page: number = 1): Promise<MovieResults> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session token not found" };

    const { movies } = await API.searchMovies(query, page, cookie.value);

    return { success: true, movies: movies };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}

type ShowResults = {
  success: true, shows: Show[]
} | {
  success: false, error: string
};

export async function searchShowsAction(query: string, page: number = 1): Promise<ShowResults> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session token not found" };

    const { shows } = await API.searchShows(query, page, cookie.value);

    return { success: true, shows: shows };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}

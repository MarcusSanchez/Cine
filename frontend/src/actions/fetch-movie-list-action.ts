"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { Movie, MovieList } from "@/models/models";

type Result = {
  success: true, movies: Movie[]
} | {
  success: false, error: string
};

export default async function fetchMovieListAction(list: MovieList): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session token not found" };

    const { movies } = await API.fetchMovieList(list, cookie.value);

    return { success: true, movies: movies };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}


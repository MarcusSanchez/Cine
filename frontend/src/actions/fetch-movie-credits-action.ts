"use server";

import { cookies } from "next/headers";
import API from "@/api/api";
import { MovieCredits } from "@/models/models";

type Result = {
  success: true, movieCredits: MovieCredits
} | {
  success: false, error: string
}

export default async function fetchMovieCreditsAction(ref: number): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    const { movie_credits } = await API.fetchMovieCredits(ref, cookie.value);

    return { success: true, movieCredits: movie_credits };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}
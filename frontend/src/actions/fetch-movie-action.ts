"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { DetailedMovie } from "@/models/models";

type Result = {
  success: true, detailedMovie: DetailedMovie
} | {
  success: false, error: string
}

export default async function fetchMovieAction(ref: number): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    const { detailed_movie } = await API.fetchMovie(ref, cookie.value);
    return { success: true, detailedMovie: detailed_movie };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}
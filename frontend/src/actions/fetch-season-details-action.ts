"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { DetailedSeason } from "@/models/models";

type Result = {
  success: true, detailedSeason: DetailedSeason
} | {
  success: false, error: string
}

export default async function fetchSeasonDetailsAction(ref: number, season: number): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    const { detailed_season } = await API.fetchShowSeasonDetails(ref, season, cookie.value);
    return { success: true, detailedSeason: detailed_season };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}

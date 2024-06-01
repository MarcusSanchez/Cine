"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { DetailedShow } from "@/models/models";

type Result = {
  success: true, detailedShow: DetailedShow
} | {
  success: false, error: string
}

export default async function fetchMovieAction(ref: number): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    const { detailed_show } = await API.fetchShow(ref, cookie.value);
    return { success: true, detailedShow: detailed_show };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}

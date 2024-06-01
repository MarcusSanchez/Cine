"use server";

import { cookies } from "next/headers";
import API from "@/api/api";
import { ShowCredits } from "@/models/models";

type Result = {
  success: true, showCredits: ShowCredits
} | {
  success: false, error: string
}

export default async function fetchShowCreditsAction(ref: number): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    const { show_credits } = await API.fetchShowCredits(ref, cookie.value);

    return { success: true, showCredits: show_credits };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}

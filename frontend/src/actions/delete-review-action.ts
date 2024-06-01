"use server";

import { cookies } from "next/headers";
import API from "@/api/api";

type Result = { success: true } | { success: false, error: string }

export default async function deleteReviewAction(csrf: string, reviewID: string): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    const result = await API.deleteReview(csrf, cookie.value, reviewID);
    if (!result.success) return { success: false, error: "failed to delete review" };

    return { success: true }
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}

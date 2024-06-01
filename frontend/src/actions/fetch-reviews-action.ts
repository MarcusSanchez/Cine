"use server";

import { cookies } from "next/headers";
import API from "@/api/api";
import { DetailedReview, MediaType } from "@/models/models";

type Result = {
  success: true, detailedReviews: DetailedReview[]
} | {
  success: false, error: string
}

export default async function fetchReviewsAction(media: MediaType, ref: number): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    const { detailed_reviews } = await API.fetchReviews(media, ref, cookie.value);

    return { success: true, detailedReviews: detailed_reviews };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}

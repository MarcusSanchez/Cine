"use server";

import { cookies } from "next/headers";
import API from "@/api/api";
import { FormSchema as ReviewFormSchema } from "@/components/UserReview";
import { MediaType, Review } from "@/models/models";

type Result = { success: true, review: Review } | { success: false, error: string }

export default async function editReviewAction(
  csrf: string,
  data: ReviewFormSchema,
  reviewID: string,
): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    const { review } = await API.editReview(csrf, cookie.value, data, reviewID);
    return { success: true, review }
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}


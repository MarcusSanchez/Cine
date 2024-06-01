"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { DetailedComment, MediaType } from "@/models/models";

type Result = {
  success: true, detailedComments: DetailedComment[]
} | {
  success: false, error: string
}

export default async function fetchCommentsAction(media: MediaType, ref: number): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    const { detailed_comments } = await API.fetchComments(cookie.value, media, ref);
    return { success: true, detailedComments: detailed_comments };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}

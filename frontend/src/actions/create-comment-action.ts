"use server";

import { cookies } from "next/headers";
import API from "@/api/api";
import { Comment, MediaType } from "@/models/models";

type Result = { success: true, comment: Comment } | { success: false, error: string }

export default async function createCommentAction(
  csrf: string,
  content: string,
  media: MediaType,
  ref: number,
  replyingTo?: string
): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    const { comment } = await API.createComment(csrf, cookie.value, content, media, ref, replyingTo);
    return { success: true, comment }
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}


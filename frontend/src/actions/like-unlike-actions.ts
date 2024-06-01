"use server";

import { cookies } from "next/headers";
import API from "@/api/api";

type Result = { success: true } | { success: false, error: string }

export async function likeCommentAction(
  csrf: string,
  commentID: string,
): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    await API.likeComment(csrf, cookie.value, commentID);
    return { success: true }
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}

export async function unlikeCommentAction(
  csrf: string,
  commentID: string,
): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    await API.unlikeComment(csrf, cookie.value, commentID);
    return { success: true }
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}




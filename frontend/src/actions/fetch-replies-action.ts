"use server";

import { cookies } from "next/headers";
import API from "@/api/api";
import { DetailedComment } from "@/models/models";

type Result = {
  success: true, replies: DetailedComment[]
} | {
  success: false, error: string
}

export default async function fetchRepliesAction(commentID: string): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    const { replies } = await API.fetchReplies(cookie.value, commentID);

    return { success: true, replies: replies };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}

"use server";

import API from "@/api/api";
import { cookies } from "next/headers";

type Result = {
  success: true
} | {
  success: false, error: string
}

export async function updateListAction(csrf: string, list: string, title?: string, isPublic?: boolean): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    await API.updateList(csrf, cookie.value, list, title, isPublic);
    return { success: true };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}


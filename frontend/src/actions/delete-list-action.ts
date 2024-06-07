"use server";

import API from "@/api/api";
import { cookies } from "next/headers";

type Result = {
  success: true
} | {
  success: false, error: string
}

export async function deleteListAction(csrf: string, list: string): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    await API.deleteList(csrf, cookie.value, list);
    return { success: true };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}


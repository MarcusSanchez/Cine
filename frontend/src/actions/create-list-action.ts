"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { List } from "@/models/models";

type Result = {
  success: true, list: List
} | {
  success: false, error: string
}

export async function createListAction(csrf: string, title: string): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    const { list } = await API.createList(csrf, cookie.value, title);
    return { success: true, list: list };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}


"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { DetailedList } from "@/models/models";

type Result = {
  success: true, list: DetailedList
} | {
  success: false, error: string
}

export async function fetchDetailedListAction(listID: string): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    const { list } = await API.fetchDetailedList(listID, cookie.value);
    return { success: true, list: list };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}


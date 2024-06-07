"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { DetailedList } from "@/models/models";

type Result = {
  success: true, lists: DetailedList[]
} | {
  success: false, error: string
}

export async function fetchMyListsAction(): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    const { detailed_lists } = await API.fetchMyLists(cookie.value);
    return { success: true, lists: detailed_lists };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}

export async function fetchUserListsAction(user: string): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    const { detailed_lists } = await API.fetchUserLists(user, cookie.value);
    return { success: true, lists: detailed_lists };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}

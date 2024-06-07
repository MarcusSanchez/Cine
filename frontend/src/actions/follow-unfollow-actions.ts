"use server";

import { cookies } from "next/headers";
import API from "@/api/api";

type Result = { success: true } | { success: false, error: string }

export async function followUserAction(
  csrf: string,
  followeeID: string,
): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    await API.followUser(csrf, cookie.value, followeeID);
    return { success: true }
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}

export async function unfollowUserAction(
  csrf: string,
  followeeID: string,
): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    await API.unfollowUser(csrf, cookie.value, followeeID);
    return { success: true }
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}


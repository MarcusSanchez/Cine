"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { UserStats } from "@/models/models";

type Result = {
  success: true, data: { stats: UserStats }
} | {
  success: false, error: string
};

export default async function fetchUserStatsAction(id: string): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session token not found" };

    const { stats } = await API.fetchUserStats(id, cookie.value);

    return { success: true, data: { stats: stats } };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}

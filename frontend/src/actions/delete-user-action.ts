"use server";

import { cookies } from "next/headers";
import API from "@/api/api";

type Result = { success: true } | { success: false, error: string }

export default async function deleteUserAction(csrf: string): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    const result = await API.deleteUser(csrf, cookie.value);
    if (!result.success) return { success: false, error: "failed to delete user" };

    cookies().delete("session");
    cookies().delete("csrf");

    return { success: true }
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}
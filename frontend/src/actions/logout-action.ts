"use server";

import { cookies } from "next/headers";
import API from "@/api/api";

type Result = {
  success: true
} | {
  success: false,
  error: string
}

export default async function logoutAction(csrf: string): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    await API.logout(csrf, cookie.value);

    cookies().delete("session");
    cookies().delete("csrf");

    return { success: true };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}
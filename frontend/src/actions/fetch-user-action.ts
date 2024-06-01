"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { User } from "@/models/models";

type Result = {
  success: true, data: { user: User }
} | {
  success: false, error: string
};

export default async function fetchUserAction(id: string): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session token not found" };

    const { user } = await API.fetchUser(id, cookie.value);

    return { success: true, data: { user: user } };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}
"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { maxAge } from "@/lib/utils";
import { User } from "@/models/models";

type Result = {
  success: true, data: { user: User }
} | {
  success: false, error: string
};

export default async function authenticateAction(csrf: string): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "Session token not found" };

    const { user, session } = await API.authenticate(cookie.value, csrf);
    const expiration = maxAge(session.expiration);

    cookies().set("session", session.token, { maxAge: expiration, sameSite: "lax" });
    cookies().set("csrf", session.csrf, { maxAge: expiration, httpOnly: false, sameSite: "lax" });

    return { success: true, data: { user } };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}
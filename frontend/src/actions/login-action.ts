"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { maxAge } from "@/lib/utils";

type Result = {
  success: true, data: { user: User }
} | {
  success: false, error: string
};

export async function loginAction(username: string, password: string): Promise<Result> {
  try {
    const { user, session } = await API.login(username, password);
    const expiration = maxAge(session.expiration);

    cookies().set("session", session.token, { maxAge: expiration, sameSite: "lax" });
    cookies().set("csrf", session.csrf, { maxAge: expiration, httpOnly: false, sameSite: "lax" });

    return { success: true, data: { user } };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}

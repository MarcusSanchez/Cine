"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { revalidatePath } from "next/cache";
import { User } from "@/models/models";

type Result = {
  data: { user: User }, success: true
} | {
  success: false, error: string
};

type Input = {
  username?: string,
  display_name?: string,
  password?: string,
  profile_picture?: string,
};

export default async function updateUserAction(csrf: string, input: Input): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session token not found" };

    const { user } = await API.updateUser(csrf, cookie.value, input);

    revalidatePath("/profile/me");
    return { success: true, data: { user: user } };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}

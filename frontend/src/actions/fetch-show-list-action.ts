"use server";

import API from "@/api/api";
import { cookies } from "next/headers";
import { Show, ShowList } from "@/models/models";

type Result = {
  success: true, shows: Show[]
} | {
  success: false, error: string
};

export default async function fetchShowListAction(list: ShowList): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session token not found" };

    const { shows } = await API.fetchShowList(list, cookie.value);

    return { success: true, shows: shows };
  } catch (e: any) {
    return { success: false, error: e.message };
  }
}

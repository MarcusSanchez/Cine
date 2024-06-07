"use server";

import API from "@/api/api";
import { cookies } from "next/headers";

type Result = {
  success: true
} | {
  success: false, error: string
}

export async function addMovieToListAction(csrf: string, list: string, ref: number): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    await API.addMovieToList(csrf, cookie.value, list, ref);
    return { success: true };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}

export async function removeMovieFromListAction(csrf: string, list: string, ref: number): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    await API.removeMovieFromList(csrf, cookie.value, list, ref);
    return { success: true };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}

export async function addShowToListAction(csrf: string, list: string, ref: number): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    await API.addShowToList(csrf, cookie.value, list, ref);
    return { success: true };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}

export async function removeShowFromListAction(csrf: string, list: string, ref: number): Promise<Result> {
  try {
    const cookie = cookies().get("session");
    if (!cookie) return { success: false, error: "session not found" }

    await API.removeShowFromList(csrf, cookie.value, list, ref);
    return { success: true };
  } catch (e: any) {
    return { success: false, error: e.message }
  }
}




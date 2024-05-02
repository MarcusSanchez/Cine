import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

// converts golang's string representation of time to a cookie's maxAge
export function maxAge(time: string) {
  const date = new Date(time);
  const now = new Date();
  const diffMS = date.getTime() - now.getTime();

  return Math.round(diffMS / 1000);
}

export function getCookie(name: string) {
  const cookies = document.cookie.split("; ");
  for (const cookie of cookies) {
    const [key, value] = cookie.split("=");
    if (key === name) return value;
  }
  return null;
}

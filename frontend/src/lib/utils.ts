import { type ClassValue, clsx } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

// converts a string representation of time to a human-readable time ago
export function timeAgo(input: string): string {
  const now = new Date();
  const past = new Date(input);
  const diffInSeconds = Math.floor((now.getTime() - past.getTime()) / 1000);

  const units = [
    { name: 'year', seconds: 31536000 },
    { name: 'month', seconds: 2592000 },
    { name: 'week', seconds: 604800 },
    { name: 'day', seconds: 86400 },
    { name: 'hour', seconds: 3600 },
    { name: 'minute', seconds: 60 },
  ];

  for (const unit of units) {
    const quotient = Math.floor(diffInSeconds / unit.seconds);
    if (quotient >= 1) return `${quotient} ${unit.name}${quotient > 1 ? 's' : ''} ago`;
  }

  return 'a few seconds ago';
}

// converts string representation of time to a cookie's maxAge
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

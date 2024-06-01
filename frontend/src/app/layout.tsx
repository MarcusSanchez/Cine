import type { Metadata } from "next";
import { Roboto } from "next/font/google";
import "./globals.css";

import "@/env/environment";
import API from "@/api/api";
import React from "react";
import { App } from "@/app/App";

API.url = process.env.API_URL;

const roboto = Roboto({
  weight: ["500"],
  subsets: ["latin"],
  display: "swap",
  variable: "--font-roboto",
});

export const metadata: Metadata = {
  title: "Cine | Cinema Social Platform",
  description: "A cinema social platform",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={roboto.className}>
      <body>
        <App>{children}</App>
      </body>
    </html>
  );
}

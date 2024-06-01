"use client";

import Navbar from "@/components/Navbar";
import Footer from "@/components/Footer";
import { useUserStore } from "@/app/state";
import { ReactNode, useEffect, useState } from "react";
import authenticateAction from "@/actions/authenticate-action";
import { getCookie } from "@/lib/utils";

export function App({ children }: Readonly<{ children: ReactNode }>) {
  const { setUser } = useUserStore();
  const [loaded, setLoaded] = useState(false);

  useEffect(() => {
    const authenticate = async () => {
      const csrf = getCookie("csrf");
      if (!csrf) return;

      const result = await authenticateAction(csrf);
      if (!result.success) {
        console.error(result.error);
        return;
      }

      setUser({ ...result.data.user, loggedIn: true, csrf: csrf });
    }

    authenticate().then(() => setLoaded(true));
  }, [])

  return (
    <>
      <Navbar />
      {loaded && children}
      <Footer />
    </>
  );
}
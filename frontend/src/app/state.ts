"use client";

import { atom, useAtom } from 'jotai';
import logoutAction from "@/actions/logout-action";
import { useRouter } from "next/navigation";
import { useToast } from "@/components/ui/use-toast";

const userAtom = atom({
  id: "",
  username: "",
  display_name: "",
  profile_picture: "",
  csrf: "",
  loggedIn: false,
});

export function useUserStore() {
  const [user, setUser] = useAtom(userAtom);
  const router = useRouter();
  const { toast } = useToast();

  const clearUser = () => setUser({
    id: "",
    username: "",
    display_name: "",
    profile_picture: "",
    csrf: "",
    loggedIn: false
  });

  const logout = async () => {
    const result = await logoutAction(user.csrf);
    if (!result.success) return toast({
      title: "Failed to logout",
      description: result.error,
    });

    clearUser();
    router.push("/");
  };

  return { user, setUser, logout, clearUser };
}



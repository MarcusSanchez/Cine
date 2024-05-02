import { atom, useAtom } from 'jotai';
import logoutAction from "@/actions/logout-action";
import { useRouter } from "next/navigation";

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

  const logout = async () => {
    const result = await logoutAction(user.csrf);
    if (!result.success) {
      console.error(result.error);
      return;
    }

    setUser({ id: "", username: "", display_name: "", profile_picture: "", csrf: "", loggedIn: false });
    router.push("/");
  }

  return { user, setUser, logout }
}



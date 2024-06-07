"use client";

import { useUserStore } from "@/app/state";
import React, { useEffect, useState } from "react";
import Link from "next/link";
import { cn, errorToast } from "@/lib/utils";
import { buttonVariants } from "@/components/ui/button";
import { useRouter } from "next/navigation";
import fetchUserAction from "@/actions/fetch-user-action";
import { User } from "@/models/models";
import { Plus, X } from "lucide-react";
import { useAtom } from "jotai";
import { followedAtom } from "@/app/profile/[user]/state";
import { followUserAction, unfollowUserAction } from "@/actions/follow-unfollow-actions";
import { useToast } from "@/components/ui/use-toast";

type ProfileLayoutProps = {
  children: React.ReactNode,
  params: { user: string }
}

export default function ProfileLayout({ children, params }: ProfileLayoutProps) {
  const { user } = useUserStore();
  const { toast } = useToast();

  const router = useRouter();

  let [currentUser, setCurrentUser] = useState<User>(user);

  const [followed, setFollowed] = useAtom(followedAtom);

  if (params.user === "me" && !user.loggedIn) router.replace("/login");
  if (params.user === user.id) router.replace("/profile/me")

  useEffect(() => {
    const fetchUser = async () => {
      const result = await fetchUserAction(params.user);
      if (!result.success) return router.replace("/404");
      setCurrentUser(result.data.user);
    }

    if (params.user === "me") return setCurrentUser(user);
    fetchUser();
  }, [user]);

  const followUser = async () => {
    const result = await followUserAction(user.csrf, currentUser.id);
    if (!result.success) return errorToast(toast, "Failed to follow user", result.error);
    setFollowed(true);
  }

  const unfollowUser = async () => {
    const result = await unfollowUserAction(user.csrf, currentUser.id);
    if (!result.success) return errorToast(toast, "Failed to unfollow user", result.error);
    setFollowed(false);
  }

  return (
    <div className="container max-w-[1200px]">
      <div className="grid grid-cols-1 sm:grid-cols-4 gap-4">
        <div className="pb-4 mb-4">
          <div className="flex flex-col justify-center content-center">
            <div className="flex m-auto md:m-0">
              <img
                src={currentUser.profile_picture}
                alt="avatar"
                className="rounded-full w-[150px] h-[150px] mb-4 ring-2 ring-brand-yellow"
              />
              {params.user !== "me" && followed &&
                <button className="self-end ml-[-35px]" onClick={unfollowUser}>
                  <X
                    className="
                      w-8 h-8 text-brand-darker border border-brand-darker bg-brand-yellow rounded-3xl mt-[-50px]
                      hover:bg-brand-light
                    "
                  />
                </button>
              }
              {params.user !== "me" && !followed &&
                <button className="self-end ml-[-35px]" onClick={followUser}>
                  <Plus
                    className="
                      w-8 h-8 text-brand-darker border border-brand-darker bg-brand-yellow rounded-3xl mt-[-50px]
                      hover:bg-brand-light
                    "
                  />
                </button>
              }
            </div>
            <h1 className="text-center md:text-left text-2xl text-brand-yellow font-bold">{currentUser.display_name}</h1>
            <p className="text-center md:text-left text-stone-400 text-md">@{currentUser.username}</p>
          </div>

          <div className="mt-8 flex sm:flex-col gap-2 border-y border-brand-yellow sm:border-none sm:border-t">
            <Link
              href={`/profile/${params.user}`}
              className={cn(
                buttonVariants({ variant: "link" }),
                "text-lg md:text-2xl text-brand-light hover:text-brand-yellow w-full sm:w-min px-0",
              )}
            >
              Profile
            </Link>
            <Link
              href={`/profile/${params.user}/lists`}
              className={cn(
                buttonVariants({ variant: "link" }),
                "text-lg md:text-2xl text-brand-light hover:text-brand-yellow w-full sm:w-min px-0",
              )}
            >
              Lists
            </Link>
          </div>
        </div>
        <div className="sm:col-span-3 sm:border-l-2 border-brand-yellow">
          {children}
        </div>
      </div>
    </div>
  );
}

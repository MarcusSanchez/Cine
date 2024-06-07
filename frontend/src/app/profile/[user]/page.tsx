"use client";

import { Handshake, Heart, List, MessageCircle, Star, Users } from "lucide-react";
import React, { ReactNode, useEffect, useState } from "react";
import fetchUserStatsAction from "@/actions/fetch-user-stats-action";
import DeleteProfileDialog from "@/app/profile/[user]/(components)/DeleteProfileDialog";
import EditProfileDialog from "@/app/profile/[user]/(components)/EditProfileDialog";
import { useAtom } from "jotai";
import { followedAtom } from "@/app/profile/[user]/state";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

export default function Profile({ params }: { params: { user: string } }) {
  const { toast } = useToast();

  const [followed, setFollowed] = useAtom(followedAtom);

  const [likes, setLikes] = useState(0);
  const [comments, setComments] = useState(0);
  const [reviews, setReviews] = useState(0);
  const [lists, setLists] = useState(0);
  const [following, setFollowing] = useState(0);
  const [followers, setFollowers] = useState(0);

  useEffect(() => {
    const fetchStats = async () => {
      const result = await fetchUserStatsAction(params.user);
      if (!result.success) return errorToast(toast, "Failed to fetch user stats", "Please try again later");

      const { stats } = result.data;

      setLikes(stats.likes_count);
      setComments(stats.comments_count);
      setReviews(stats.reviews_count);
      setLists(stats.lists_count);
      setFollowing(stats.following_count);
      setFollowed(stats.followed);
      setFollowers(stats.followers_count);

      // due to race condition, we adjust the count based on the followed status
      if (stats.followers_count === 0) return;
      if (stats.followed) setFollowers(stats.followers_count - 1);
    }

    fetchStats();
  }, []);

  useEffect(() => {
    setFollowers(followed ? followers + 1 : followers - 1);
  }, [followed]);

  return (
    <div className="container mb-8">
      {params.user === "me" ?
        (
          <div className="flex justify-between border-b border-brand-yellow pb-3 mb-4">
            <h1 className="text-xl sm:text-2xl md:text-4xl font-bold text-brand-light">My Profile</h1>
            <div className="flex gap-2">
              <EditProfileDialog />
              <DeleteProfileDialog />
            </div>
          </div>
        ) :
        <h1 className="text-4xl font-bold text-brand-light mb-4">Profile</h1>
      }

      <p className="text-stone-400 text-lg mb-4">Stats</p>
      <div className=" grid grid-cols-2 gap-4">
        <StatCard value={following} label="Following" Icon={<Handshake />} />
        <StatCard value={followers} label="Followers" Icon={<Users />} />
        <StatCard value={likes} label="Likes Given" Icon={<Heart />} />
        <StatCard value={comments} label="Comments" Icon={<MessageCircle />} />
        <StatCard value={reviews} label="Reviews" Icon={<Star />} />
        <StatCard value={lists} label="Lists" Icon={<List />} />
      </div>
    </div>
  );
}

type StatCardProps = {
  value: number,
  label: string,
  Icon: ReactNode,
}

function StatCard({ value, label, Icon }: StatCardProps) {
  return (
    <div className="bg-brand-darker rounded-lg shadow-md p-6 border border-brand-yellow">
      <div className="flex items-center justify-between mb-4">
        <h1 className="text-4xl font-bold text-brand-light">{value}</h1>
        <span className="text-brand-light">{Icon}</span>
      </div>
      <p className="text-stone-400">{label}</p>
    </div>
  )
}


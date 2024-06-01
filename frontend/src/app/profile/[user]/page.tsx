"use client";

import { Handshake, Heart, List, MessageCircle, Star, Users } from "lucide-react";
import React, { ReactNode, useEffect, useState } from "react";
import fetchUserStatsAction from "@/actions/fetch-user-stats-action";
import DeleteProfileDialog from "@/app/profile/[user]/(components)/DeleteProfileDialog";
import EditProfileDialog from "@/app/profile/[user]/(components)/EditProfileDialog";

export default function Profile({ params }: { params: { user: string } }) {
  const [likes, setLikes] = useState(0);
  const [comments, setComments] = useState(0);
  const [reviews, setReviews] = useState(0);
  const [lists, setLists] = useState(0);
  const [following, setFollowing] = useState(0);
  const [followers, setFollowers] = useState(0);

  useEffect(() => {
    const fetchStats = async () => {
      const result = await fetchUserStatsAction(params.user);
      if (!result.success) return;

      const { stats } = result.data;

      setLikes(stats.likes_count);
      setComments(stats.comments_count);
      setReviews(stats.reviews_count);
      setLists(stats.lists_count);
      setFollowing(stats.following_count);
      setFollowers(stats.followers_count);
    }

    fetchStats();
  }, []);

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


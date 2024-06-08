"use client";

import React, { useEffect, useState } from "react";
import { MediaType, Movie, MovieList, Show, ShowList } from "@/models/models";
import fetchMovieListAction from "@/actions/fetch-movie-list-action";
import fetchShowListAction from "@/actions/fetch-show-list-action";
import List, { ListsSkeleton } from "@/components/List";
import { useUserStore } from "@/app/state";
import { useRouter } from "next/navigation";
import mascot from "@/../public/mascot.png";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

type ListType = [MediaType.Movie, MovieList] | [MediaType.Show, ShowList];

const listOrder: ListType[] = [
  [MediaType.Movie, MovieList.Upcoming],
  [MediaType.Show, ShowList.Popular],
  [MediaType.Movie, MovieList.Popular],
  [MediaType.Show, ShowList.TopRated],
  [MediaType.Movie, MovieList.TopRated],
];

export default function Home() {
  const { user } = useUserStore();
  const { toast } = useToast();

  const router = useRouter();

  const [lists, setLists] = useState<([MediaType.Movie, MovieList, Movie[]] | [MediaType.Show, ShowList, Show[]])[]>([]);

  if (!user.loggedIn) router.replace("/login");

  useEffect(() => {
    const fetchLists = async () => {
      for (const [mediaType, list] of listOrder) {
        switch (mediaType) {
          case MediaType.Movie:
            fetchMovieListAction(list).then((result) => {
              if (!result.success) return errorToast(toast, "Failed to fetch content", "Please try again later.");
              setLists((prev) => [...prev, [mediaType, list, result.movies]]);
            });
            break;
          case MediaType.Show:
            fetchShowListAction(list).then((result) => {
              if (!result.success) return errorToast(toast, "Failed to fetch content", "Please try again later.");
              setLists((prev) => [...prev, [mediaType, list, result.shows]]);
            });
            break;
        }
      }
    }

    if (!user.loggedIn) return;
    fetchLists();
  }, []);

  return (
    <div className="container max-w-[1200px] mb-8">
      <Opener />
      {lists.length < 1 && <ListsSkeleton />}
      {lists.map(([mediaType, list, items]) => (
        <List key={mediaType + list} mediaType={mediaType} list={list} items={items} />
      ))}
    </div>
  );
}

const Opener = () => (
  <div className="flex flex-col md:grid md:grid-cols-5 ">
    <div className="order-last text-center md:order-first col-span-3 flex flex-col justify-center">
      <h1 className="text-4xl md:text-6xl lg:text-8xl font-bold text-brand-yellow">Welcome.</h1>
      <p className="text-sm sm:text-base md:text-xl lg:text-2xl text-brand-light">
        With access to millions of Movies and TV Shows, you're cinematic cravings will never perish.
        Share your thoughts, reviews and lists with the community!
      </p>
    </div>
    <div className="w-full flex justify-center col-span-2">
      <img src={mascot.src} alt="cinema-mascot" className="max-h-[10rem] md:max-h-[20rem] lg:max-h-[24rem]" />
    </div>
  </div>
);


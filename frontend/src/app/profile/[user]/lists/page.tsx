"use client";

import React, { useEffect, useState } from "react";
import { fetchMyListsAction, fetchUserListsAction } from "@/actions/fetch-lists-actions";
import { DetailedList, type List, Media, MediaType } from "@/models/models";
import CreateListDialog from "@/app/profile/[user]/lists/(components)/CreateListDialog";
import Link from "next/link";
import { Carousel, CarouselContent, CarouselItem } from "@/components/ui/carousel";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

export default function ProfileLists({ params }: { params: { user: string } }) {
  const [lists, setLists] = useState<DetailedList[]>([]);
  const { toast } = useToast();

  useEffect(() => {
    const fetchLists = async () => {
      switch (params.user) {
        case "me":
          const myListsResult = await fetchMyListsAction();
          if (!myListsResult.success) return errorToast(toast, "Failed to fetch lists", "Please try again later");
          setLists(myListsResult.lists);
          break;
        default:
          const listsResult = await fetchUserListsAction(params.user);
          if (!listsResult.success) return;
          setLists(listsResult.lists);
          break;
      }
    }

    fetchLists();
  }, []);

  return (
    <div className="container max-w-[1200px]">
      <div className="flex justify-between border-b border-brand-yellow pb-3 mb-4">
        <h1 className="text-xl sm:text-2xl md:text-4xl font-bold text-brand-light">Lists</h1>
        {params.user === "me" && <CreateListDialog setLists={setLists} />}
      </div>
      <div>
        {lists.length < 1 && <p className="text-brand-light">No lists found</p>}
        {lists.map(l =>
          <List
            params={params}
            list={l.list}
            medias={[...l.shows ?? [], ...l.movies ?? []]}
          />
        )}
      </div>
    </div>
  );
}

type ListProps = {
  list: List;
  medias: Media[];
  params: { user: string };
};

const List = ({ list, medias, params }: ListProps) => (
  <div>
    <Link href={`/profile/${params.user}/lists/${list.id}`} key={list.id}>
      <h2 className="text-xl md:text-2xl font-bold text-brand-yellow mt-4 mb-4 hover:text-brand-light">
        {list.name}
      </h2>
    </Link>
    {medias.length < 1 && <p className="text-brand-light">No items found</p>}
    <Carousel
      opts={{
        align: "start",
      }}
    >
      <CarouselContent className="w-[95%]">
        {medias.map((media, i) => (
          <CarouselItem key={i} className="basis-1/3 md:basis-1/4 lg:basis-1/5 h-min">
            <ItemCard media={media} mediaType={media.media_type} />
          </CarouselItem>
        ))}
      </CarouselContent>
    </Carousel>
  </div>
)

const imageBase = "https://image.tmdb.org/t/p/w500";

const ItemCard = ({ media, mediaType }: { media: Media, mediaType: MediaType }) => (
  <div className="flex flex-col group">
    <Link href={`/${mediaType}s/${media.ref}`}>
      <img
        src={`${imageBase}${media.poster_path}`}
        alt={media.title}
        className="rounded-lg w-full object-cover group-hover:opacity-60 transition-opacity duration-200 ease-in-out"
      />
    </Link>
    <Link href={`/${mediaType}s/${media.ref}`}>
      <p className="text-sm sm:text-base md:text-lg text-brand-light mt-2 group-hover:text-brand-yellow">
        {media.title}
      </p>
    </Link>
    <p className="text-sm md:text-base text-stone-400">
      {media.release_date?.slice(0, 4)}
    </p>
  </div>
);

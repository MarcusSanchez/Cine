"use client";

import React, { useEffect, useState } from "react";
import { DetailedList, Media, MediaType } from "@/models/models";
import { fetchDetailedListAction } from "@/actions/fetch-detailed-list-action";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { X } from "lucide-react";
import { removeMovieFromListAction, removeShowFromListAction } from "@/actions/add-remove-media-from-list-action";
import { useUserStore } from "@/app/state";
import EditListDialog from "@/app/profile/[user]/lists/[listID]/(components)/EditListDialog";
import DeleteListDialog from "@/app/profile/[user]/lists/[listID]/(components)/DeleteListDialog";
import { useToast } from "@/components/ui/use-toast";
import { useRouter } from "next/navigation";
import { errorToast } from "@/lib/utils";

export default function ListPage({ params }: { params: { user: string, listID: string } }) {
  const { user } = useUserStore();
  const { toast } = useToast();

  const router = useRouter();

  const [list, setList] = useState<DetailedList | null>(null);
  const medias = [...(list?.movies ?? []), ...(list?.shows ?? [])];

  useEffect(() => {
    const fetchList = async () => {
      const result = await fetchDetailedListAction(params.listID);
      if (!result.success) return router.replace("/uh-oh");
      setList(result.list);
    };

    fetchList();
  }, []);

  const removeMedia = async (media: Media) => {
    switch (media.media_type) {
      case MediaType.Movie:
        const movieResult = await removeMovieFromListAction(user.csrf, params.listID, media.ref);
        if (!movieResult.success) return errorToast(toast, "Failed to remove movie from list", movieResult.error);
        break;
      case MediaType.Show:
        const showResult = await removeShowFromListAction(user.csrf, params.listID, media.ref);
        if (!showResult.success) return errorToast(toast, "Failed to remove show from list", showResult.error);
        break;
    }

    setList({
      ...list!,
      movies: list!.movies.filter(m => m.ref !== media.ref),
      shows: list!.shows.filter(m => m.ref !== media.ref),
    });
  };

  return (
    <div className="container max-w-[1200px]">
      <div className="justify-between border-b border-brand-yellow pb-3 mb-4">
        <h1 className="text-3xl font-bold text-brand-yellow">{list?.list.name}</h1>
      </div>

      {params.user === "me" && list &&
        <div className="flex gap-2 mb-2">
          <EditListDialog list={list} setList={setList} />
          <DeleteListDialog list={list} />
        </div>
      }

      <p className="text-brand-light">{medias.length} items</p>
      {medias.length < 1 && <p className="text-brand-light">No items found. Try adding your favorite movie or show.</p>}
      {medias.map(media =>
        <MediaCard key={media.ref} media={media} params={params} removeMedia={removeMedia} />
      )}
    </div>
  );
}

const baseImageURL = "https://image.tmdb.org/t/p/w500";
const posterErrorImage = "https://conversionfanatics.com/wp-content/themes/seolounge/images/no-image/No-Image-Found-400x264.png";

type MediaCardProps = {
  media: Media,
  params: { user: string },
  removeMedia: (M: Media) => void,
};

const MediaCard = ({ media, params, removeMedia }: MediaCardProps) => (
  <div className="flex gap-2 bg-brand-darker border border-brand-yellow rounded-xl p-2 mb-2 hover:bg-brand-dark">
    <Link href={`/movies/${media.ref}`}>
      <img
        src={`${baseImageURL}${media.poster_path}`}
        alt={media.title}
        onError={(e) => (e.target as HTMLImageElement).src = posterErrorImage}
        className="w-32 lg:w-48 object-cover rounded-xl hover:opacity-60 transition duration-100 ease-in-out"
      />
    </Link>
    <div className="flex flex-col justify-between w-full">
      <div>
        <div className="flex">
          <Link href={`/movies/${media.ref}`} className="w-full group">
            <h2 className="text-brand-yellow font-bold text-lg md:text-2xl group-hover:text-brand-light">{media.title}</h2>
          </Link>
          {params.user === "me" &&
            <Button
              variant="link"
              size="sm"
              className="text-brand-yellow hover:text-brand-light h-min w-min p-0"
              onClick={() => removeMedia(media)}
            >
              <X className="h-8 w-6" />
            </Button>
          }
        </div>
        <p className="text-sm md:text-base text-brand-light">{media.overview.slice(0, media.overview.indexOf(".") + 1)}</p>
      </div>
      <div className="flex justify-between">
        <p className="text-base md:text-lg text-stone-400 self-end">{media.media_type === MediaType.Movie ? "Movie" : "Show"}</p>
        <p className="text-base md:text-lg text-stone-400 self-end">{media.release_date?.slice(0, 4) || "unreleased"}</p>
      </div>
    </div>
  </div>
);
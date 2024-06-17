"use client";

import { DetailedSeason, DetailedShow, Episode, Season } from "@/models/models";
import React, { Dispatch, SetStateAction, useEffect, useState } from "react";
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import fetchSeasonDetailsAction from "@/actions/fetch-season-details-action";
import { Carousel, CarouselContent, CarouselItem } from "@/components/ui/carousel";
import { Skeleton } from "@/components/ui/skeleton";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";
import { backdropBase, episodeErrorImage } from "@/lib/constants";

export default function ShowSeasons({ show }: { show: DetailedShow | null }) {
  const seasons = show?.seasons.filter(s => s.name !== "Specials") ?? [];
  const [seasonName, setSeasonName] = useState<string>("");
  const season = seasons.find(s => s.name === seasonName);

  useEffect(() => {
    setSeasonName(seasons[0]?.name ?? "")
  }, [show]);

  return (
    <div className="container max-w-[1200px] mt-8">
      <h1 className="text-2xl md:text-4xl mb-2 text-brand-yellow font-bold">Seasons</h1>
      <ShowSeasonSelect
        setSeason={setSeasonName}
        seasons={seasons.map(s => s.name)}
        selected={seasonName}
      />
      <ShowSeasonContent show={show} season={season} />
    </div>
  );
}

const fetchedSeasons = new Map<[number, number], DetailedSeason>();

function ShowSeasonContent({ show, season }: { show: DetailedShow | null, season?: Season }) {
  const { toast } = useToast();

  const [content, setContent] = useState<DetailedSeason | null>(null);

  useEffect(() => {
    const fetchSeason = async () => {
      const res = await fetchSeasonDetailsAction(show!.id, season!.season_number);
      if (!res.success) return errorToast(toast, "Failed to fetch season details", "Please try again later");

      fetchedSeasons.set([show!.id, season!.season_number], res.detailedSeason);
      setContent(res.detailedSeason);
    }

    if (!show || !season) return;
    const fetched = fetchedSeasons.get([show.id, season.season_number]);
    if (fetched) return setContent(fetched);

    fetchSeason();
  }, [show, season])

  if (!season) return <ShowSeasonContentSkeleton />;

  return (
    <div>
      <p className="text-sm md:text-base text-brand-light mb-4">{season.overview}</p>
      <Carousel
        opts={{
          align: "start",
        }}
      >
        <CarouselContent className="w-[90%]">
          {content?.episodes.map((episode) => (
            <CarouselItem key={episode.name} className="basis-[100%] md:basis-1/3 lg:basis-1/4 h-min">
              <EpisodeCard episode={episode} />
            </CarouselItem>
          ))}
        </CarouselContent>
      </Carousel>
    </div>
  );
}

const EpisodeCard = ({ episode }: { episode: Episode }) => (
  <div className="flex flex-col">
    <img
      className="rounded-xl border border-brand-yellow mb-2"
      src={`${backdropBase}${episode.still_path}`}
      alt={episode.name}
      onError={(e) => (e.target as HTMLImageElement).src = episodeErrorImage}
    />
    <div>
      <h3 className="text-brand-yellow">{episode.episode_number}. {episode.name}</h3>
      <p className="text-sm text-brand-light">{episode.overview.slice(0, episode.overview.indexOf(".") + 1)}</p>
    </div>
  </div>
);

type ShowSeasonSelectProps = {
  seasons: string[],
  selected: string,
  setSeason: Dispatch<SetStateAction<string>>
}

const ShowSeasonSelect = ({ seasons, selected, setSeason }: ShowSeasonSelectProps) => (
  <Select defaultValue={selected} onValueChange={v => setSeason(v)}>
    <SelectTrigger className="w-[180px] mb-4 text-brand-yellow bg-brand-darker border-brand-yellow">
      <SelectValue placeholder={selected} />
    </SelectTrigger>
    <SelectContent className="bg-brand-darker border-brand-yellow">
      <SelectGroup>
        {seasons.map((season) => (
          <SelectItem
            key={season}
            value={season}
            className="hover:cursor-pointer focus:bg-brand-dark text-brand-yellow focus:text-brand-yellow"
          >
            {season}
          </SelectItem>
        ))}
      </SelectGroup>
    </SelectContent>
  </Select>
);

const ShowSeasonContentSkeleton = () => (
  <div>
    <Skeleton className="w-full h-4 mb-1" />
    <Skeleton className="w-3/4 h-4 mb-4" />
    <Carousel
      opts={{
        align: "start",
      }}
    >
      <CarouselContent className="w-[90%]">
        {Array.from({ length: 8 }).map((_, i) => (
          <CarouselItem key={i} className="basis-[100%] md:basis-1/3 lg:basis-1/4 h-min">
            <div className="flex flex-col">
              <Skeleton className="w-full h-28 rounded-xl border border-brand-yellow mb-2" />
              <div>
                <Skeleton className="w-3/4 h-4 mb-2" />
                <Skeleton className="w-full h-2 mb-1" />
                <Skeleton className="w-full h-2 mb-1" />
                <Skeleton className="w-3/4 h-2" />
              </div>
            </div>
          </CarouselItem>
        ))}
      </CarouselContent>
    </Carousel>
  </div>
);
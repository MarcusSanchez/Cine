"use client";

import React, { useEffect, useState } from "react";
import fetchMovieCreditsAction from "@/actions/fetch-movie-credits-action";
import { Carousel, CarouselContent, CarouselItem, } from "@/components/ui/carousel"
import { Button } from "@/components/ui/button";
import { Skeleton } from "@/components/ui/skeleton";
import type { Cast, Crew, MovieCredits } from "@/models/models";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

const defaultCast = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png";
const castBaseURL = "https://image.tmdb.org/t/p/w500";

export default function MovieCredits({ refID }: { refID: number }) {
  const { toast } = useToast();

  const [credits, setCredits] = useState<MovieCredits | null>(null);
  const [viewCrew, setViewCrew] = useState(false);

  useEffect(() => {
    const fetchCredits = async () => {
      const result = await fetchMovieCreditsAction(refID);
      if (!result.success) return errorToast(toast, "Failed to fetch movie credits", "Please try again later.");
      setCredits(result.movieCredits);
    }

    fetchCredits();
  }, []);

  // Filter and sort cast and crew members
  const cast = credits?.cast.filter((member) => member.profile_path).slice(0, 20) ?? [];
  const crew = credits?.crew.filter((member) => member.profile_path) ?? [];
  const director = crew.find((member) => member.job === "Director");
  const screenplay = crew.find((member) => member.job === "Screenplay");
  const executiveProducer = crew.find((member) => member.job === "Executive Producer");
  const otherCrew = crew.filter((member) => !(["Director", "Screenplay", "Executive Producer"].includes(member.job)));
  const sortedCrew = [director, screenplay, executiveProducer, ...otherCrew].filter(m => m).slice(0, 20) as Crew[];

  if (!credits) return <MovieCreditsSkeleton />;

  return (
    <div className="container max-w-[1200px] mt-8">
      <CreditSection title="Cast" members={cast} />
      <Button variant="link" onClick={() => setViewCrew(!viewCrew)} className="text-brand-yellow p-0">
        {!viewCrew ? "View Crew" : "Hide Crew"}
      </Button>
      <div className={!viewCrew ? "hidden" : ""}>
        <CreditSection title="Crew" members={sortedCrew} />
      </div>
    </div>
  );
}

const CreditSection = ({ title, members }: { title: string, members: (Cast | Crew)[] }) => (
  <div>
    <h1 className="text-2xl md:text-4xl mb-4 text-brand-yellow font-bold">{title}</h1>
    <Carousel
      opts={{
        align: "start",
      }}
    >
      <CarouselContent className="w-[96%]">
        {members.map((member, i) => (
          <CarouselItem key={member.name + i} className="basis-1/3 sm:basis-1/4 md:basis-1/5 lg:basis-1/6 xl:basis-[12.5%] h-min">
            <CreditCard member={member} />
          </CarouselItem>
        ))}
      </CarouselContent>
    </Carousel>
  </div>
);

const CreditCard = ({ member }: { member: Cast | Crew }) => (
  <div className="flex flex-col">
    <img
      src={member.profile_path ? `${castBaseURL}/${member.profile_path}` : defaultCast}
      alt={member.name}
      className="w-auto rounded-xl max-h-[150px] sm:max-h-[250px] md:max-h-[400px] border border-brand-yellow"
    />
    <div className="mt-2 flex flex-col">
      <span className="text-xs md:text-base text-brand-yellow font-bold">{member.name}</span>
      <span className="text-xs md:text-base text-brand-light font-normal">{('character' in member) ? member.character : member.job}</span>
    </div>
  </div>
);

const MovieCreditsSkeleton = () => (
  <>
    <div className="container max-w-[1200px] mt-12">
      <h1 className="text-2xl md:text-4xl mb-4 text-brand-yellow font-bold">Cast</h1>
      <div className="mx-4 flex justify-center content-center h-min">
        <Carousel
          opts={{
            align: "start",
          }}
          className="w-[90%] md:w-[95%] h-min"
        >
          <CarouselContent>
            {Array.from({ length: 20 }).map((_, i) => (
              <CarouselItem key={i} className="basis-1/3 sm:basis-1/4 md:basis-1/5 lg:basis-1/6 xl:basis-[12.5%] h-min">
                <div className="flex flex-col">
                  <Skeleton className="w-[80px] sm:w-[120px] rounded-xl h-[120px] md:h-[180px]" />
                  <div className="mt-2 flex flex-col">
                    <Skeleton className="w-[60%] h-4" />
                    <Skeleton className="w-[65%] h-4 mt-1" />
                  </div>
                </div>
              </CarouselItem>
            ))}
          </CarouselContent>
        </Carousel>
      </div>
    </div>
  </>
);
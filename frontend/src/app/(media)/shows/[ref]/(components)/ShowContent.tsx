"use client";

import React from "react";
import { Skeleton } from "@/components/ui/skeleton";
import { DetailedShow } from "@/models/models";
import Image from "next/image";

const posterBaseURL = "https://image.tmdb.org/t/p/original";
const backdropBaseURL = "https://image.tmdb.org/t/p/original";

const ShowContent = ({ show }: { show: DetailedShow | null }) => (
  !show ? <ShowContentSkeleton /> : (
    <div
      className="w-auto h-fit bg-cover bg-center bg-no-repeat mt-[-2rem] border-y border-brand-yellow"
      style={{ backgroundImage: `url(${backdropBaseURL}${show?.backdrop_path})` }}
    >
      <div className="inset-0 bg-black bg-opacity-85 backdrop-filter pt-8 pb-12">
        <div className="container max-w-[1200px]">
          <div className="grid grid-cols-1 md:grid-cols-5 gap-4">
            <div className="flex justify-center col-span-2 xl:col-span-1">
              <img
                src={`${posterBaseURL}${show?.poster_path}`}
                alt={show?.name}
                className="w-auto rounded-xl max-h-[250px] md:max-h-[344px]"
              />
            </div>
            <div className="col-span-3 xl:col-span-4">
              <h1 className="text-4xl font-bold text-white flex flex-col">
                <div>
                  {show?.name} {" "}
                  <span className="text-brand-light font-normal">({show?.first_air_date?.slice(0, 4)})</span>
                </div>
                <span className="text-brand-yellow font-normal text-sm">{show?.tagline}</span>
              </h1>
              <p className="text-brand-light md:text-lg mt-2">{show?.overview}</p>
              <div className="mt-4">
                <span className="text-white">Genres: </span>
                {show?.genres.map((genre, index) => (
                  <span key={index} className="text-brand-light">{genre.name}{index < show.genres.length - 1 ? ", " : ""}</span>
                ))}
              </div>
              <div className="mt-4">
                <span className="text-white">Seasons: </span>
                <span className="text-brand-light">{show?.number_of_seasons} </span>
              </div>
              <div className="mt-4">
                <span className="text-white">First Aired: </span>
                <span className="text-brand-light">{show?.first_air_date}</span>
              </div>
              <div className="mt-4">
                <span className="text-white">Status: </span>
                <span className="text-brand-light">{show?.status}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
);

const ShowContentSkeleton = () => (
  <div
    className="w-auto h-fit bg-cover bg-no-repeat mt-[-2rem] border-y border-brand-yellow"
  >
    <div className="inset-0 bg-black bg-opacity-85 backdrop-filter pt-8 pb-12">
      <div className="container max-w-[1200px]">
        <div className="grid grid-cols-1 md:grid-cols-5 gap-4">
          <div className="flex justify-center col-span-2 xl:col-span-1">
            <Skeleton className="w-[155px] md:w-[214px] rounded-xl h-[250px] md:h-[344px]" />
          </div>
          <div className="col-span-3 xl:col-span-4">
            <h1 className="text-4xl font-bold text-white flex flex-col">
              <div>
                <Skeleton className="w-[80%] h-12" />
              </div>

              <Skeleton className="w-[40%] h-6 mt-2" />
            </h1>

            <Skeleton className="w-[100%] h-6 mt-4" />
            <Skeleton className="w-[100%] h-6 mt-2" />

            <div className="mt-4">
              <Skeleton className="w-[30%] md:w-[25%] h-6" />
            </div>
            <div className="mt-4">
              <Skeleton className="w-[25%] md:w-[20%] h-6" />
            </div>
            <div className="mt-4">
              <Skeleton className="w-[30%] md:w-[25%] h-6" />
            </div>
            <div className="mt-4">
              <Skeleton className="w-[25%] md:w-[20%] h-6" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
);

export default ShowContent;

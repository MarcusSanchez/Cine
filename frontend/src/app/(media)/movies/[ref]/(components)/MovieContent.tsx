"use client";

import React from "react";
import { Skeleton } from "@/components/ui/skeleton";
import { DetailedMovie, MediaType } from "@/models/models";
import AddToListDialog from "@/components/AddToListDialog";
import { backdropBase, imageBase } from "@/lib/constants";

const MovieContent = ({ movie }: { movie: DetailedMovie | null }) => (
  !movie ? <MainMovieContentSkeleton /> : (
    <div
      className="w-auto h-fit bg-cover bg-center bg-no-repeat mt-[-2rem] border-y border-brand-yellow"
      style={{ backgroundImage: `url(${backdropBase}${movie?.backdrop_path})` }}
    >
      <div className="inset-0 bg-black bg-opacity-85 backdrop-filter pt-8 pb-12">
        <div className="container max-w-[1200px]">
          <div className="grid grid-cols-1 md:grid-cols-5 gap-4">
            <div className="flex justify-center col-span-2 xl:col-span-1">
              <img
                src={`${imageBase}${movie?.poster_path}`}
                alt={movie?.title}
                className="w-auto rounded-xl max-h-[250px] md:max-h-[344px]"
              />
            </div>
            <div className="col-span-3 xl:col-span-4">
              <h1 className="text-4xl font-bold text-white flex flex-col">
                <div>
                  {movie?.title} {" "}
                  <span className="text-brand-light font-normal">({movie?.release_date?.slice(0, 4)})</span>
                </div>
                <span className="text-brand-yellow font-normal text-sm">{movie?.tagline}</span>
              </h1>
              <p className="text-brand-light md:text-lg mt-2">{movie?.overview}</p>
              <div className="mt-4">
                <span className="text-white">Genres: </span>
                {movie?.genres.map((genre, index) => (
                  <span key={index} className="text-brand-light">{genre.name}{index < movie.genres.length - 1 ? ", " : ""}</span>
                ))}
              </div>
              <div className="mt-4">
                <span className="text-white">Runtime: </span>
                <span className="text-brand-light">{min2hrs(movie?.runtime)} </span>
              </div>
              <div className="mt-4">
                <span className="text-white">Release Date: </span>
                <span className="text-brand-light">{movie?.release_date}</span>
              </div>
              <div className="mt-4 mb-4">
                <span className="text-white">Status: </span>
                <span className="text-brand-light">{movie?.status}</span>
              </div>
              <AddToListDialog mediaType={MediaType.Movie} refID={movie.id} />
            </div>
          </div>
        </div>
      </div>
    </div>
  )
);

function min2hrs(min: number = 0) {
  const hours = Math.floor(min / 60);
  const minutes = min % 60;
  return `${hours}h ${minutes}m`;
}

const MainMovieContentSkeleton = () => (
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

export default MovieContent;

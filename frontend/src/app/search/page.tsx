"use client";

import { useRouter, useSearchParams } from "next/navigation";
import React, { useEffect, useState } from "react";
import { Movie, Show } from "@/models/models";
import { searchMoviesAction, searchShowsAction } from "@/actions/search-actions";
import { MovieCard, ShowCard } from "@/components/Cards";
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";
import { Simulate } from "react-dom/test-utils";

export default function SearchPage() {
  const { toast } = useToast();

  const router = useRouter();
  const params = useSearchParams();

  const [movies, setMovies] = useState<Movie[]>([]);
  const [shows, setShows] = useState<Show[]>([]);

  const q = params.get("q");
  const filter = params.get("filter") ?? "movies";

  if (!q) router.replace("/");

  useEffect(() => {
    const search = async () => {
      switch (filter) {
        case "movies":
          const movieResult = await searchMoviesAction(q!);
          if (!movieResult.success) return errorToast(toast, "Failed to search movies", "Please try again later");
          setMovies(movieResult.movies);
          break;
        case "shows":
          const showResult = await searchShowsAction(q!);
          if (!showResult.success) return errorToast(toast, "Failed to search shows", "Please try again later");
          setShows(showResult.shows);
          break;
      }
    }

    if (!q) return;
    search();
  }, [q, filter]);

  const setFilter = (filter: string) => router.push(`/search?q=${q}&filter=${filter}`);

  return (
    <div className="container max-w-[1200px]">
      <h1 className="text-3xl font-bold text-brand-yellow">Search results for "{q}"</h1>
      <p className="text-brand-light">{filter === "movies" ? movies.length : shows.length} results</p>
      <SearchFilterSelect selected={filter} setFilter={setFilter} />
      {filter === "movies"
        ? movies.length > 0
          ? movies.map(movie =>
            <MovieCard key={movie.id} movie={movie} />
          )
          : <p className="text-white">No movies found</p>

        : shows.length > 0
          ? shows.map(show =>
            <ShowCard key={show.id} show={show} />
          )
          : <p className="text-white">No shows found</p>
      }
    </div>
  );
}

const SearchFilterSelect = ({ selected, setFilter }: { selected: string, setFilter: (s: string) => void }) => (
  <Select defaultValue={selected} onValueChange={v => setFilter(v)}>
    <SelectTrigger className="h-10 text-brand-yellow bg-brand-darker border-brand-yellow rounded mt-2 mb-4">
      <SelectValue placeholder={selected} />
    </SelectTrigger>
    <SelectContent className="bg-brand-darker border-brand-yellow">
      <SelectGroup>
        {["Movies", "Shows"].map((filter) => (
          <SelectItem
            key={filter}
            value={filter.toLowerCase()}
            className="hover:cursor-pointer focus:bg-brand-dark text-brand-yellow focus:text-brand-yellow"
          >
            {filter}
          </SelectItem>
        ))}
      </SelectGroup>
    </SelectContent>
  </Select>
)

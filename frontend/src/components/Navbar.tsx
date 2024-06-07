import Logo from '@/../public/logo.png';
import { Button, buttonVariants } from "@/components/ui/button";
import Link from 'next/link';
import { ListChecks, LogOut, Search, User } from "lucide-react";
import { cn, errorToast } from "@/lib/utils";
import { useUserStore } from "@/app/state";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import React, { FormEvent, useEffect, useRef, useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger
} from "@/components/ui/dialog";
import { DialogBody } from "next/dist/client/components/react-dev-overlay/internal/components/Dialog";
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { Movie, Show } from "@/models/models";
import { searchMoviesAction, searchShowsAction } from "@/actions/search-actions";
import { useRouter } from "next/navigation";
import { useToast } from "@/components/ui/use-toast";

export default function Navbar() {
  const { user } = useUserStore();

  return (
    <>
      <nav className="container max-w-[1200px] flex justify-between py-2 my-1">
        <Link href="/" className="flex items-center">
          <img src={Logo.src} alt="Cine" className="w-20 sm:w-24" />
        </Link>

        <div className="flex items-center gap-1 md:gap-2">
          {!user.loggedIn && <>
            <Link
              href={"/login"}
              className={cn(
                buttonVariants({ variant: "link" }),
                "text-sm md:text-base text-brand-light hover:text-brand-yellow w-16",
              )}
            >
              Login
            </Link>
            <Link
              href={"/register"}
              className={cn(
                buttonVariants({ variant: "link" }),
                "text-sm md:text-base text-brand-light hover:text-brand-yellow w-16",
              )}
            >
              Register
            </Link>
          </>}
          {user.loggedIn && <>
            <SearchBar />
            <ProfileDropDown />
          </>}
        </div>
      </nav>
      <hr className="border-black mb-8" />
    </>
  );
}

export function ProfileDropDown() {
  const { user, logout } = useUserStore();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="ml-2" asChild>
        <img
          src={user.profile_picture}
          alt="Profile Picture"
          className="
            w-[2.25rem] h-[2.25rem] border-b-brand-darker rounded-full hover:cursor-pointer hover:ring-[2px]
            hover:ring-brand-yellow hover:ring-opacity-80
          "
        />
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56 bg-brand-darker border-brand-yellow">
        <DropdownMenuLabel className="text-brand-yellow">My Account</DropdownMenuLabel>
        <DropdownMenuSeparator className="bg-brand-yellow" />
        <Link href={"/profile/me"}>
          <DropdownMenuItem className="text-brand-yellow hover:cursor-pointer">
            <User className="mr-2 h-4 w-4" />
            <span>Profile</span>
            <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
          </DropdownMenuItem>
        </Link>
        <Link href={"/profile/me/lists"}>
          <DropdownMenuItem className="text-brand-yellow hover:cursor-pointer">
            <ListChecks className="mr-2 h-4 w-4" />
            <span>Lists</span>
            <DropdownMenuShortcut>⇧⌘L</DropdownMenuShortcut>
          </DropdownMenuItem>
        </Link>
        <DropdownMenuItem onClick={logout} className="text-brand-yellow hover:cursor-pointer">
          <LogOut className="mr-2 h-4 w-4" />
          <span>Log out</span>
          <DropdownMenuShortcut>⇧⌘Q</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}

function SearchBar() {
  const { toast } = useToast();

  const router = useRouter();

  const [filter, setFilter] = useState("Movies");
  const [query, setQuery] = useState("");

  const [movies, setMovies] = useState<Movie[]>([]);
  const [shows, setShows] = useState<Show[]>([]);
  const currentResults = filter === "Movies" ? movies : shows;

  const previousQuery = useRef("");
  const searchButtonRef = useRef<HTMLButtonElement>(null);

  useEffect(() => {
    const timer = setTimeout(() => {
      if (query === "") {
        filter === "Movies" ? setMovies([]) : setShows([]);
      } else if (query === previousQuery.current) {
        search(query, filter);
      }
    }, 500);

    previousQuery.current = query;

    return () => clearTimeout(timer);
  }, [query, filter]);

  const search = async (query: string, filter: string) => {
    switch (filter) {
      case "Movies":
        const movieResult = await searchMoviesAction(query);
        if (!movieResult.success) return errorToast(toast, "Failed to search movies", "Please try again later");
        setMovies(movieResult.movies);
        break;
      case "Shows":
        const showResult = await searchShowsAction(query);
        if (!showResult.success) return errorToast(toast, "Failed to search shows", "Please try again later");
        setShows(showResult.shows);
        break;
    }
  }

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    if (query.length === 0) return;

    router.push(`/search/?q=${query}&filter=${filter.toLowerCase()}`);
    searchButtonRef.current?.click();
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button
          ref={searchButtonRef}
          className="
            flex justify-start p-2 text-sm md:text-base text-stone-400 bg-brand-darker text-opacity-80 w-32 md:w-40 h-8
            rounded-xl hover:ring-[1px] hover:ring-brand-yellow hover:ring-opacity-80 hover:text-brand-light
          "
        >
          <Search className="h-[12px] md:h-[16px]" />
          Search...
        </Button>
      </DialogTrigger>
      <DialogContent className="bg-brand-dark border-brand-yellow md:min-w-[40%]">
        <DialogHeader>
          <DialogTitle className="text-brand-yellow">Search</DialogTitle>
          <DialogDescription>Find the cinema you're looking for.</DialogDescription>
        </DialogHeader>

        <DialogBody>
          <form onSubmit={handleSubmit} className="grid grid-cols-4">
            <input
              type="text"
              placeholder={`Search ${filter}...`}
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              className="
                col-span-2 border border-brand-yellow h-10
                flex justify-start p-2 text-sm md:text-base text-stone-400 bg-brand-darker text-opacity-80 w-full
                rounded-l-xl hover:ring-[1px] hover:ring-brand-yellow hover:ring-opacity-80 hover:text-brand-light
              "
            />
            <SearchFilterSelect selected={filter} setFilter={setFilter} />
            <Button
              className="
                rounded-l-none rounded-r-xl h-10
                flex justify-center p-2 text-sm md:text-base text-black bg-brand-yellow text-opacity-80 w-full
                hover:ring-[1px] hover:ring-brand-yellow hover:ring-opacity-80 hover:text-brand-light
              "
            >
              Search
            </Button>
          </form>

          <div className="flex flex-col mt-4">
            {currentResults.slice(0, 8).map((result) => (
              <Link
                key={result.id}
                onClick={() => searchButtonRef.current?.click()}
                href={`/${filter.toLowerCase()}/${result.id}`}
                className="flex items-center gap-2 p-2 border-b border-stone-400 hover:bg-brand-darker"
              >
                <h3 className="text-brand-light">
                  {"title" in result ? result.title : result.name} {" "}
                  <span className="text-stone-500">
                    ({("release_date" in result ? result.release_date : result.first_air_date).slice(0, 4)})
                  </span>
                </h3>
              </Link>
            ))}
          </div>
        </DialogBody>
      </DialogContent>
    </Dialog>
  );
}

const SearchFilterSelect = ({ selected, setFilter }: { selected: string, setFilter: (s: string) => void }) => (
  <Select defaultValue={selected} onValueChange={v => setFilter(v)}>
    <SelectTrigger className="h-10 text-brand-yellow bg-brand-darker border-brand-yellow rounded-none">
      <SelectValue placeholder={selected} />
    </SelectTrigger>
    <SelectContent className="bg-brand-darker border-brand-yellow">
      <SelectGroup>
        {["Movies", "Shows"].map((filter) => (
          <SelectItem
            key={filter}
            value={filter}
            className="hover:cursor-pointer focus:bg-brand-dark text-brand-yellow focus:text-brand-yellow"
          >
            {filter}
          </SelectItem>
        ))}
      </SelectGroup>
    </SelectContent>
  </Select>
)

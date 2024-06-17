import { Movie, Show } from "@/models/models";
import Link from "next/link";
import { imageBase, posterErrorImage } from "@/lib/constants";

export const MovieCard = ({ movie }: { movie: Movie }) => (
  <div className="flex gap-2 bg-brand-darker border border-brand-yellow rounded-xl p-2 mb-2 hover:bg-brand-dark">
    <Link href={`/movies/${movie.id}`}>
      <img
        src={`${imageBase}${movie.poster_path}`}
        alt={movie.title}
        onError={(e) => (e.target as HTMLImageElement).src = posterErrorImage}
        className="w-24 h-36 lg:w-32 lg:h-48 object-cover rounded-xl hover:opacity-60 transition duration-100 ease-in-out"
      />
    </Link>
    <div className="flex flex-col justify-between w-full">
      <div>
        <Link href={`/movies/${movie.id}`} className="w-full group">
          <h2 className="text-brand-yellow font-bold text-2xl group-hover:text-brand-light">{movie.title}</h2>
        </Link>
        <p className="text-brand-light">{movie.overview.slice(0, movie.overview.indexOf(".") + 1)}</p>
      </div>
      <div className="flex justify-between">
        <p className="text-lg text-stone-400 self-end">Movie</p>
        <p className="text-lg text-stone-400 self-end">{movie.release_date.slice(0, 4) || "unreleased"}</p>
      </div>
    </div>
  </div>
);

export const ShowCard = ({ show }: { show: Show }) => (
  <div className="flex gap-2 bg-brand-darker border border-brand-yellow rounded-xl p-2 mb-2 hover:bg-brand-dark">
    <Link href={`/movies/${show.id}`}>
      <img
        src={`${imageBase}${show.poster_path}`}
        alt={show.name}
        onError={(e) => (e.target as HTMLImageElement).src = posterErrorImage}
        className="w-24 h-36 lg:w-32 lg:h-48 object-cover rounded-xl hover:opacity-60 transition duration-100 ease-in-out"
      />
    </Link>
    <div className="flex flex-col justify-between w-full">
      <div>
        <Link href={`/shows/${show.id}`} className="w-full group">
          <h2 className="text-brand-yellow font-bold text-2xl group-hover:text-brand-light">{show.name}</h2>
        </Link>
        <p className="text-brand-light">{show.overview.slice(0, show.overview.indexOf(".") + 1)}</p>
      </div>
      <div className="flex justify-between">
        <p className="text-lg text-stone-400 self-end">Show</p>
        <p className="text-lg text-stone-400 self-end">{show.first_air_date.slice(0, 4) || "unreleased"}</p>
      </div>
    </div>
  </div>
);
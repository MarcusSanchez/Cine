import { MediaType, Movie, MovieList, Show, ShowList } from "@/models/models";
import { Carousel, CarouselContent, CarouselItem } from "@/components/ui/carousel";
import Link from "next/link";
import { Skeleton } from "@/components/ui/skeleton";
import { imageBase } from "@/lib/constants";

type ListProps = {
  mediaType: MediaType,
  list: MovieList | ShowList,
  items: Movie[] | Show[]
};

const List = ({ mediaType, list, items }: ListProps) => (
  <div>
    <h2 className="text-2xl md:text-3xl font-bold text-brand-yellow mt-8 mb-4">{listName(mediaType, list)}</h2>
    <Carousel
      opts={{
        align: "start",
      }}
    >
      <CarouselContent className="w-[95%]">
        {items.map((item, i) => (
          <CarouselItem key={i} className="basis-1/3 sm:basis-1/4 md:basis-1/5 lg:basis-1/6 xl:basis-1/7 h-min">
            <ItemCard item={item} mediaType={mediaType} />
          </CarouselItem>
        ))}
      </CarouselContent>
    </Carousel>
  </div>
)

const ItemCard = ({ item, mediaType }: { item: Movie | Show, mediaType: MediaType }) => (
  <div className="flex flex-col group">
    <Link href={`/${mediaType}s/${item.id}`}>
      <img
        src={`${imageBase}${item.poster_path}`}
        alt={"title" in item ? item.title : item.name}
        className="rounded-lg w-full object-cover group-hover:opacity-60 transition-opacity duration-200 ease-in-out"
      />
    </Link>
    <Link href={`/${mediaType}s/${item.id}`}>
      <p className="text-sm sm:text-base md:text-lg text-brand-light mt-2 group-hover:text-brand-yellow">
        {"title" in item ? item.title : item.name}
      </p>
    </Link>
    <p className="text-sm md:text-base text-stone-400">
      {("release_date" in item ? item.release_date : item.first_air_date).slice(0, 4)}
    </p>
  </div>
);

export default List;

function listName(mediaType: MediaType, list: MovieList | ShowList) {
  switch (mediaType) {
    case MediaType.Movie:
      switch (list) {
        case MovieList.NowPlaying:
          return "Now Playing Movies";
        case MovieList.TopRated:
          return "Top Rated Movies";
        case MovieList.Popular:
          return "Popular Movies";
        case MovieList.Upcoming:
          return "Upcoming Movies";
      }
    case MediaType.Show:
      switch (list) {
        case ShowList.AiringToday:
          return "Airing Today Shows";
        case ShowList.OnTheAir:
          return "On The Air Shows";
        case ShowList.Popular:
          return "Popular Shows";
        case ShowList.TopRated:
          return "Top Rated Shows";
      }
  }
}

export const ListsSkeleton = () => (
  Array.from({ length: 5 }).map((_, i) => (
    <div key={i}>
      <Skeleton className="text-2xl md:text-3xl font-bold text-brand-yellow mt-8 mb-4 h-8 w-[40%]" />
      <Carousel
        opts={{
          align: "start",
        }}
      >
        <CarouselContent className="w-[95%]">
          {Array.from({ length: 9 }).map((_, i) => (
            <CarouselItem key={i} className="basis-1/3 sm:basis-1/4 md:basis-1/5 lg:basis-1/6 xl:basis-1/7 h-min">
              <div className="flex flex-col group">
                <Skeleton className="rounded-lg w-full h-[175px] md:h-[225px]" />
                <Skeleton className="h-4 mt-2 w-[80%]" />
                <Skeleton className="h-4 mt-2 w-[40%]" />
              </div>
            </CarouselItem>
          ))}
        </CarouselContent>
      </Carousel>
    </div>
  ))
);
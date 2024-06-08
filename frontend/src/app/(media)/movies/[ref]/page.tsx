"use client";

import MovieContent from "./(components)/MovieContent";
import MovieCredits from "./(components)/MovieCredits";
import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import fetchMovieAction from "@/actions/fetch-movie-action";
import Reviews from "@/components/Reviews";
import { useUserStore } from "@/app/state";
import ReviewForm from "@/components/UserReview";
import { DetailedMovie, DetailedReview, MediaType } from "@/models/models";
import { Button } from "@/components/ui/button";
import fetchReviewsAction from "@/actions/fetch-reviews-action";
import Comments from "@/components/Comments";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";


export default function MoviePage({ params }: { params: { ref: number } }) {
  const { user } = useUserStore();
  const { toast } = useToast();

  const [movie, setMovie] = useState<DetailedMovie | null>(null);

  const [reviews, setReviews] = useState<DetailedReview[] | null>(null);
  const userReview = reviews?.find((review) => review.user.id === user?.id) ?? null;

  const [viewReviewForm, setViewReviewForm] = useState(userReview === null);

  const router = useRouter();
  const buttonRef = React.createRef<HTMLButtonElement>();

  // handle close after edit review success
  useEffect(() => {
    if (!viewReviewForm) return;
    buttonRef.current?.click();
  }, [reviews])

  useEffect(() => {
    const fetchMovie = async () => {
      const result = await fetchMovieAction(params.ref);
      if (!result.success) return router.replace("/uh-oh");
      setMovie(result.detailedMovie);
    }

    const fetchReviews = async () => {
      const result = await fetchReviewsAction(MediaType.Movie, params.ref);
      if (!result.success) return errorToast(toast, "Failed to fetch reviews", "Please try again later");

      const userReview = result.detailedReviews.find((r) => r.user.id === user.id);

      userReview
        ? setReviews([userReview, ...result.detailedReviews.filter((r) => r.user.id !== user.id)])
        : setReviews(result.detailedReviews);
    }

    fetchMovie().then(() => fetchReviews());
  }, []);

  return (
    <>
      <MovieContent movie={movie} />
      <MovieCredits refID={params.ref} />

      <div className="container max-w-[1200px] mt-8">
        <div className="flex">
          <h1 className="text-2xl md:text-4xl mb-2 text-brand-yellow font-bold">Reviews</h1>
          {reviews && userReview &&
            <div className="flex items-end">
              <Button
                ref={buttonRef}
                variant="link"
                onClick={() => setViewReviewForm(!viewReviewForm)}
                className="text-brand-yellow"
              >
                {!viewReviewForm ? "Edit or Delete Review" : "Hide Form"}
              </Button>
            </div>
          }
        </div>
        {reviews && viewReviewForm &&
          <ReviewForm refID={params.ref} reviews={reviews} setReviews={setReviews} setViewReviewForm={setViewReviewForm} />
        }
      </div>
      <Reviews reviews={reviews} />

      <Comments media={MediaType.Movie} refID={params.ref} />
    </>
  );
}

"use client";

import React, { createRef, useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import fetchShowAction from "@/actions/fetch-show-action";
import Reviews from "@/components/Reviews";
import { useUserStore } from "@/app/state";
import ReviewForm from "@/components/UserReview";
import { DetailedReview, DetailedShow, MediaType } from "@/models/models";
import { Button } from "@/components/ui/button";
import fetchReviewsAction from "@/actions/fetch-reviews-action";
import Comments from "@/components/Comments";
import ShowContent from "@/app/(media)/shows/[ref]/(components)/ShowContent";
import ShowCredits from "@/app/(media)/shows/[ref]/(components)/ShowCredits";
import ShowSeasons from "@/app/(media)/shows/[ref]/(components)/ShowSeasons";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

export default function ShowPage({ params }: { params: { ref: number } }) {
  const { user } = useUserStore();
  const { toast } = useToast();

  const [show, setShow] = useState<DetailedShow | null>(null);

  const [reviews, setReviews] = useState<DetailedReview[] | null>(null);
  const userReview = reviews?.find((review) => review.user.id === user?.id) ?? null;

  const [viewReviewForm, setViewReviewForm] = useState(userReview === null);

  const router = useRouter();
  const buttonRef = createRef<HTMLButtonElement>();

  // handle close after edit review success
  useEffect(() => {
    if (!viewReviewForm) return;
    buttonRef.current?.click();
  }, [reviews])

  useEffect(() => {
    const fetchShow = async () => {
      const result = await fetchShowAction(params.ref);
      if (!result.success) return router.replace("/404");
      setShow(result.detailedShow);
    }

    const fetchReviews = async () => {
      const result = await fetchReviewsAction(MediaType.Show, params.ref);
      if (!result.success) return errorToast(toast, "Failed to fetch reviews", "Please try again later");

      const userReview = result.detailedReviews.find((r) => r.user.id === user.id);

      userReview
        ? setReviews([userReview, ...result.detailedReviews.filter((r) => r.user.id !== user.id)])
        : setReviews(result.detailedReviews);
    }

    fetchShow().then(() => fetchReviews());
  }, []);

  return (
    <>
      <ShowContent show={show} />
      <ShowSeasons show={show} />
      <ShowCredits refID={params.ref} />

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
          <ReviewForm
            refID={params.ref}
            media={MediaType.Show}
            reviews={reviews}
            setReviews={setReviews}
            setViewReviewForm={setViewReviewForm}
          />
        }
      </div>
      <Reviews reviews={reviews} />

      <Comments media={MediaType.Show} refID={params.ref} />
    </>
  );
}

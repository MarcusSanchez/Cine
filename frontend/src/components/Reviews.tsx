"use client";

import { Skeleton } from "@/components/ui/skeleton";
import Link from "next/link";
import React from "react";
import { Button } from "@/components/ui/button";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger, } from "@/components/ui/dialog"
import { DialogBody } from "next/dist/client/components/react-dev-overlay/internal/components/Dialog";
import { DetailedReview } from "@/models/models";

const Reviews = ({ reviews }: { reviews: DetailedReview[] | null }) => (
  !reviews ? <ReviewsSkeleton /> : (
    <div className="container max-w-[1200px]">
      {reviews.length < 1 &&
        <div className="bg-brand-dark p-4 rounded-xl">
          <p className="text-brand-light">No reviews yet. Be the first to create one!</p>
        </div>
      }

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {reviews.slice(0, 3).map((review) => (
          <ReviewCard key={review.review.id} review={review} />
        ))}
      </div>

      {reviews.length > 3 &&
        <Dialog>
          <DialogTrigger asChild>
            <Button variant="link" className="text-brand-yellow px-0">
              Show all reviews
            </Button>
          </DialogTrigger>
          <DialogContent className="bg-brand-dark border-brand-yellow overflow-y-scroll h-min h-max-[80%]">
            <DialogHeader>
              <DialogTitle className="text-brand-yellow">Reviews</DialogTitle>
            </DialogHeader>

            <DialogBody className="w-full">
              <div className="grid grid-cols-1 gap-4">
                {reviews.map((review) => (
                  <ReviewCard key={review.review.id} review={review} />
                ))}
              </div>
            </DialogBody>
          </DialogContent>
        </Dialog>
      }
    </div>
  )
);

export const ReviewCard = ({ review }: { review: DetailedReview }) => (
  <div className="bg-brand-dark p-4 rounded-xl border border-brand-yellow">
    <div className="flex items-center justify-between">
      <Link href={`/profile/${review.user.id}`}>
        <div className="flex items-center">
          <img src={review.user.profile_picture} alt={review.user.username} className="w-12 h-12 rounded-full" />
          <div className="ml-4 flex flex-col">
            <span className="text-xl text-brand-yellow font-bold">{review.user.username}</span>
            <span className="text-sm text-brand-light font-bold">@{review.user.username}</span>
          </div>
        </div>
      </Link>
      <span className="border-4 border-brand-yellow w-8 h- rounded-3xl text-brand-yellow font-bold flex items-center justify-center">
        {review.review.rating}
      </span>
    </div>
    <p className="ml-4 mt-4 text-brand-light break-words hyphens-auto">{review.review.content}</p>
  </div>
);

const ReviewsSkeleton = () => (
  <div className="container max-w-[1200px]">
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
      {[...Array(3)].map((_, index) => (
        <div key={index} className="bg-brand-dark p-4 rounded-xl">
          <div className="flex items-center justify-between">
            <div className="flex items-center">
              <Skeleton className="w-12 h-12 bg-brand-light rounded-full"></Skeleton>
              <div className="ml-4">
                <Skeleton className="w-24 h-4 bg-brand-light rounded"></Skeleton>
                <Skeleton className="w-16 h-4 bg-brand-light rounded mt-1"></Skeleton>
              </div>
            </div>
            <Skeleton className="w-12 h-12 bg-brand-light rounded-full"></Skeleton>
          </div>
          <Skeleton className="w-full h-4 bg-brand-light rounded mt-4"></Skeleton>
          <Skeleton className="w-full h-4 bg-brand-light rounded mt-2"></Skeleton>
          <Skeleton className="w-full h-4 bg-brand-light rounded mt-2"></Skeleton>
          <Skeleton className="w-1/2 h-4 bg-brand-light rounded mt-4"></Skeleton>
        </div>
      ))}
    </div>
  </div>
);

export default Reviews;


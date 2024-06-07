import React, { ChangeEvent, FormEvent, useState } from "react";
import { z } from "zod";
import { Textarea } from "@/components/ui/textarea";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import createReviewAction from "@/actions/create-review-action";
import { useUserStore } from "@/app/state";
import { DetailedReview, MediaType } from "@/models/models";
import editReviewAction from "@/actions/edit-review-action";
import deleteReviewAction from "@/actions/delete-review-action";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

const formSchema = z.object({
  content: z.string()
    .min(1, "content must be at least 1 character long")
    .max(140, "content must be at most 140 characters long"),
  rating: z.number()
    .gt(0, "rating must be greater than 0")
    .lte(10, "rating must be less than or equal to 10"),
});
export type FormSchema = z.infer<typeof formSchema>;

type ReviewFormProps = {
  refID: number,
  media?: MediaType,
  reviews: DetailedReview[],
  setReviews: (reviews: DetailedReview[]) => void,
  setViewReviewForm: (view: boolean) => void
};

export default function ReviewForm({
  refID,
  media = MediaType.Movie,
  reviews,
  setReviews,
  setViewReviewForm
}: ReviewFormProps) {
  const { user } = useUserStore();
  const { toast } = useToast();

  const userReview = reviews.find((review) => review.user.id === user.id);

  const [data, setData] = useState({
    content: userReview?.review.content ?? "",
    rating: userReview?.review.rating.toString() ?? "",
  });

  const handleChange = (key: keyof FormSchema) => (e: ChangeEvent<HTMLTextAreaElement | HTMLInputElement>) => {
    switch (key) {
      case "content":
        setData({ ...data, content: e.target.value.slice(0, 140) });
        break;
      case "rating":
        if (e.target.value === "") return setData({ ...data, rating: "" });
        if (!/^\d+$/.test(e.target.value)) return;

        const rating = parseInt(e.target.value);
        if (rating > 10 || rating < 1) return;

        setData({ ...data, rating: e.target.value });
        break;
    }
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    const values = formSchema.safeParse({ ...data, rating: parseInt(data.rating) });
    if (!values.success) return;

    if (!userReview) {
      const result = await createReviewAction(user.csrf, values.data, media, refID);
      if (!result.success) return errorToast(toast, "Failed to create review.", result.error);
      setReviews([{ user, review: result.review }, ...reviews]);
    } else {
      const result = await editReviewAction(user.csrf, values.data, userReview.review.id);
      if (!result.success) return errorToast(toast, "Failed to edit review.", result.error);
      setReviews(reviews.map((review) => review.user.id === user.id ? { user, review: result.review } : review));
    }
  };

  const deleteReview = async () => {
    const result = await deleteReviewAction(user.csrf, userReview!.review.id)
    if (!result.success) return errorToast(toast, "Failed to delete review.", result.error);

    setReviews(reviews.filter((review) => review.user.id !== user.id));
    setViewReviewForm(true);
    setData({ content: "", rating: "" });
  };

  return (
    <form onSubmit={handleSubmit} className="bg-brand-dark p-4 pt-2 rounded-xl mb-3">
      <Label className="text-xl text-brand-yellow">Review Content</Label>
      <Textarea
        value={data.content}
        onChange={handleChange("content")}
        placeholder="Type your review here."
        className="bg-brand-darker text-white text-base border border-brand-yellow"
      />
      <p className="text-sm text-brand-light flex justify-end mt-1">{data.content.length}/140</p>

      <div className="flex flex-col">
        <Label className="text-xl text-brand-yellow">Rating</Label>
        <div className="flex gap-2 justify-between">
          <Input
            value={data.rating}
            onChange={handleChange("rating")}
            placeholder="1-10"
            className="bg-brand-darker text-white text-base p-[.35rem] border border-brand-yellow rounded-xl w-24"
          />

          <div className="flex gap-1">
            {userReview &&
              <Button
                onClick={deleteReview}
                className="
                  bg-brand-red text-white hover:text-brand-red hover:border-brand-red hover:border
                  border-brand-red px-4 py-2 rounded-xl
                "
              >
                Delete Review
              </Button>
            }
            <Button
              disabled={data.content === "" || data.rating === ""}
              className="
                bg-brand-yellow text-brand-darker hover:text-brand-yellow hover:border-brand-yellow hover:border
                border-brand-yellow px-4 py-2 rounded-xl
              "
            >
              {userReview ? "Edit Review" : "Submit Review"}
            </Button>
          </div>
        </div>
      </div>
    </form>
  );
}
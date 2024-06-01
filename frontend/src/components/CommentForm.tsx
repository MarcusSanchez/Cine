import { Dispatch, FormEvent, SetStateAction, useState } from "react";
import { DetailedComment, MediaType } from "@/models/models";
import { useUserStore } from "@/app/state";
import { Button } from "@/components/ui/button";
import createCommentAction from "@/actions/create-comment-action";

type CommentFormProps = {
  comments: DetailedComment[],
  setComments: Dispatch<SetStateAction<DetailedComment[]>>,
  media: MediaType,
  refID: number,
  replyingTo?: string
};

export default function CommentForm({ comments, setComments, media, refID, replyingTo }: CommentFormProps) {
  const { user } = useUserStore();
  const [content, setContent] = useState("");

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    const result = await createCommentAction(user.csrf, content, media, refID, replyingTo);
    if (!result.success) return;
    setContent("");
    setComments([{
      user,
      comment: result.comment,
      liked_by_user: false,
      likes_count: 0,
      replies_count: 0
    }, ...comments])
    if (replyingTo) {
      setComments(comments => comments.map((c) =>
        c.comment.id === replyingTo
          ? { ...c, replies_count: c.replies_count + 1 }
          : c
      ));
    }
  };

  return (
    <form onSubmit={handleSubmit} className="bg-brand-dark rounded-xl flex">
      <div className="w-full">
        <input
          type="text"
          placeholder={`Leave a ${replyingTo ? "reply" : "comment"}`}
          value={content}
          onChange={(e) => setContent((e.target as HTMLInputElement).value.slice(0, 280))}
          className="bg-brand-dark text-white w-full py-1 px-2 rounded-l-xl border border-brand-yellow h-10"
        />
        <p className="text-sm text-brand-light flex justify-end mt-1">{content.length}/280</p>
      </div>
      <Button
        disabled={content === ""}
        type="submit"
        className="
          bg-brand-yellow text-brand-dark p-2 rounded-l-none h-10 rounded-r-xl hover:text-brand-yellow
          hover:border-brand-yellow hover:border border-brand-yellow px-4 py-2
        "
      >
        {replyingTo ? "Reply" : "Comment"}
      </Button>
    </form>
  );
}

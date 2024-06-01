import React, { Dispatch, SetStateAction, useEffect, useState } from "react";
import { DetailedComment, MediaType } from "@/models/models";
import { useUserStore } from "@/app/state";
import fetchCommentsAction from "@/actions/fetch-comments-action";
import CommentForm from "@/components/CommentForm";
import Link from "next/link";
import { ThumbsUp } from "lucide-react";
import { likeCommentAction, unlikeCommentAction } from "@/actions/like-unlike-actions";
import { Button } from "@/components/ui/button";
import { timeAgo } from "@/lib/utils";
import deleteCommentAction from "@/actions/delete-comment-action";
import Replies from "@/components/Replies";

type CommentsProps = {
  media: MediaType;
  refID: number;
};

export default function Comments({ media, refID }: CommentsProps) {
  const { user } = useUserStore();

  let [comments, setComments] = useState<DetailedComment[]>([]);
  comments = [...comments.filter((c) => c.user.id === user.id), ...comments.filter((c) => c.user.id !== user.id)];
  const topLevelComments = comments.filter((c) => !c.comment.replying_to_id);

  const [viewReplyForms, setViewReplyForms] = useState<DetailedComment[]>([]);

  useEffect(() => {
    const fetchComments = async () => {
      const result = await fetchCommentsAction(media, refID);
      if (!result.success) return;
      setComments(result.detailedComments);
    };

    fetchComments();
  }, []);

  return (
    <div className="container max-w-[1200px] mt-8">
      <h1 className="text-2xl md:text-4xl mb-2 text-brand-yellow font-bold">Comments</h1>
      <CommentForm comments={comments} setComments={setComments} media={media} refID={refID} />
      <div>
        {topLevelComments.map((c) => (
          <Comment
            key={c.comment.id}
            c={c}
            viewReplyForms={viewReplyForms}
            setViewReplyForms={setViewReplyForms}
            comments={comments}
            setComments={setComments}
            media={media}
            refID={refID}
          />
        ))}
      </div>
    </div>
  );
}

type CommentProps = {
  c: DetailedComment,
  topLevel?: boolean,
  viewReplyForms: DetailedComment[],
  setViewReplyForms: Dispatch<SetStateAction<DetailedComment[]>>,
  comments: DetailedComment[],
  setComments: Dispatch<SetStateAction<DetailedComment[]>>,
  media: MediaType,
  refID: number
};

export function Comment({
  c,
  topLevel = true,
  viewReplyForms,
  setViewReplyForms,
  comments,
  setComments,
  media,
  refID
}: CommentProps) {
  const { user } = useUserStore();
  const [showReplies, setShowReplies] = useState<DetailedComment[]>([])

  const handleDelete = async (commentID: string) => {
    const result = await deleteCommentAction(user.csrf, commentID);
    if (!result.success) return;

    // if the comment is replying to another comment, decrement the replies_count of the parent comment
    if (comments.find((c) => c.comment.id === commentID)?.comment.replying_to_id) {
      const parentComment = comments.find((c) => c.comment.id === comments.find((c) => c.comment.id === commentID)?.comment.replying_to_id);
      if (parentComment) {
        setComments(comments => comments.map((c) => c.comment.id === parentComment.comment.id ? {
          ...c,
          replies_count: c.replies_count - 1
        } : c));
      }
    }

    // if the comment has replies, cascade delete them.
    let comment = comments.find((c) => c.comment.id === commentID);
    let replies = [comment];
    while (replies.length > 0) {
      const curr = replies.pop();
      const currReplies = comments.filter((c) => c.comment.replying_to_id === curr?.comment.id);
      replies = [...replies, ...currReplies];
      setComments(comments => comments.filter((c) => c.comment.id !== curr?.comment.id));
    }

    setComments(comments => comments.filter((c) => c.comment.id !== commentID));
  };

  const handleLike = async ({ comment: { id: commentID }, likes_count, liked_by_user }: DetailedComment) => {
    const action = liked_by_user ? unlikeCommentAction : likeCommentAction;

    const result = await action(user.csrf, commentID);
    if (!result.success) return;

    const newLikeCount = likes_count + (liked_by_user ? -1 : 1);

    setComments(comments.map((c) => c.comment.id === commentID ? {
      ...c,
      likes_count: newLikeCount,
      liked_by_user: !liked_by_user
    } : c));
  };

  const hideOrShowReplyForm = (comment: DetailedComment) => {
    viewReplyForms.includes(c)
      ? setViewReplyForms(viewReplyForms.filter(c => c.comment.id !== comment.comment.id))
      : setViewReplyForms([...viewReplyForms, comment]);
  };

  const hideOrShowReplies = (comment: DetailedComment) => {
    showReplies.includes(c)
      ? setShowReplies(showReplies.filter(c => c.comment.id !== comment.comment.id))
      : setShowReplies([...showReplies, comment]);
  }

  const getUserFromReplyingToID = (replyingToID: string) => {
    return comments.find((c) => c.comment.id === replyingToID)?.user;
  }

  return (
    <div className="mb-2">
      <div key={c.comment.id} className="bg-brand-dark rounded-xl py-4 flex">
        <img src={c.user.profile_picture} alt="avatar" className="w-12 h-12 rounded-full mr-2" />

        <div className="flex flex-col w-full">
          <div className="w-full">
            <Link href={`/profile/${c.user.id}`} className="flex flex-col w-min">
              <p className="text-base md:text-lg text-brand-yellow font-bold w-min text-nowrap">{c.user.display_name}</p>
              <p className="text-sm md:text-base text-brand-light mt-[-5px] w-min text-nowrap">@{c.user.username}</p>
            </Link>
            <p className="text-white mt-1 text-sm md:text-base">
              {c.comment.replying_to_id &&
                <Link
                  href={`/profile/${getUserFromReplyingToID(c.comment.replying_to_id)?.id}`}
                  className="text-brand-yellow font-bold"
                >
                  {`@${getUserFromReplyingToID(c.comment.replying_to_id)?.username} `}
                </Link>
              }
              {c.comment.content}
            </p>
          </div>
          <div className="flex gap-2 mt-2">
            <Button
              variant="link"
              onClick={() => hideOrShowReplyForm(c)}
              className="text-sm md:text-base text-brand-yellow h-2 w-min px-0"
            >
              {viewReplyForms.includes(c) ? "Cancel" : "Reply"}
            </Button>
            {c.replies_count > 0 &&
              <Button
                variant="link"
                onClick={() => hideOrShowReplies(c)}
                className="text-sm md:text-base text-brand-yellow h-2 w-min px-0"
              >
                {showReplies.includes(c) ? "Hide" : "View"} {c.replies_count} Replies
              </Button>
            }
            {c.comment.user_id === user.id &&
              <Button
                variant="link"
                onClick={() => handleDelete(c.comment.id)}
                className="text-sm md:text-base text-brand-red h-2 w-min px-0"
              >
                Delete
              </Button>
            }
            <p className="text-xs md:text-sm mt-[-1px] text-stone-500">{timeAgo(c.comment.created_at)}</p>
          </div>
        </div>

        <div className="flex flex-col items-center">
          <ThumbsUp
            onClick={() => handleLike(c)}
            {...(c.liked_by_user ? { fill: "currentColor" } : {})}
            size={24}
            className={"text-white hover:cursor-pointer"}
          />
          <p className="text-white">{c.likes_count}</p>
        </div>
      </div>
      {viewReplyForms.includes(c) &&
        <div className="w-[95%] ml-auto">
          <CommentForm
            comments={comments}
            setComments={setComments}
            media={media}
            refID={refID}
            replyingTo={c.comment.id}
          />
        </div>
      }
      {showReplies.includes(c) &&
        <Replies
          c={c}
          topLevel={topLevel}
          viewReplyForms={viewReplyForms}
          setViewReplyForms={setViewReplyForms}
          comments={comments}
          setComments={setComments}
          showReplies={showReplies}
          media={media}
          refID={refID}
        />
      }
    </div>
  );
}
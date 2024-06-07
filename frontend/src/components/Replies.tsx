import { DetailedComment, MediaType } from "@/models/models";
import { Dispatch, SetStateAction, useEffect } from "react";
import { Comment } from "./Comments";
import fetchRepliesAction from "@/actions/fetch-replies-action";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

const alreadyFetchedReplies = new Set<DetailedComment>();

type RepliesProps = {
  c: DetailedComment,
  topLevel: boolean,
  viewReplyForms: DetailedComment[],
  setViewReplyForms: Dispatch<SetStateAction<DetailedComment[]>>,
  comments: DetailedComment[],
  setComments: Dispatch<SetStateAction<DetailedComment[]>>,
  showReplies: DetailedComment[],
  media: MediaType,
  refID: number,
};

export default function Replies({
  c, /* comment */
  topLevel,
  viewReplyForms, setViewReplyForms,
  comments, setComments,
  showReplies,
  media, refID
}: RepliesProps) {
  const { toast } = useToast();

  let orderedReplies = [...comments.filter((r) => r.comment.replying_to_id === c.comment.id)];

  const fetchReplies = async (c: DetailedComment) => {
    if (alreadyFetchedReplies.has(c)) return;

    const result = await fetchRepliesAction(c.comment.id);
    if (!result.success) return errorToast(toast, "Failed to fetch replies", "Please try again later");
    setComments(removeDuplicates([...comments, ...result.replies]));
    alreadyFetchedReplies.add(c);
  }

  for (const reply of showReplies) {
    if (reply.replies_count < 0) continue;
    if (alreadyFetchedReplies.has(reply)) {
      const index = orderedReplies.findIndex((r) => r.comment.id === reply.comment.id);
      if (index === -1) continue;
      orderedReplies = [
        ...orderedReplies.slice(0, index),
        ...comments.filter((c) => c.comment.replying_to_id === reply.comment.id),
        ...orderedReplies.slice(index + 1)
      ];

      continue;
    }
    fetchReplies(reply);
  }

  useEffect(() => {
    fetchReplies(c)
  }, []);

  return (
    <div className={topLevel ? "ml-14" : ""}>
      {orderedReplies.map((r) => (
        <Comment
          key={r.comment.id}
          c={r}
          topLevel={false}
          viewReplyForms={viewReplyForms}
          setViewReplyForms={setViewReplyForms}
          comments={comments}
          setComments={setComments}
          media={media}
          refID={refID}
        />
      ))}
    </div>
  );
}

function removeDuplicates(arr: DetailedComment[]) {
  return arr.filter((v, i, a) => a.findIndex(t => (t.comment.id === v.comment.id)) === i);
}

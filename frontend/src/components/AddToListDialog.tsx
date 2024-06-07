import { Button } from "@/components/ui/button";
import React, { useEffect, useRef, useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger
} from "@/components/ui/dialog";
import { DialogBody } from "next/dist/client/components/react-dev-overlay/internal/components/Dialog";
import { DetailedList, MediaType } from "@/models/models";
import { fetchMyListsAction } from "@/actions/fetch-lists-actions";
import { addMovieToListAction, addShowToListAction } from "@/actions/add-remove-media-from-list-action";
import { useUserStore } from "@/app/state";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

export default function AddToListDialog({ mediaType, refID }: { mediaType: "movie" | "show"; refID: number }) {
  const { user } = useUserStore();
  const { toast } = useToast();

  const triggerRef = useRef<HTMLButtonElement>(null);

  const [lists, setLists] = useState<DetailedList[]>([]);

  useEffect(() => {
    const fetchLists = async () => {
      const result = await fetchMyListsAction();
      if (!result.success) return errorToast(toast, "Failed to fetch lists", "Please try again later");
      setLists(result.lists);
    };

    fetchLists();
  }, []);

  const handleAddToList = async (listID: string) => {
    switch (mediaType) {
      case MediaType.Movie:
        const movieResult = await addMovieToListAction(user.csrf, listID, refID);
        if (!movieResult.success) return errorToast(toast, "Failed to add movie to list", movieResult.error);
        break;
      case MediaType.Show:
        const showResult = await addShowToListAction(user.csrf, listID, refID);
        if (!showResult.success) return errorToast(toast, "Failed to add show to list", showResult.error);
        break
    }

    triggerRef.current?.click();
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button
          ref={triggerRef}
          variant="outline"
          className="
            text-brand-yellow border-brand-yellow w-[80px] sm:w-[100px]  h-[32px] md:h-[40px] bg-transparent
            hover:bg-brand-yellow text-sm sm:text-base
          "
        >
          Add To List
        </Button>
      </DialogTrigger>
      <DialogContent className="bg-brand-dark border-brand-yellow md:min-w-[40%] overflow-y-scroll no-scrollbar">
        <DialogHeader>
          <DialogTitle className="text-brand-yellow">Add To List</DialogTitle>
          <DialogDescription>Click a list to add {mediaType === MediaType.Movie ? "movie" : "show"} cinema to.</DialogDescription>
        </DialogHeader>

        <DialogBody>
          <div className="flex flex-col">
            {lists.length < 1 && <h3 className="text-brand-light">No lists found</h3>}
            {lists.map((l) => (
              <div
                onClick={() => handleAddToList(l.list.id)}
                key={l.list.id}
                className="flex items-center gap-2 p-2 border-b border-stone-400 hover:bg-brand-darker hover:cursor-pointer"
              >
                <h3 className="text-brand-light">{l.list.name}</h3>
              </div>
            ))}
          </div>
        </DialogBody>
      </DialogContent>
    </Dialog>
  );
}

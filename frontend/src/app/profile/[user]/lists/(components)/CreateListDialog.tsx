import { Button } from "@/components/ui/button";
import React, { Dispatch, FormEvent, SetStateAction, useRef, useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger
} from "@/components/ui/dialog";
import { DialogBody } from "next/dist/client/components/react-dev-overlay/internal/components/Dialog";
import { createListAction } from "@/actions/create-list-action";
import { useUserStore } from "@/app/state";
import { DetailedList } from "@/models/models";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

type CreateListDialogProps = {
  setLists: Dispatch<SetStateAction<DetailedList[]>>;
};

export default function CreateListDialog({ setLists }: CreateListDialogProps) {
  const { user } = useUserStore();
  const { toast } = useToast();

  const [title, setTitle] = useState("");
  const triggerRef = useRef<HTMLButtonElement>(null);

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();

    const result = await createListAction(user.csrf, title);
    if (!result.success) return errorToast(toast, "Failed to create list", "Please try again later");
    setLists(l => [...l, { list: result.list, movies: [], shows: [], members: [] }])

    setTitle("");
    triggerRef.current?.click();
  }

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
          Create List
        </Button>
      </DialogTrigger>
      <DialogContent className="bg-brand-dark border-brand-yellow md:min-w-[40%]">
        <DialogHeader>
          <DialogTitle className="text-brand-yellow">Create List</DialogTitle>
          <DialogDescription>Time for a new collection.</DialogDescription>
        </DialogHeader>

        <DialogBody>
          <form onSubmit={handleSubmit} className="grid grid-cols-4">
            <input
              type="text"
              placeholder={`List title...`}
              value={title}
              onChange={(e) => setTitle(e.target.value.slice(0, 50))}
              className="
                col-span-3 border border-brand-yellow h-10
                flex justify-start p-2 text-sm md:text-base text-stone-400 bg-brand-darker text-opacity-80 w-full
                rounded-l-xl hover:ring-[1px] hover:ring-brand-yellow hover:ring-opacity-80 hover:text-brand-light
              "
            />
            <Button
              type="submit"
              className="
                rounded-l-none rounded-r-xl h-10
                flex justify-center p-2 text-sm md:text-base text-black bg-brand-yellow text-opacity-80 w-full
                hover:ring-[1px] hover:ring-brand-yellow hover:ring-opacity-80 hover:text-brand-light
              "
              disabled={title.length < 1}
            >
              Create
            </Button>
          </form>
        </DialogBody>
      </DialogContent>
    </Dialog>
  );
}

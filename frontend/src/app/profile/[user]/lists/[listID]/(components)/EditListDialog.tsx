import { Button } from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import React, { Dispatch, SetStateAction, useRef, useState } from "react";
import { DetailedList } from "@/models/models";
import { Select, SelectContent, SelectGroup, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { updateListAction } from "@/actions/update-list-action";
import { useUserStore } from "@/app/state";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

type EditListDialogProps = {
  list: DetailedList;
  setList: Dispatch<SetStateAction<DetailedList | null>>;
};

export default function EditListDialog({ list, setList }: EditListDialogProps) {
  const { user } = useUserStore();
  const { toast } = useToast();

  const triggerRef = useRef<HTMLButtonElement>(null);

  const [title, setTitle] = useState(list.list.name);
  const [visibility, setVisibility] = useState(list.list.is_public ? "Public" : "Private");

  const updateList = async () => {
    const isPublic = visibility === "Public";
    const result = await updateListAction(user.csrf, list.list.id, title, isPublic);
    if (!result.success) return errorToast(toast, "Failed to update list", result.error);

    setList({
      ...list,
      list: {
        ...list.list,
        name: title,
        is_public: isPublic,
      }
    });
    triggerRef.current?.click();
  };

  return (
    <Dialog>
      <DialogTrigger asChild ref={triggerRef}>
        <Button
          variant="outline"
          className="
            text-brand-yellow border-brand-yellow w-[80px] sm:w-[100px]  h-[32px] md:h-[40px] bg-transparent
            hover:bg-brand-yellow text-sm sm:text-base
          "
        >
          Edit List
        </Button>
      </DialogTrigger>
      <DialogContent className="bg-brand-dark border-brand-yellow">
        <DialogHeader>
          <DialogTitle className="text-brand-yellow">Edit profile</DialogTitle>
          <DialogDescription>
            Edit the list title, or mark/unmark it as public or private.
          </DialogDescription>
        </DialogHeader>

        <label className="text-stone-400 text-md w-full flex flex-col">
          <span className="mb-1">Title</span>
          <input
            autoFocus={false}
            autoComplete="off"
            type="text"
            placeholder={list.list.name ?? "List title"}
            value={title}
            onChange={(e) => setTitle(e.target.value.slice(0, 50))}
            className="
              p-2 bg-brand-darker border-[1px] border-brand-yellow border-opacity-80 text-opacity-80 rounded-xl
              text-brand-light placeholder:text-stone-500
            "
          />
        </label>

        <label className="text-stone-400 text-md w-full flex flex-col">
          <span className="mb-1">Visibility</span>
          <Select defaultValue={visibility} onValueChange={v => setVisibility(v)}>
            <SelectTrigger className="h-10 text-brand-yellow bg-brand-darker border-brand-yellow rounded-none">
              <SelectValue placeholder={visibility} />
            </SelectTrigger>
            <SelectContent className="bg-brand-darker border-brand-yellow">
              <SelectGroup>
                {["Public", "Private"].map((filter) => (
                  <SelectItem
                    key={filter}
                    value={filter}
                    className="hover:cursor-pointer focus:bg-brand-dark text-brand-yellow focus:text-brand-yellow"
                  >
                    {filter}
                  </SelectItem>
                ))}
              </SelectGroup>
            </SelectContent>
          </Select>
        </label>

        <DialogFooter>
          <Button
            onClick={updateList}
            disabled={title.length < 1}
            className="border border-brand-yellow hover:bg-brand-yellow mt-2"
          >
            Save changes
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

"use client";

import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger
} from "@/components/ui/alert-dialog";
import { useUserStore } from "@/app/state";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { deleteListAction } from "@/actions/delete-list-action";
import { DetailedList } from "@/models/models";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

export default function DeleteListDialog({ list }: { list: DetailedList }) {
  const { user } = useUserStore();
  const { toast } = useToast();

  const router = useRouter();

  const deleteList = async () => {
    const result = await deleteListAction(user.csrf, list.list.id);
    if (!result.success) return errorToast(toast, "Failed to delete list", result.error);
    router.replace("/profile/me/lists");
  };

  return (
    <AlertDialog>
      <AlertDialogTrigger>
        <Button
          variant="outline"
          className="
            text-brand-red border-brand-red w-[95px] sm:w-[120px] h-[32px] md:h-[40px] bg-transparent
            hover:bg-brand-red text-sm sm:text-base border-2
          "
        >
          Delete List
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent className="w-200px sm:w-full container bg-brand-darker border-brand-yellow">
        <AlertDialogHeader>
          <AlertDialogTitle className="text-brand-yellow">
            Are you absolutely sure?
          </AlertDialogTitle>
          <AlertDialogDescription>
            This action cannot be undone. This will permanently delete this list.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel className="bg-brand-yellow border-none">Cancel</AlertDialogCancel>
          <AlertDialogAction onClick={deleteList} className="bg-brand-red hover:bg-brand-dark">
            Continue
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}


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
import deleteUserAction from "@/actions/delete-user-action";
import { Button } from "@/components/ui/button";
import { useToast } from "@/components/ui/use-toast";
import { errorToast } from "@/lib/utils";

export default function DeleteProfileDialog() {
  const { user, clearUser } = useUserStore();
  const { toast } = useToast();

  const router = useRouter();

  const deleteUser = async () => {
    const result = await deleteUserAction(user.csrf);
    if (!result.success) return errorToast(toast, "Failed to delete profile", result.error);

    clearUser();
    router.push("/");
  }

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
          Delete Profile
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent className="w-200px sm:w-full container bg-brand-darker border-brand-yellow">
        <AlertDialogHeader>
          <AlertDialogTitle className="text-brand-yellow">
            Are you absolutely sure?
          </AlertDialogTitle>
          <AlertDialogDescription>
            This action cannot be undone. This will permanently delete your account.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel className="bg-brand-yellow border-none">Cancel</AlertDialogCancel>
          <AlertDialogAction onClick={deleteUser} className="bg-brand-red hover:bg-brand-dark">
            Continue
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}


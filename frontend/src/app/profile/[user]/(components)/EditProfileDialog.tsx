"use client";

import { z } from "zod";
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
import { ChangeEvent, FormEvent, useRef, useState } from "react";
import { useUserStore } from "@/app/state";
import updateUserAction from "@/actions/update-user-action";

const formSchema = z.object({
  username: z.string()
    .min(3, "username must be at least 3 characters")
    .max(32, "username must be at most 32 characters")
    .optional(),
  display_name: z.string()
    .min(3, "display name must be at least 3 characters")
    .max(32, "display name must be at most 32 characters")
    .optional(),
  password: z.string()
    .min(8, "password must be at least 8 characters")
    .max(32, "password must be at most 32 characters")
    .regex(new RegExp(`[A-Z]`), `password must contain at least one uppercase letter`)
    .regex(new RegExp(`[0-9]`), `password must contain at least one digit`)
    .regex(new RegExp(`[!@#$%^&*()_+{}|:<>?~]`), `password must contain at least one special character`).optional(),
  profile_picture: z.string().url("profile picture must be a valid image URL").optional(),
});
type FormSchema = z.infer<typeof formSchema>;

export default function EditProfileDialog() {
  const { user, setUser } = useUserStore();

  const dialogRef = useRef<HTMLButtonElement>(null);
  const [error, setError] = useState<string | null>(null);
  const [form, setForm] = useState<FormSchema>({
    username: undefined,
    display_name: undefined,
    password: undefined,
    profile_picture: undefined,
  });

  const onChange = (key: keyof FormSchema) => (e: ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [key]: e.target.value });
  }

  const updateUser = async (e: FormEvent) => {
    e.preventDefault();

    const values = formSchema.safeParse(form);
    if (!values.success) {
      setError(values.error.errors[0].message);
      return;
    }
    setError(null);

    const { username, display_name, password, profile_picture } = values.data;

    if (
      username === "" &&
      display_name === "" &&
      password === "" &&
      profile_picture === ""
    ) {
      setError("no fields to update");
      return;
    }

    const result = await updateUserAction(user.csrf, { username, display_name, password, profile_picture });
    if (!result.success) {
      setError(result.error);
      return;
    }

    setUser({ ...user, ...result.data.user });
    dialogRef.current?.click();
  }

  return (
    <Dialog>
      <DialogTrigger asChild ref={dialogRef}>
        <Button
          variant="outline"
          className="
            text-brand-yellow border-brand-yellow w-[80px] sm:w-[100px]  h-[32px] md:h-[40px] bg-transparent
            hover:bg-brand-yellow text-sm sm:text-base
          "
        >
          Edit Profile
        </Button>
      </DialogTrigger>
      <DialogContent className="bg-brand-dark border-brand-yellow">
        <DialogHeader>
          <DialogTitle className="text-brand-yellow">Edit profile</DialogTitle>
          <DialogDescription>
            Make changes to your profile here. Click save when you're done.
          </DialogDescription>
        </DialogHeader>
        <label className="text-stone-400 text-md flex gap-1">
          Username
        </label>
        <input
          autoComplete="off"
          type="text"
          placeholder={user.username}
          value={form.username}
          onChange={onChange("username")}
          className="
            p-2 text-lg bg-brand-darker border-[1px] border-brand-yellow border-opacity-80 text-opacity-80 rounded-xl
            text-brand-light placeholder:text-stone-500
          "
        />

        <label className="text-stone-400 text-md flex gap-1">
          Display Name
        </label>
        <input
          autoComplete="off"
          type="text"
          placeholder={user.display_name}
          value={form.display_name}
          onChange={onChange("display_name")}
          className="
            p-2 text-lg bg-brand-darker border-[1px] border-brand-yellow border-opacity-80 text-opacity-80 rounded-xl
            text-brand-light placeholder:text-stone-500
          "
        />

        <label className="text-stone-400 text-md flex gap-1">
          Password
        </label>
        <input
          autoComplete="off"
          type="text"
          placeholder="Password..."
          value={form.password}
          onChange={onChange("password")}
          className="
            p-2 text-lg bg-brand-darker border-[1px] border-brand-yellow border-opacity-80 text-opacity-80 rounded-xl
            text-brand-light placeholder:text-stone-500
          "
        />

        <label className="text-stone-400 text-md flex gap-1">
          Profile Picture
        </label>
        <input
          autoComplete="off"
          type="text"
          placeholder="Profile Picture..."
          value={form.profile_picture}
          onChange={onChange("profile_picture")}
          className="
            p-2 text-lg bg-brand-darker border-[1px] border-brand-yellow border-opacity-80 text-opacity-80 rounded-xl
            text-brand-light placeholder:text-stone-500
          "
        />

        {error && <div className="text-center text-sm text-red-500 mt-3">{error}</div>}

        <DialogFooter>
          <Button
            onClick={updateUser}
            className="border border-brand-yellow hover:bg-brand-yellow mt-2"
          >
            Save changes
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

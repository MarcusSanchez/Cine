"use client";

import { z } from "zod";
import React, { ChangeEvent, FormEvent, useState } from "react";
import { registerAction } from "@/actions/register-action";
import { Button } from "@/components/ui/button";
import { KeyRound, Mail, Monitor, User, UserPlus } from "lucide-react";
import Link from "next/link";
import { useUserStore } from "@/app/state";
import { useRouter } from "next/navigation";
import { getCookie } from "@/lib/utils";

const defaultPFP = "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png";

const formSchema = z.object({
  username: z.string()
    .min(3, "username must be at least 3 characters")
    .max(32, "username must be at most 32 characters"),
  display_name: z.string()
    .min(3, "display name must be at least 3 characters")
    .max(32, "display name must be at most 32 characters"),
  email: z.string()
    .min(3, "email must be at least 3 characters")
    .max(254, "email must be at most 254 characters")
    .email("email must be a valid email address"),
  password: z.string()
    .min(8, "password must be at least 8 characters")
    .max(32, "password must be at most 32 characters")
    .regex(new RegExp(`[A-Z]`), `password must contain at least one uppercase letter`)
    .regex(new RegExp(`[0-9]`), `password must contain at least one digit`)
    .regex(new RegExp(`[!@#$%^&*()_+{}|:<>?~]`), `password must contain at least one special character`),
  profile_picture: z.string().url("profile picture must be a valid image URL"),
});
type FormSchema = z.infer<typeof formSchema>;

export default function Register() {
  const { user, setUser } = useUserStore();
  const router = useRouter();

  const [error, setError] = useState<string | null>(null);
  const [form, setForm] = useState<FormSchema>({
    username: "",
    display_name: "",
    email: "",
    password: "",
    profile_picture: defaultPFP,
  });

  if (user.loggedIn) router.push("/");

  const handleChange = (key: keyof FormSchema) => (e: ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [key]: e.target.value });
  }

  const register = async (e: FormEvent) => {
    e.preventDefault();

    const values = formSchema.safeParse(form);
    if (!values.success) {
      setError(values.error.errors[0].message);
      return;
    }
    setError(null);

    const result = await registerAction(values.data);
    if (!result.success) {
      setError(result.error);
      return;
    }

    setUser({ ...result.data.user, loggedIn: true, csrf: getCookie("csrf")! })
    router.push("/");
  };

  return (
    <div className="container w-[400px]">
      <form onSubmit={register}>
        <h1 className="text-3xl text-center font-bold text-brand-light mb-1">Get Started!</h1>
        <p className="text-base text-center font-semibold text-brand-yellow mb-4">
          Gain access to all the best cinema.
        </p>

        <hr className="border-brand-yellow border-b-1 border-opacity-80 mb-4" />

        <div className="flex flex-col">
          <label className="text-stone-400 text-md p-1 flex gap-1">
            <User className="h-[20px] w-[20px]" />
            Username
          </label>
          <input
            type="text"
            placeholder="Username..."
            onChange={handleChange("username")}
            className="
              mb-3 p-2 text-lg bg-brand-darker border-[1px] border-brand-yellow border-opacity-80 text-opacity-80 rounded-xl
              text-brand-light placeholder:text-stone-500
            "
          />

          <label className="text-stone-400 text-md p-1 flex gap-1">
            <Monitor className="h-[20px] w-[20px]" />
            Display Name
          </label>
          <input
            type="text"
            placeholder="Display Name..."
            value={form.display_name}
            onChange={handleChange("display_name")}
            className="
              mb-3 p-2 text-lg bg-brand-darker border-[1px] border-brand-yellow border-opacity-80 text-opacity-80 rounded-xl
              text-brand-light placeholder:text-stone-500
            "
          />

          <label className="text-stone-400 text-md p-1 flex gap-1">
            <Mail className="h-[20px] w-[20px]" />
            Email
          </label>
          <input
            type="email"
            placeholder="Email..."
            value={form.email}
            onChange={handleChange("email")}
            className="
              mb-3 p-2 text-lg bg-brand-darker border-[1px] border-brand-yellow border-opacity-80 text-opacity-80 rounded-xl
              text-brand-light placeholder:text-stone-500
            "
          />

          <label className="text-stone-400 text-md p-1 flex gap-1">
            <KeyRound className="h-[20px] w-[20px]" />
            Password
          </label>
          <input
            type="password"
            placeholder="Password..."
            value={form.password}
            onChange={handleChange("password")}
            className="
              mb-2 p-2 text-lg bg-brand-darker border-[1px] border-brand-yellow border-opacity-80 text-opacity-80 rounded-xl
              text-brand-light placeholder:text-stone-500
            "
          />

          {error && <div className="text-center text-sm text-red-500 mt-3">{error}</div>}

          <Button
            type="submit"
            className="my-4 text-lg font-bold text-brand-darker bg-brand-yellow hover:bg-brand-light"
          >
            <UserPlus strokeWidth="2.5px" className="mr-1" />
            Register
          </Button>

          <hr className="border-brand-yellow border-b-1 border-opacity-80 mb-4" />

          <p className="text-sm text-center text-brand-light mb-4">
            Already have an account? {" "}
            <Link
              href={"/login"}
              className="text-brand-yellow hover:underline"
            >
              Click here to login.
            </Link>
          </p>
        </div>
      </form>
    </div>
  );
}


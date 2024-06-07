"use client";

import { z } from "zod";
import React, { ChangeEvent, FormEvent, useState } from "react";
import { Button } from "@/components/ui/button";
import { KeyRound, LogIn, User } from "lucide-react";
import { loginAction } from "@/actions/login-action";
import Link from "next/link";
import { useUserStore } from "@/app/state";
import { useRouter } from "next/navigation";
import { getCookie } from "@/lib/utils";

const formSchema = z.object({
  username: z.string()
    .min(3, "username must be at least 3 characters")
    .max(32, "username must be at most 32 characters"),
  password: z.string()
    .min(8, "password must be at least 8 characters")
    .max(32, "password must be at most 32 characters")
    .regex(new RegExp(`[A-Z]`), `password must contain at least one uppercase letter`)
    .regex(new RegExp(`[0-9]`), `password must contain at least one digit`)
    .regex(new RegExp(`[!@#$%^&*()_+{}|:<>?~]`), `password must contain at least one special character`),
});
type FormSchema = z.infer<typeof formSchema>;

export default function Register() {
  const { user, setUser } = useUserStore();
  const router = useRouter();

  const [error, setError] = useState<string | null>(null);
  const [form, setForm] = useState<FormSchema>({ username: "", password: "" });

  if (user.loggedIn) router.push("/");

  const onChange = (key: keyof FormSchema) => (e: ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [key]: e.target.value });
  }

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault();

    const values = formSchema.safeParse(form);
    if (!values.success) {
      setError(values.error.errors[0].message);
      return;
    }
    setError(null);

    const { username, password } = values.data;
    const result = await loginAction(username, password);
    if (!result.success) return setError(result.error);

    setUser({ ...result.data.user, loggedIn: true, csrf: getCookie("csrf")! })
    router.push("/");
  };

  return (
    <div className="container w-[400px]">
      <form onSubmit={onSubmit}>
        <h1 className="text-3xl text-center font-bold text-brand-light mb-1">Welcome back!</h1>
        <p className="text-base text-center font-semibold text-brand-yellow mb-4">
          Did you miss us?
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
            value={form.username}
            onChange={onChange("username")}
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
            onChange={onChange("password")}
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
            <LogIn strokeWidth="2.5px" className="mr-1" />
            Login
          </Button>

          <hr className="border-brand-yellow border-b-1 border-opacity-80 mb-4" />

          <p className="text-sm text-center text-brand-light mb-4">
            Don't have an account? {" "}
            <Link
              href={"/register"}
              className="text-brand-yellow hover:underline"
            >
              Click here to register.
            </Link>
          </p>

        </div>
      </form>
    </div>
  );
}


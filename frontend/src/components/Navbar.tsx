import Logo from '@/../public/logo.png';
import { Button, buttonVariants } from "@/components/ui/button";
import Link from 'next/link';
import { LogOut, Search, User } from "lucide-react";
import { cn } from "@/lib/utils";
import { useUserStore } from "@/app/state";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import React from "react";

export default function Navbar() {
  const { user } = useUserStore();

  return (
    <>
      <nav className="container flex justify-between py-2 max-w-[1200px] my-1">
        <Link href="/" className="flex items-center">
          <img src={Logo.src} alt="Cine" className="w-20 sm:w-24 h-min" />
        </Link>

        <div className="flex items-center gap-1 md:gap-2">
          <Button
            className="
            flex justify-start
            p-2 text-sm md:text-base text-stone-400 bg-brand-darker text-opacity-80 w-32 md:w-40 h-8 rounded-xl
            hover:ring-[1px] hover:ring-brand-yellow hover:ring-opacity-80 hover:text-brand-light
          "
          >
            <Search className="h-[12px] md:h-[16px]" />
            Search...
          </Button>
          {!user.loggedIn &&
            <>
              <Link
                href={"/login"}
                className={cn(
                  buttonVariants({ variant: "link" }),
                  "text-sm md:text-base text-brand-light hover:text-brand-yellow w-16",
                )}
              >
                Login
              </Link>
              <Link
                href={"/register"}
                className={cn(
                  buttonVariants({ variant: "link" }),
                  "text-sm md:text-base text-brand-light hover:text-brand-yellow w-16",
                )}
              >
                Register
              </Link>
            </>
          }
          {user.loggedIn && <ProfileDropDown />}
        </div>
      </nav>
      <hr className="border-black mb-8" />
    </>
  );
}

export function ProfileDropDown() {
  const { user, logout } = useUserStore();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="ml-2" asChild>
        <img
          src={user.profile_picture}
          alt="Profile Picture"
          className="
            w-[2.25rem] h-[2.25rem] border-b-brand-darker rounded-full hover:cursor-pointer hover:ring-[2px]
            hover:ring-brand-yellow hover:ring-opacity-80
          "
        />
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56 bg-brand-darker border-brand-yellow">
        <DropdownMenuLabel className="text-brand-yellow">My Account</DropdownMenuLabel>
        <DropdownMenuSeparator className="bg-brand-yellow" />
        <Link href={"/profile"}>
          <DropdownMenuItem className="text-brand-yellow hover:cursor-pointer">
            <User className="mr-2 h-4 w-4" />
            <span>Profile</span>
            <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut>
          </DropdownMenuItem>
        </Link>
        <DropdownMenuItem onClick={logout} className="text-brand-yellow hover:cursor-pointer">
          <LogOut className="mr-2 h-4 w-4" />
          <span>Log out</span>
          <DropdownMenuShortcut>⇧⌘Q</DropdownMenuShortcut>
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
import * as React from "react";
import { Slot } from "@radix-ui/react-slot";
import { cva, type VariantProps } from "class-variance-authority";

import { cn } from "@/lib/utils";
import Image from "next/image";

import GoogleImage from "@/assets/google.png";
import GithubImage from "@/assets/github.svg";

const buttonVariants = cva(
  "inline-flex uppercase items-center font-grotesk justify-center whitespace-nowrap rounded-sm text-base font-semibold ring-offset-white transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-slate-950 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 dark:ring-offset-slate-950 dark:focus-visible:ring-slate-300",
  {
    variants: {
      variant: {
        default:
          "bg-white flex gap-2 items-center border border-[#18181B] text-[#18181B] hover:bg-slate-100 dark:bg-slate-800 dark:text-slate-50 dark:hover:bg-slate-800/80",
      },
      size: {
        default: "h-10 px-5 py-2",
        sm: "h-9 rounded-sm px-3",
        lg: "h-11 rounded-sm px-8",
        icon: "h-10 w-10",
      },
      socialMedia: {
        github: "github",
        google: "google",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "default",
    },
  }
);

export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonVariants> {
  asChild?: boolean;
}

const SocialLoginButton = React.forwardRef<HTMLButtonElement, ButtonProps>(
  (
    { className, variant, size, socialMedia, asChild = false, ...props },
    ref
  ) => {
    const Comp = asChild ? Slot : "button";
    return (
      <Comp
        className={cn(buttonVariants({ variant, size, className }))}
        ref={ref}
        {...props}
      >
        {socialMedia === "google" && (
          <Image src={GoogleImage} alt="Google" className="w-5" />
        )}

        {socialMedia === "github" && (
          <Image src={GithubImage} alt="Github" className="w-5" />
        )}

        {props.children}
      </Comp>
    );
  }
);
SocialLoginButton.displayName = "SocialLoginButton";

export { SocialLoginButton, buttonVariants };

import { Button } from "@/components/ui/button";
import { SocialLoginButton } from "@/components/ui/social-login";
import React from "react";

const Cta = () => {
  return (
    <>
      <h2 className="text-5xl font-bold">
        Express yourself.{" "}
        <span className="text-primary">Connect with others.</span>
      </h2>

      <div className="flex gap-4 items-center">
        <Button>Sign up with email</Button>

        <SocialLoginButton socialMedia="google">
          Sign up with google
        </SocialLoginButton>
      </div>

      <p>
        <span className="font-semibold">PostGo is free to use</span> for as long
        as you’ll live
      </p>
    </>
  );
};

export default Cta;

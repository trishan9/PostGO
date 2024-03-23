import React from "react";
import Navbar from "./_components/navbar";

const AuthLayout = ({ children }: { children: React.ReactElement }) => {
  return (
    <div>
      <Navbar />

      {children}
    </div>
  );
};

export default AuthLayout;

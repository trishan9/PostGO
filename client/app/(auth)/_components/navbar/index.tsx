"use client";

import { useWindowScroll } from "@uidotdev/usehooks";
import { cn } from "@/lib/utils";
import NavLogo from "./nav-logo";
import NavActions from "./nav-actions";

const Navbar = () => {
  const [{ y }] = useWindowScroll();

  return (
    <nav
      className={cn(
        "flex justify-between items-center",
        (y as number) >= 288
          ? "bg-white rounded-full border shadow-xl z-[100] py-4 px-8 w-[1290px] top-2 left-0 sticky"
          : "w-full py-6 px-80"
      )}
    >
      <div className="flex items-center gap-14">
        <NavLogo />
      </div>

      <NavActions primaryLabel="Create Post" secondaryLabel="Settings" />
    </nav>
  );
};

export default Navbar;

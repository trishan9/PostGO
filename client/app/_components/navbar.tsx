import { Button } from "@/components/ui/button";
import Link from "next/link";

const Navbar = () => {
  return (
    <nav className="flex w-full p-4 px-6 border bg-slate-200 justify-between items-center">
      <h1 className="text-xl font-medium font-grotesk">
        <Link href="/">PostGO</Link>
      </h1>

      <ul className="flex items-center gap-4">
        <li>
          <Button variant="ghost" size="sm">
            Login
          </Button>
        </li>

        <li>
          <Button size="sm">Signup</Button>
        </li>
      </ul>
    </nav>
  );
};

export default Navbar;

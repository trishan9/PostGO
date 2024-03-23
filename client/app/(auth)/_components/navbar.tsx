import { Button } from "@/components/ui/button";

const Navbar = () => {
  return (
    <nav className="flex w-full p-4 px-6 border bg-slate-200 justify-between items-center">
      <h1 className="text-xl font-medium font-grotesk">PostGO</h1>

      <ul className="flex items-center gap-4">
        <li>
          <Button variant="ghost" size="sm">
            Posts
          </Button>
        </li>
      </ul>
    </nav>
  );
};

export default Navbar;

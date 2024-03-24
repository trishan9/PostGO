import { Button } from "@/components/ui/button";

type NavActionsProps = {
  primaryLabel: string;
  secondaryLabel: string;
};

const NavActions = ({ primaryLabel, secondaryLabel }: NavActionsProps) => {
  return (
    <ul className="flex items-center gap-4">
      <li>
        <Button variant="secondary">{secondaryLabel}</Button>
      </li>

      <li>
        <Button>{primaryLabel}</Button>
      </li>
    </ul>
  );
};

export default NavActions;

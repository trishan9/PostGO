const MENU_ITEMS = [
  "Features",
  "Solutions",
  "Enterprise",
  "Resources",
  "Pricing",
];

const NavMenus = () => {
  return (
    <ul className="flex items-center gap-6 text-black font-semibold text-sm">
      {MENU_ITEMS.map((item) => (
        <li className="cursor-pointer" key={item}>
          {item}
        </li>
      ))}
    </ul>
  );
};

export default NavMenus;

import Image from "next/image";
import Link from "next/link";

const META_DATA = {
  LOGO: "/logo.webp",
  LOGO_ALT: "Logo",
  LOGO_WIDTH: 506,
  LOGO_HEIGHT: 329,
  APP_NAME: "PostGO",
} as const;

const NavLogo = () => {
  return (
    <Link href="/">
      <div className="flex gap-2 items-center">
        <Image
          src={META_DATA.LOGO}
          alt={META_DATA.LOGO_ALT}
          width={META_DATA.LOGO_WIDTH}
          height={META_DATA.LOGO_HEIGHT}
          className="w-16"
        />

        <h1 className="text-xl font-medium font-grotesk">
          {META_DATA.APP_NAME}
        </h1>
      </div>
    </Link>
  );
};

export default NavLogo;

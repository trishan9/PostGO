import Image from "next/image";
import { cn } from "@/lib/utils";
import NextLogo from "@/assets/next.svg";
import ReactLogo from "@/assets/react.png";
import GoLogo from "@/assets/go.png";
import GinLogo from "@/assets/gin.png";
import PostgresLogo from "@/assets/postgresql.png";

const LOGOS = [
  {
    id: 1,
    alt: "gin",
    src: GinLogo,
    className: "w-9",
  },
  {
    id: 2,
    alt: "go",
    src: GoLogo,
    className: "w-16",
  },
  {
    id: 3,
    alt: "postgres",
    src: PostgresLogo,
    className: "w-10",
  },
  {
    id: 4,
    alt: "next",
    src: NextLogo,
    className: "w-10",
  },
  {
    id: 5,
    alt: "react",
    src: ReactLogo,
    className: "w-10",
  },
];

const LogoSection = () => {
  return (
    <div className="flex items-center gap-6 py-4">
      {LOGOS.map((logo) => (
        <Image
          key={logo.id}
          src={logo.src}
          alt={logo.alt}
          className={cn(logo.className)}
        />
      ))}
    </div>
  );
};

export default LogoSection;

import Image from "next/image";
import DemoImage from "@/assets/demo.png";
import LogoSection from "./logos";
import Cta from "./hero-cta";

const Hero = () => {
  return (
    <section className="py-6 flex flex-col w-full justify-between items-center gap-4">
      <Cta />

      <LogoSection />

      <div className="w-full flex flex-col justify-between items-center relative">
        <Image
          src={DemoImage}
          alt="demo"
          className="w-[900px] border rounded-md"
        />

        <div className="bg-gray-50 w-full h-[1000px] absolute -z-10 top-72"></div>
      </div>
    </section>
  );
};

export default Hero;

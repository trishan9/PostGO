import Hero from "./_components/hero";
import Navbar from "./_components/navbar";

export default function Home() {
  return (
    <main className="flex min-h-screen font-primary flex-col items-center">
      <Navbar />

      <Hero />
    </main>
  );
}

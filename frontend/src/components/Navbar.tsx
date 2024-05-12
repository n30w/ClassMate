import Link from "next/link";

const Navbar: React.FC = () => {
  return (
    <nav>
      <div className="relative">
        <div className="absolute inset-0 opacity-70"></div>
        <div className="py-8 px-32">
          <div className="flex items-center gap-4">
            <h1 className="text-white text-2xl font-bold">
              <Link href={`/homepage`}>Darkspace</Link>
            </h1>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;

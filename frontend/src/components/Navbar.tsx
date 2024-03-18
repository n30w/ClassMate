import Image from "next/image";

const Navbar: React.FC = () => {
  return (
    <nav
      style={{
        backgroundImage: `url('/backgrounds/dashboard-bg.jpeg')`,
        backgroundSize: "cover",
        backgroundPosition: "center",
      }}
    >
      <div className="relative">
        <div className="absolute inset-0 bg-black opacity-70"></div>
        <div className="py-8 px-32">
          <div className="flex items-center gap-4">
            <Image
              src="/backgrounds/NYU-logo.png"
              width="150"
              height="39"
              alt="NYU Logo"
              className="z-10"
            />
            <Image
              src="/backgrounds/darkspace.png"
              width="200"
              height="39"
              alt="Darkspace Logo"
              className="z-10"
            />
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;

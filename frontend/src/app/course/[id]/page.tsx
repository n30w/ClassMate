import Image from "next/image";
import Announcements from "@/components/coursepage/Announcements";
import Assignments from "@/components/coursepage/Assignments";
import Discussions from "@/components/Discussions";

export default function Page() {
  return (
    <div style={{ backgroundColor: "black", minHeight: "100vh" }}>
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
      <div
        style={{
          backgroundImage: `url('/backgrounds/course-bg.jpg')`,
          backgroundSize: "cover",
          backgroundPosition: "center",
          width: "100%",
          height: "300px",
        }}
      >
        <div className="relative">
          <div className="py-4 px-8 ml-32 mt-32 h-32 w-96 absolute bg-black bg-opacity-70 flex flex-col justify-center">
            <h1 className="text-white text-3xl font-bold pb-2 block text-opacity-100">
              Software Engineering
            </h1>
            <h2 className="text-white text-2xl block text-opacity-100">
              with, Xu Lihua
            </h2>
          </div>
          <div className="flex justify-end">
            <Discussions />
          </div>
        </div>
      </div>
      <div className="flex justify-around p-16">
        <div className="flex flex-col w-96">
          <Announcements />
        </div>
        <div className="flex flex-col">
          <Assignments />
        </div>
      </div>
    </div>
  );
}

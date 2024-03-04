import Image from "next/image";
import Courses from "./components/Courses";

export default function Home() {

  const currentDate = new Date().toLocaleDateString();

  return (
    <div>
      <nav
        style={{
          backgroundImage: `url('/backgrounds/dashboard-bg.png')`,
          backgroundSize: 'cover',
          backgroundPosition: 'center'
        }}
      >
        <div className="relative">
          <div className="absolute inset-0 bg-black opacity-70"></div>
          <div className="py-8 px-32">
            <div className="flex items-center gap-16">
              <Image src="/backgrounds/NYU-logo.png" width="100" height="39" alt="NYU Logo" className="z-10"/>
              <h3 className="text-white z-10 font-light">Calendar</h3>
              <h3 className="text-white z-10 font-light">Announcements</h3>
              <h3 className="text-white z-10 font-light">Messages</h3>
            </div>
            <div className="flex items-center justify-between py-8">
              <h1 className="text-white font-bold text-5xl z-10">Welcome to Darkspace!</h1>
              <h3 className="text-white text-xl z-10">{currentDate}</h3>
            </div>
          </div>
        </div>
      </nav>
      <div className="bg-white">
        <div className="flex items-center justify-between py-8 px-32">
          <h1 className="font-bold text-2xl">Current Courses</h1>
          <h2 className="font-light text-xl">Spring 2024</h2>
        </div>
        <Courses 
          coursename="Software Engineering"
          professor="Xu, Lihua"
          time="Tue,Thu 3.45 PM - 5.00 PM"
          loc="Room S311"
        />
        <Courses 
          coursename="Software Engineering"
          professor="Xu, Lihua"
          time="Tue,Thu 3.45 PM - 5.00 PM"
          loc="Room S311"
        />
      </div>
    </div>
  );
}

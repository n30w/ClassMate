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
        <div class="relative">
          <div class="absolute inset-0 bg-black opacity-70"></div>
          <div class="py-8 px-32">
            <div class="flex items-center gap-16">
              <Image src="/backgrounds/NYU-logo.png" width="100" height="39" alt="NYU Logo" class="z-10"/>
              <h3 class="text-white z-10 font-light">Calendar</h3>
              <h3 class="text-white z-10 font-light">Announcements</h3>
              <h3 class="text-white z-10 font-light">Messages</h3>
            </div>
            <div class="flex items-center justify-between py-8">
              <h1 class="text-white font-bold text-5xl z-10">Welcome to Darkspace!</h1>
              <h3 class="text-white text-xl z-10">{currentDate}</h3>
            </div>
          </div>
        </div>
      </nav>
      <div class="bg-white">
        <div class="flex items-center justify-between py-8 px-32">
          <h1 class="font-bold text-2xl">Current Courses</h1>
          <h2 class="font-light text-xl">Spring 2024</h2>
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

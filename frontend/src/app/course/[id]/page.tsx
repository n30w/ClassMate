import Announcements from "@/components/coursepage/Announcements";
import Assignments from "@/components/coursepage/Assignments";
import Discussions from "@/components/coursepage/Discussions";
import { Announcement, Assignment, Discussion, Course } from "@/lib/types";
import Navbar from "@/components/Navbar";

// These data names must match what the API returns.
interface HomepageData {
  name: string;
  teacher_name: string;
  assignments: Assignment[];
  discussions: Discussion[];
  announcements: Announcement[];
}

async function getData(id: string): Promise<HomepageData> {
  const path = `http://localhost:6789/v1/course/homepage/${id}`;

  const res = await fetch(path, {
    method: "POST",
    body: JSON.stringify({
      courseid: id,
    }),
  });

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  } else {
    const data: Course = await res.json();
    console.log(data);
  }

  return res.json();
}

// Dynamic route example found here:
// https://nextjs.org/docs/app/building-your-application/routing/dynamic-routes#example
export default async function Page({ params }: { params: { slug: string } }) {
  const data = await getData(params.slug);

  return (
    <div style={{ backgroundColor: "black", minHeight: "100vh" }}>
      <Navbar />
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
              {data.name}
            </h1>
          </div>
          <div className="flex justify-end">
            <Discussions />
          </div>
        </div>
      </div>
      <div className="flex justify-around p-16">
        <div className="flex flex-col w-96">
          <Announcements entries={data.discussions} />
        </div>
        <div className="flex flex-col">
          <Assignments entries={data.assignments} />
        </div>
      </div>
    </div>
  );
}

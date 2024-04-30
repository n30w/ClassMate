import Announcements from "@/components/coursepage/Announcements";
import Assignments from "@/components/coursepage/Assignments";
import Discussions from "@/components/coursepage/Discussions";
import { Announcement, Assignment, Discussion } from "@/lib/types";
import Navbar from "@/components/Navbar";

// These data names must match what the API returns.
interface HomepageData {
  name: string;
  teacher_name: string;
  assignments: Assignment[];
  discussions: Discussion[];
  announcements: Announcement[];
}

// This function is adapted from:
// https://nextjs.org/docs/app/building-your-application/data-fetching/fetching-caching-and-revalidating#fetching-data-on-the-server-with-fetch
// async function getData(of: string): Promise<HomepageData> {
//   const path = `http://localhost:6789/v1/course/homepage/${of}`;
//   console.log(path);

//   const res = await fetch(path, {
//     method: "GET",
//     headers: {
//       "Content-Type": "application/json",
//       "Access-Control-Allow-Origin": "*",
//       "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
//       "Access-Control-Allow-Headers": "Content-Type, Authorization",
//     },
//   });

//   if (!res.ok) {
//     // This will activate the closest `error.js` Error Boundary
//     throw new Error("Failed to fetch data");
//   }

//   return res.json();
// }

async function getData(): Promise<HomepageData> {
  const path =
    "http://localhost:6789/v1/course/homepage/c3b34a9f-8f59-4818-a684-9cda56f42d02";

  const res = await fetch(path, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
      "Access-Control-Allow-Headers": "Content-Type, Authorization",
    },
  });

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error("Failed to fetch data");
  }

  return res.json();
}

// Dynamic route example found here:
// https://nextjs.org/docs/app/building-your-application/routing/dynamic-routes#example
export default async function Page({ params }: { params: { slug: string } }) {
  // const data = await getData(params.slug);
  const data = await getData();

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
              {data.course.name}
            </h1>
          </div>
          <div className="flex justify-end">
            <Discussions />
          </div>
        </div>
      </div>
      <div className="flex justify-around p-16">
        <div className="flex flex-col w-96">
          <Announcements entries={data.course.discussions} />
        </div>
        <div className="flex flex-col">
          <Assignments entries={data.course.assignments} />
        </div>
      </div>
    </div>
  );
}

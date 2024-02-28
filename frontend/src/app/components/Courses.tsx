import Announcements from "./Announcements";
import Assignments from "./Assignments";
import Discussions from "./Discussions";
import Image from "next/image";

interface CourseProps {
  coursename: string;
  professor: string;
  time: string;
  loc: string;
}

const Course: React.FC<CourseProps> = (props) => {
  return (
    <div className="py-4 px-32">
      <div className="flex border border-gray-300">
        <div className="relative flex-col">
          <Image
            src="/backgrounds/course-bg.png"
            width="340"
            height="340"
            alt="Course Background"
          />
          <div className="absolute inset-0 flex flex-col justify-end items-start text-white p-4">
            <h1 className="text-white font-bold text-2xl">
              {props.coursename}
            </h1>
            <h2 className="text-white mb-2">with {props.professor}</h2>
            <h3 className="text-white font-light text-sm">{props.time}</h3>
            <h3 className="text-white font-light text-sm">{props.loc}</h3>
          </div>
        </div>
        <div className="flex my-8 mx-4 justify-between">
          <div className="flex flex-col">
            <Announcements />
          </div>
          <div className="flex flex-col">
            <Assignments />
          </div>
          <div className="flex flex-col">
            <Discussions />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Course;

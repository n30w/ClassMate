import { Course } from "@/lib/types";
import Image from "next/image";

interface props {
  data: Course;
  onClick: () => void;
  banner: any;
}

const CourseItem: React.FC<props> = (props: props) => {
  console.log("BANNER PATH: ", props.data.banner);
  return (
    <div className="py-4 px-32" onClick={props.onClick}>
      <div>
        <div className="relative flex-col h-80 w-80">
          <Image
            src={props.banner}
            className="object-cover w-full h-full"
            width={340}
            height={340}
            alt="Course Background"
          />
          <div className="absolute inset-0 flex flex-col justify-end items-start text-white p-4">
            <h1 className="text-white font-bold text-2xl">{props.data.name}</h1>
            {/* <h2 className="text-white mb-2">with {props.data.professor}</h2> */}
            {/* <h3 className="text-white font-light text-sm">
              {props.data.location}
            </h3> */}
          </div>
        </div>
        {/* <div className="flex my-8 mx-4 justify-between">
          <div className="flex flex-col">
            <Announcements />
          </div>
          <div className="flex flex-col">
            <Assignments />
          </div>
          <div className="flex flex-col">
            <Discussions />
          </div>
        </div> */}
      </div>
    </div>
  );
};

export default CourseItem;

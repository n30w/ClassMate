import { Announcement } from "@/lib/types";

interface props {
  announcements: Announcement[];
}

const AnnouncementDisplay: React.FC<props> = ({ announcements }: props) => {
  {
    /* Checks if announced is null. This won't work by checking if the array is greater than zero, because announced is a promise. */
  }
  console.log(announcements);
  return (
    <ul className="w-full bg-red-400 h-full">
      {announcements ? (
        announcements.map((announcement: Announcement, key: number) => (
          <li
            className="flex flex-col p-6 h-46 border shadow bg-gray-900 border-gray-700 hover:bg-gray-700"
            key={key}
          >
            <h2 className="text-white text-3xl mb-1">{announcement.title}</h2>
            <h3 className="text-white text-sm mb-2">{announcement.date}</h3>
            <p className="text-white text-lg font-light">
              {announcement.description}
            </p>
          </li>
        ))
      ) : (
        <p className="text-white">No announcements yet.</p>
      )}
      ;
    </ul>
  );
};

export default AnnouncementDisplay;

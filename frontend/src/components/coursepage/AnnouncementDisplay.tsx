const AnnouncementDisplay: React.FC<{ announcements: any }> = ({
  announcements,
}) => {
  {
    /* Checks if announced is null. This won't work by checking if the array is greater than zero, because announced is a promise. */
  }
  return announcements ? (
    announcements.map((announcement: any, key: number) => (
      <div key={key}>
        <h2 className="text-white text-2xl mb-1">{announcement.Title}</h2>
        {/* <h3 className="text-white text-sm mb-2">{announcement.date}</h3> */}
        <p className="text-white text-sm font-light border-b-2 border-white mb-4 pb-4">
          {announcement.Description}
        </p>
      </div>
    ))
  ) : (
    <p className="text-white">No announcements yet.</p>
  );
};

export default AnnouncementDisplay;

const AnnouncementDisplay: React.FC<{ announcements: any[] }> = ({
  announcements,
}) => {
  return announcements.length > 0 ? (
    announcements.map((announcement, key) => (
      <div key={key}>
        <h2 className="text-white text-2xl mb-1">{announcement.title}</h2>
        <h3 className="text-white text-sm mb-2">{announcement.date}</h3>
        <p className="text-white text-sm font-light border-b-2 border-white mb-4 pb-4">
          {announcement.description}
        </p>
      </div>
    ))
  ) : (
    <p className="text-white">No announcements yet.</p>
  );
};

export default AnnouncementDisplay;

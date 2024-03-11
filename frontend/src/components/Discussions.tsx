import React from "react";

const Discussions: React.FC = () => {
  const currentDate = new Date();

  const formattedDate = currentDate
    .toLocaleDateString("en-US", {
      month: "2-digit",
      day: "2-digit",
      year: "numeric",
    })
    .replace(/\//g, "-");

  return (
    <div className="bg-black bg-opacity-70 p-8 mt-8">
      <h2 className="text-white font-bold">
        Balancing Innovation and Maintenance
      </h2>
      <h3 className="text-white">Neo Alabastro posted on {formattedDate}</h3>
      <a
        href=""
        className="text-white font-light flex justify-end text-sm pt-2"
      >
        View Discussions
      </a>
    </div>
  );
};

export default Discussions;

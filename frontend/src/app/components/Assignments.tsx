const Assignments = () => {
  const homeworks = [
    { name: "Homework 1", duedate: "Feb 8" },
    { name: "Homework 2", duedate: "Feb 25" },
    { name: "Homework 3", duedate: "March 3" },
    { name: "Homework 4", duedate: "March 10" },
  ];

  return (
    <div className="w-60 h-70 px-4">
      <h1 className="font-bold text-xl border-b-2 border-black mb-4 pb-4">
        Assignments
      </h1>
      <ul>
        {homeworks.map((homework, key) => {
          return (
            <li key={key}>
              <h2 className="text-m font-bold">{homework.name}</h2>
              <p className="text-sm font-light pb-1 mb-1 border-b border-black">
                {homework.duedate}
              </p>
            </li>
          );
        })}
      </ul>
    </div>
  );
};

export default Assignments;

const formattedDate = (isoDateString: string) => {
  return new Date(isoDateString).toLocaleDateString("en-US", {
    day: "2-digit",
    month: "long",
    year: "numeric",
  });
  // .replace(/\//g, "-");
};

export default formattedDate;

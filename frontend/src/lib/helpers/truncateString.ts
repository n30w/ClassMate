// truncateString slices a string up to a specified length, then adds
// ellipses to the end of the string.
const truncateString = (s: string, end: number) => {
  if (s.length <= 50) return s;
  return s.slice(0, end) + "...";
};

export default truncateString;

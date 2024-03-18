const apiPath = (s: string) => {
  return `http://${process.env.API_HOSTNAME}:${process.env.API_PORT}${s}`;
};

export default apiPath;

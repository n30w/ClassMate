export interface Announcement {
  id: string;
  title: string;
  date: string;
  description: string;
}

export interface Assignment {
  id: string;
  title: string;
  duedate: string;
  description: string;
}

export interface Course {
  id: string;
  name: string;
  professor: string;
  location: string;
}

export interface Discussion {
  title: string;
  description: string;
}

export interface User {
  id: string;
  username: string;
  fullname: string;
}

export interface Token {
  authentication_token: { token: string };
  permissions: string;
}

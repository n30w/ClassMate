export interface Entity {
  id: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

export interface Announcement extends Entity {
  title: string;
  date: string;
  description: string;
}

export interface Assignment extends Entity {
  title: string;
  due_date: string;
  description: string;
}

export interface Course extends Entity {
  name: string;
  description: string;
  professor: string;
  banner: string;
  roster: string[];
  assignments?: Assignment[];
}

export interface Discussion {
  title: string;
  description: string;
}

export interface User extends Entity {
  username: string;
  full_name: string;
  email: string;
}

export interface Token {
  authentication_token: { token: string };
  permissions: string;
}

export interface Submission {
  userid: string;
  assignmentid: string;
  grade: string;
  feedback: string;
  file: any;
  file_path: string;
}

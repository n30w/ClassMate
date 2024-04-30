-- Initial setup
CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users Table
CREATE TABLE IF NOT EXISTS users (
                                     id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
                                     net_id VARCHAR UNIQUE,
                                     created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                     updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                     deleted_at TIMESTAMP WITHOUT TIME ZONE,
                                     full_name VARCHAR,
                                     profile_picture_id UUID,
                                     bio TEXT,
                                     username VARCHAR NOT NULL,
                                     password VARCHAR NOT NULL,
                                     email VARCHAR NOT NULL,
                                     membership INT NOT NULL
);

-- Courses Table
CREATE TABLE IF NOT EXISTS courses (
                                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                       title VARCHAR NOT NULL,
                                       description TEXT,
                                       teacher_id UUID[] REFERENCES users(net_id) ON DELETE CASCADE,
                                       created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                       updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                       archived BOOLEAN NOT NULL DEFAULT FALSE
);

-- Media Table
CREATE TABLE IF NOT EXISTS media (
                                     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                     type VARCHAR NOT NULL,
                                     url VARCHAR NOT NULL,
                                     created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                     updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

-- Projects Table
CREATE TABLE IF NOT EXISTS projects (
                                        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                        name VARCHAR NOT NULL,
                                        description TEXT,
                                        created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                        updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

-- Messages Table
CREATE TABLE IF NOT EXISTS messages (
                                        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                        title VARCHAR NOT NULL,
                                        description TEXT,
                                        date TIMESTAMP WITHOUT TIME ZONE,
                                        type BOOLEAN
);

-- Assignments Table
CREATE TABLE IF NOT EXISTS assignments (
                                           id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                           title VARCHAR NOT NULL,
                                           description TEXT,
                                           date TIMESTAMP WITHOUT TIME ZONE,
                                           due_date TIMESTAMP WITHOUT TIME ZONE
);

-- Submissions Table
CREATE TABLE IF NOT EXISTS submissions (
                                           id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                           file_type VARCHAR,
                                           submission_time TIMESTAMP WITHOUT TIME ZONE,
                                           on_time BOOLEAN,
                                           grade INT,
                                           feedback VARCHAR
);

-----------------
--- JUNCTIONS ---
-----------------

-- Junction Table for Users and Courses (Many-to-Many)
CREATE TABLE IF NOT EXISTS user_courses (
                                            user_net_id VARCHAR REFERENCES users(net_id) ON DELETE CASCADE,
                                            course_id UUID REFERENCES courses(id) ON DELETE CASCADE ,
                                            PRIMARY KEY (user_net_id, course_id)
);

-- Junction Table for Courses and Messages (Many-to-Many)
CREATE TABLE IF NOT EXISTS course_messages (
                                               course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
                                               message_id UUID REFERENCES messages(id) ON DELETE CASCADE ,
                                               PRIMARY KEY (course_id, message_id)
);

-- Junction Table for Courses and Teachers (Many-to-Many)
CREATE TABLE IF NOT EXISTS course_teachers (
                                               course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
                                               teacher_id VARCHAR REFERENCES users(net_id) ON DELETE CASCADE,
                                               PRIMARY KEY (course_id, teacher_id)
);

-- Junction Table for Courses and Roster (Students) (Many-to-Many)
CREATE TABLE IF NOT EXISTS course_roster (
                                             course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
                                             student_id VARCHAR REFERENCES users(net_id) ON DELETE CASCADE,
                                             PRIMARY KEY (course_id, student_id)
);

-- Junction Table for Courses and Assignments (Many-to-Many)
CREATE TABLE IF NOT EXISTS course_assignments (
                                                  course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
                                                  assignment_id UUID REFERENCES assignments(id) ON DELETE CASCADE,
                                                  PRIMARY KEY (course_id, assignment_id)
);

-- Junction Table for Assignments and Submissions (Many-to-Many)
CREATE TABLE IF NOT EXISTS assignment_submissions (
                                                      assignment_id UUID REFERENCES assignments(id) ON DELETE CASCADE,
                                                      submission_id UUID REFERENCES submissions(id) ON DELETE CASCADE,
                                                      PRIMARY KEY (assignment_id, submission_id)
);

-- Junction Table for Messages and Media (Many-to-Many)
CREATE TABLE IF NOT EXISTS message_media (
                                             message_id UUID REFERENCES messages(id) ON DELETE CASCADE,
                                             media_id UUID REFERENCES media(id) ON DELETE CASCADE,
                                             PRIMARY KEY (message_id, media_id)
);

-- Authentication Table
CREATE TABLE IF NOT EXISTS tokens (
                                      hash bytea PRIMARY KEY,
                                      net_id VARCHAR UNIQUE REFERENCES users(net_id) ON DELETE CASCADE,
                                      expiry timestamp(0) with time zone NOT NULL,
                                      scope text NOT NULL
);


-- Adding foreign key constraints after all tables are established and maintain direct single relationships
-- Use a cascade deletion.
ALTER TABLE projects ADD COLUMN user_net_id VARCHAR REFERENCES users(net_id) ON DELETE CASCADE;
ALTER TABLE assignments ADD COLUMN media_id UUID REFERENCES media(id) ON DELETE SET NULL;
ALTER TABLE assignments ADD COLUMN course_id UUID REFERENCES courses(id) ON DELETE SET NULL;
ALTER TABLE assignments ADD COLUMN owner_id INT REFERENCES users(id) ON DELETE SET NULL;
ALTER TABLE submissions ADD COLUMN user_id INT REFERENCES users(id) ON DELETE CASCADE;

-- Foreign key for profile picture which relates to the Media table
ALTER TABLE users ADD CONSTRAINT fk_profile_picture
    FOREIGN KEY (profile_picture_id)
        REFERENCES media (id)
        ON DELETE SET NULL;

-- Insert dummy users
-- Note: Insert users before courses since courses might reference users' net_id if needed
INSERT INTO users (net_id, created_at, updated_at, username, password, email, membership, full_name) VALUES
   ('abc123', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'jcena', 'password123', 'abc123@nyu.edu', 0, 'John Cena'),
   ('xyz789', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'mmiller', 'mypass789', 'xyz789@example.com', 1, 'Mike Miller'),
   ('def456', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'ajackson', 'pass456', 'def456@example.com', 0, 'Alice Jackson'),
   ('uvw321', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'ksmith', 'pass321', 'uvw321@example.com', 1, 'Kevin Smith'),
   ('ghi987', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'jdoe', 'mysecretpass', 'ghi987@example.com', 0, 'Jane Doe');

-- Insert dummy courses
-- Removed the net_id column since it's now intended to be managed through a junction table or direct reference in projects, not stored directly in courses
INSERT INTO courses (title, description, created_at, updated_at) VALUES
   ('Introduction to Computer Science', 'Basic concepts of computer programming', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
   ('Advanced Mathematics', 'In-depth coverage of calculus and linear algebra', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
   ('Modern Art History', 'Exploration of art from the 19th century to present', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
   ('Environmental Science', 'Study of climate change and environmental impact', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
   ('Business Management', 'Principles and practices in managing modern businesses', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Courses with preset ID for test operations.
INSERT INTO courses(id, title, description, created_at, updated_at) VALUES
   ('c3b34a9f-8f59-4818-a684-9cda56f42d02', 'Clown Foundations', 'Learn how to be a clown', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
   ('98e64e88-b989-49a0-bbfd-76e158bac634', 'Delete This Course', 'In testing, this course should be deleted', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);


-- Insert dummy assignments
-- Using the new structure without course_id in the initial insert. Instead, use the junction table to link courses and assignments if needed
INSERT INTO assignments (title, description, date, due_date) VALUES
   ('Quiz 1', 'Quiz on basic programming concepts', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + interval '7 days'),
   ('Calculus Exam', 'Midterm exam on calculus topics', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + interval '14 days'),
   ('Art Essay', 'Essay on modern art movements', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + interval '10 days'),
   ('Climate Report', 'Group report on climate change effects', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + interval '20 days'),
   ('Management Case Study', 'Analysis of a business case study', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + interval '12 days');

-- Insert dummy messages
-- Updated to avoid using ARRAY and using the junction table instead
INSERT INTO messages (title, description, date, type) VALUES
   ('Welcome!', 'Welcome to the course on Computer Science', CURRENT_TIMESTAMP, TRUE),
   ('Assignment Reminder', 'Remember to submit the calculus exam by Friday', CURRENT_TIMESTAMP, FALSE),
   ('Field Trip', 'Field trip to modern art museum next week', CURRENT_TIMESTAMP, TRUE),
   ('Guest Lecture', 'Upcoming guest lecture on renewable energy', CURRENT_TIMESTAMP, TRUE),
   ('Project Groups', 'Project groups for the case study have been assigned', CURRENT_TIMESTAMP, FALSE);

-- Insert dummy projects
-- Note: Assuming that user_net_id refers to a single user managing the project
INSERT INTO projects (name, description, created_at, updated_at, user_net_id) VALUES
   ('Database Design', 'Project focusing on designing efficient databases', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'abc123'),
   ('Statistics Software', 'Develop statistical software using Python', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'xyz789'),
   ('Art Exhibition', 'Organize a virtual art exhibition', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'def456'),
   ('Water Quality', 'Study on water quality in urban areas', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'uvw321'),
   ('Startup Plan', 'Create a business plan for a new startup', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'ghi987');

-- Insert dummy submissions
-- Note: Assuming each submission is linked to a single user
INSERT INTO submissions (user_id, file_type, submission_time, on_time, grade, feedback) VALUES
   ((SELECT id FROM users WHERE username = 'jcena'), 'PDF', CURRENT_TIMESTAMP, TRUE, 85, 'Good effort, but watch out for syntax errors'),
   ((SELECT id FROM users WHERE username = 'mmiller'), 'DOCX', CURRENT_TIMESTAMP, FALSE, 78, 'Late submission, but well-written content'),
   ((SELECT id FROM users WHERE username = 'ajackson'), 'PDF', CURRENT_TIMESTAMP, TRUE, 92, 'Excellent analysis and creativity'),
   ((SELECT id FROM users WHERE username = 'ksmith'), 'ZIP', CURRENT_TIMESTAMP, TRUE, 88, 'Good collaboration, impressive research'),
   ((SELECT id FROM users WHERE username = 'jdoe'), 'PDF', CURRENT_TIMESTAMP, TRUE, 90, 'Very thorough and well-structured report');

-- Inserting courses for John Cena into the user_courses junction table
INSERT INTO user_courses (user_net_id, course_id) VALUES
                                                      ('abc123', (SELECT id FROM courses WHERE title = 'Introduction to Computer Science')),
                                                      ('abc123', (SELECT id FROM courses WHERE title = 'Advanced Mathematics')),
                                                      ('abc123', (SELECT id FROM courses WHERE title = 'Modern Art History')),
                                                      ('abc123', (SELECT id FROM courses WHERE title = 'Environmental Science')),
                                                      ('abc123', (SELECT id FROM courses WHERE title = 'Business Management')),
                                                      ('abc123', 'c3b34a9f-8f59-4818-a684-9cda56f42d02'), -- Clown Foundations
                                                      ('abc123', '98e64e88-b989-49a0-bbfd-76e158bac634'); -- Delete This Course


-- Inserting teachers for courses into the course_teachers junction table
INSERT INTO course_teachers (course_id, teacher_id) VALUES
                                                        ('c3b34a9f-8f59-4818-a684-9cda56f42d02', 'xyz789'), -- Clown Foundations
                                                      --   ('c3b34a9f-8f59-4818-a684-9cda56f42d02', 'def456'); -- Clown Foundations
                                                      --   ('c3b34a9f-8f59-4818-a684-9cda56f42d02', (SELECT net_id FROM users WHERE username = 'Kevin Smith')), -- Clown Foundations
                                                        ('98e64e88-b989-49a0-bbfd-76e158bac634', (SELECT net_id FROM users WHERE username = 'Alice Jackson')), -- Delete This Course
                                                        ((SELECT id FROM courses WHERE title = 'Introduction to Computer Science'), (SELECT net_id FROM users WHERE username = 'Kevin Smith')),
                                                        ((SELECT id FROM courses WHERE title = 'Advanced Mathematics'), (SELECT net_id FROM users WHERE username = 'Kevin Smith')),
                                                        ((SELECT id FROM courses WHERE title = 'Modern Art History'), (SELECT net_id FROM users WHERE username = 'Jane Doe')),
                                                        ((SELECT id FROM courses WHERE title = 'Environmental Science'), (SELECT net_id FROM users WHERE username = 'Alice Jackson')),
                                                        ((SELECT id FROM courses WHERE title = 'Business Management'), (SELECT net_id FROM users WHERE username = 'Mike Miller'));

CREATE OR REPLACE FUNCTION sync_user_courses()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO user_courses (user_net_id, course_id)
    VALUES (NEW.teacher_id, NEW.course_id);
    RETURN NEW;
END;
CREATE TRIGGER course_teachers_after_insert_trigger
AFTER INSERT ON course_teachers
FOR EACH ROW
EXECUTE FUNCTION sync_user_courses();
-- Further junction table entries to link data as per new structure need to be added here, for example linking courses to users, messages to courses, etc.

SELECT * FROM courses;
SELECT * FROM assignments;
SELECT * FROM messages;
SELECT * FROM projects;
SELECT * FROM submissions;
SELECT * FROM users;
SELECT * FROM tokens;
SELECT * FROM user_courses;
SELECT * FROM course_teachers;
SELECT * FROM course_id;
SELECT * FROM course_roster;
-- Initial setup
CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users Table
CREATE TABLE IF NOT EXISTS users (
                                     net_id VARCHAR PRIMARY KEY,
                                     created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                     updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                     deleted_at TIMESTAMP WITHOUT TIME ZONE,
                                     full_name VARCHAR,
                                     profile_picture_id UUID,
                                     bio TEXT,
                                     username VARCHAR NOT NULL,
                                     password VARCHAR NOT NULL,
                                     email VARCHAR NOT NULL,
                                     membership VARCHAR NOT NULL
);

-- Media Table
CREATE TABLE IF NOT EXISTS media (
                                     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                     type VARCHAR NOT NULL,
                                     url VARCHAR NOT NULL,
                                     created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                     updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                     user_net_id VARCHAR REFERENCES users(net_id)
);

-- Projects Table
CREATE TABLE IF NOT EXISTS projects (
                                        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                        name VARCHAR NOT NULL,
                                        description TEXT,
                                        created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                        updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                        user_net_id VARCHAR REFERENCES users(net_id)
);

-- Courses Table
CREATE TABLE IF NOT EXISTS courses (
                                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                       title VARCHAR NOT NULL,
                                       description TEXT,
                                       created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                       updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
                                       user_net_id VARCHAR REFERENCES users(net_id)
);

-- Foreign key for profile picture which relates to the Media table
ALTER TABLE users ADD CONSTRAINT fk_profile_picture
    FOREIGN KEY (profile_picture_id)
        REFERENCES media (id)
        ON DELETE SET NULL;
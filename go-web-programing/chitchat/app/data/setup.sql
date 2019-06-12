-- Create a new database called 'chitchat'
-- Connect to the 'mysql' database to run this snippet
USE mysql 
GO
-- Create the new database if it does not exist already
-- IF NOT EXISTS (
--     SELECT name
--         FROM sys.databases
--         WHERE name = chitchat
-- )
CREATE DATABASE chitchat CHARSET=utf8 COLLATE=utf8_general_ci;
GO
 
-- Create a new table called 'users' in schema 'chitchat'
-- Drop the table if it already exists
-- IF OBJECT_ID('chitchat.users', 'U') IS NOT NULL
-- DROP TABLE chitchat.users
GO
-- Create the table in the specified schema
CREATE TABLE chitchat.users
(
    id INT NOT NULL PRIMARY KEY, -- primary key column
    uuid VARCHAR(64) NOT NULL UNIQUE,
    name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
    -- specify more columns here
);
GO

-- Create a new table called 'session' in schema 'chitchat'
-- Drop the table if it already exists
-- IF OBJECT_ID('chitchat.session', 'U') IS NOT NULL
-- DROP TABLE chitchat.session
GO
-- Create the table in the specified schema
CREATE TABLE chitchat.session
(
    id INT NOT NULL PRIMARY KEY, -- primary key column
    uuid VARCHAR(64) NOT NULL UNIQUE,
    email VARCHAR(255),
    user_id INT REFERENCES users(id),
    created_at TIMESTAMP NOT NULL
    -- specify more columns here
);
GO

-- Create a new table called 'threads' in schema 'chitchat'
-- Drop the table if it already exists
-- IF OBJECT_ID('chitchat.threads', 'U') IS NOT NULL
-- DROP TABLE chitchat.threads
GO
-- Create the table in the specified schema
CREATE TABLE chitchat.threads
(
    id INT NOT NULL PRIMARY KEY, -- primary key column
    uuid VARCHAR(64) NOT NULL UNIQUE,
    topic TEXT,  
    user_id INT REFERENCES users(id),
    created_at TIMESTAMP NOT NULL
    -- specify more columns here
);
GO

-- Create a new table called 'posts' in schema 'chitchat'
-- Drop the table if it already exists
-- IF OBJECT_ID('chitchat.posts', 'U') IS NOT NULL
-- DROP TABLE chitchat.posts
GO
-- Create the table in the specified schema
CREATE TABLE chitchat.posts
(
    id INT NOT NULL PRIMARY KEY, -- primary key column
    uuid VARCHAR(64) NOT NULL UNIQUE,
    body TEXT,
    user_id INT REFERENCES users(id),
    thread_id INT REFERENCES threads(id),
    created_at TIMESTAMP NOT NULL
    -- specify more columns here
);
GO
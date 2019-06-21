
-- Create a new table called 'posts' in schema 'SchemaName'
-- Drop the table if it already exists

DROP TABLE IF EXISTS test.posts ;
GO
-- Create the table in the specified schema
CREATE TABLE test.posts
(
    postsId INT NOT NULL PRIMARY KEY, -- primary key column
    content TEXT,
    author VARCHAR(255)
    -- specify more columns here
);
GO
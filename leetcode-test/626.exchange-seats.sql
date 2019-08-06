--
-- @lc app=leetcode id=626 lang=mysql
--
-- [626] Exchange Seats
--
-- https://leetcode.com/problems/exchange-seats/description/
--
-- database
-- Medium (54.66%)
-- Likes:    186
-- Dislikes: 178
-- Total Accepted:    27.9K
-- Total Submissions: 50.4K
-- Testcase Example:  '{"headers": {"seat": ["id","student"]}, "rows": {"seat": ' +
  -- '[[1,"Abbot"],[2,"Doris"],[3,"Emerson"],[4,"Green"],[5,"Jeames"]]}}'
--
-- Mary is a teacher in a middle school and she has a table seat storing
-- students' names and their corresponding seat ids.
-- The column id is continuous increment.
-- 
-- 
-- Mary wants to change seats for the adjacent students.
-- 
-- 
-- Can you write a SQL query to output the result for Mary?
-- 
-- 
-- 
-- 
-- +---------+---------+
-- |    id   | student |
-- +---------+---------+
-- |    1    | Abbot   |
-- |    2    | Doris   |
-- |    3    | Emerson |
-- |    4    | Green   |
-- |    5    | Jeames  |
-- +---------+---------+
-- 
-- For the sample input, the output is:
-- 
-- 
-- 
-- 
-- +---------+---------+
-- |    id   | student |
-- +---------+---------+
-- |    1    | Doris   |
-- |    2    | Abbot   |
-- |    3    | Green   |
-- |    4    | Emerson |
-- |    5    | Jeames  |
-- +---------+---------+
-- 
-- 
-- Note:
-- If the number of students is odd, there is no need to change the last one's
-- seat.
-- 
--
-- # Write your MySQL query statement below
SELECT
    s1.id - 1 AS id,
    s1.student
FROM
    seat s1
WHERE
    s1.id MOD 2 = 0 UNION
SELECT
    s2.id + 1 AS id,
    s2.student
FROM
    seat s2
WHERE
    s2.id MOD 2 = 1
    AND s2.id != (SELECT max(s3.id) FROM seat s3) UNION
SELECT
    s4.id AS id,
    s4.student
FROM
    seat s4
WHERE
    s4.id MOD 2 = 1
    AND s4.id = (SELECT max(s3.id) FROM seat s3)
ORDER BY id;

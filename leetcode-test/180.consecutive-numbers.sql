--
-- @lc app=leetcode id=180 lang=mysql
--
-- [180] Consecutive Numbers
--
-- https://leetcode.com/problems/consecutive-numbers/description/
--
-- database
-- Medium (33.76%)
-- Likes:    263
-- Dislikes: 67
-- Total Accepted:    59.3K
-- Total Submissions: 171.6K
-- Testcase Example:  '{"headers": {"Logs": ["Id", "Num"]}, "rows": {"Logs": ' +
  -- '[[1, 1], [2, 1], [3, 1], [4, 2], [5, 1], [6, 2], [7, ' +
  -- '2]]}}'
--
-- Write a SQL query to find all numbers that appear at least three times
-- consecutively.
-- 
-- 
-- +----+-----+
-- | Id | Num |
-- +----+-----+
-- | 1  |  1  |
-- | 2  |  1  |
-- | 3  |  1  |
-- | 4  |  2  |
-- | 5  |  1  |
-- | 6  |  2  |
-- | 7  |  2  |
-- +----+-----+
-- 
-- 
-- For example, given the above Logs table, 1 is the only number that appears
-- consecutively for at least three times.
-- 
-- 
-- +-----------------+
-- | ConsecutiveNums |
-- +-----------------+
-- | 1               |
-- +-----------------+
-- 
-- 
--
-- # Write your MySQL query statement below
SELECT
    DISTINCT L1.Num ConsecutiveNums
FROM
    Logs L1,
    Logs L2,
    Logs L3
WHERE
    L1.Id = L2.Id -1
    AND L2.Id = L3.Id -1
    AND L1.Num = L2.Num
    AND L2.Num = L3.Num;


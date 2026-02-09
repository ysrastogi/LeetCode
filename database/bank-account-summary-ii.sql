# Write your MySQL query statement below
SELECT u.name AS name, SUM(t.amount) AS balance FROM Users u JOIN Transactions t on u.account = t.account
GROUP BY(u.name)
HAVING balance >10000;
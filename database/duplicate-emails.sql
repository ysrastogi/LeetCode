# Write your MySQL query statement below
SELECT p.email
From Person AS p
GROUP BY p.email
HAVING COUNT(p.email)>1
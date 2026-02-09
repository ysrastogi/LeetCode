# Write your MySQL query statement below
SELECT c.name AS Customers
From Customers as c
WHERE id NOT IN(
    SELECT customerId
    FROM Orders
)

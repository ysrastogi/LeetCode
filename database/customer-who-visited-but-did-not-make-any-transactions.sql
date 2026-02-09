SELECT
COUNT(v.visit_id) as count_no_trans,
v.customer_id as customer_id
FROM Visits as v
LEFT JOIN Transactions as t ON v.visit_id=t.visit_id
WHERE t.transaction_id IS NUll
GROUP BY v.customer_id
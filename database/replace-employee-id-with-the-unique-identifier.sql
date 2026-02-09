SELECT
e.name as name,
eu.unique_id as unique_id
FROM Employees as e
LEFT JOIN EmployeeUNI as eu ON e.id=eu.id
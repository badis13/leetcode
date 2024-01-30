--https://leetcode.com/problems/customers-who-never-order/description/

SELECT t1.name AS Customers
FROM Customers AS t1 LEFT JOIN Orders AS t2 ON t1.id = t2.CustomerId
WHERE t2.CustomerId IS NULL
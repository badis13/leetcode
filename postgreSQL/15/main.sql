--https://leetcode.com/problems/sales-person/submissions/1172122138/
SELECT t3.name
FROM Orders AS t1 JOIN Company AS t2 ON t2.com_id = t1.com_id
                    AND t2.name = 'RED'
                  RIGHT JOIN SalesPerson AS t3 ON t3.sales_id = t1.sales_id
WHERE t1.sales_id IS NULL
--https://leetcode.com/problems/find-customer-referee/
-- не решено
SELECT *
FROM Customer
WHERE referee_id <> 2
AND referee_id is NULL
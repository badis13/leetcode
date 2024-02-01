-- https://leetcode.com/problems/employees-earning-more-than-their-managers/description/


SELECT name AS Employee
FROM Employee AS t1
WHERE salary > (
    SELECT salary
    FROM Employee
    WHERE id = t1.managerId
)

SELECT t1.name as Employee
FROM Employee AS t1, Employee AS t2
WHERE t1.managerId = t2.id and t1.salary > t2.salary
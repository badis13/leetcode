--https://leetcode.com/problems/combine-two-tables/submissions/1174240246/
SELECT t1.firstName, t1.lastName, t2.city, t2.state
FROM Person AS t1 LEFT JOIN Address AS t2
ON t2.personId = t1.personId
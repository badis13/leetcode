--https://leetcode.com/problems/combine-two-tables/description/

SELECT t1.firstName, t1.lastName, t2.city, t2.state
FROM Person AS t1 LEFT JOIN Adress AS t2 ON t1.personId = t2.personId
-- https://leetcode.com/problems/big-countries/description/

SELECT name, population, area
FROM World
WHERE area > 2999999 OR population > 24999999
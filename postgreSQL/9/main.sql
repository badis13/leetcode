--https://leetcode.com/problems/delete-duplicate-emails/description/

DELETE
FROM Person AS a
USING Person  AS b
WHERE a.id > b.id
AND a.email = b.email
 
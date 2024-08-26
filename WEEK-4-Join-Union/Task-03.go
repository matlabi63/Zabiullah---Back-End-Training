-- SELECT

SELECT *
FROM Users
WHERE gender = 'Male' OR gender = 'M';


SELECT *
FROM Products
WHERE product_id = 3;

SELECT *
FROM Users
WHERE created_at >= DATE_SUB(CURRENT_DATE, INTERVAL 7 DAY)
AND username LIKE '%a%';

SELECT COUNT(*)
FROM Users
WHERE gender = 'Female';

SELECT *
FROM Users
ORDER BY username ASC;

SELECT *
FROM Products
LIMIT 5;

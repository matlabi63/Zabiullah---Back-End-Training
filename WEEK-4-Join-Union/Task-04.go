-- UPDATE

UPDATE Products
SET name = 'product dummy'
WHERE product_id = 1;

UPDATE Transaction_Details
SET qty = 3
WHERE product_id = 1;

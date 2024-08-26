SELECT * FROM Transactions WHERE user_id = 1
UNION ALL
SELECT * FROM Transactions WHERE user_id = 2;

SELECT SUM(total_price) AS total_price_user_1
FROM Transactions
WHERE user_id = 1;

SELECT COUNT(*) AS total_transactions
FROM TransactionDetails td
JOIN Products p ON td.product_id = p.product_id
WHERE p.product_type_id = 2;

SELECT p.*, pt.product_type_name
FROM Products p
JOIN ProductTypes pt ON p.product_type_id = pt.product_type_id;

SELECT t.*, p.product_name, u.name AS user_name
FROM Transactions t
JOIN TransactionDetails td ON t.transaction_id = td.transaction_id
JOIN Products p ON td.product_id = p.product_id
JOIN Users u ON t.user_id = u.user_id;

CREATE TRIGGER delete_transaction_details
AFTER DELETE ON Transactions
FOR EACH ROW
BEGIN
    DELETE FROM TransactionDetails WHERE transaction_id = OLD.transaction_id;
END;

CREATE TRIGGER update_total_qty
AFTER DELETE ON TransactionDetails
FOR EACH ROW
BEGIN
    UPDATE Transactions
    SET total_qty = total_qty - OLD.qty
    WHERE transaction_id = OLD.transaction_id;
END;

SELECT *
FROM Products
WHERE product_id NOT IN (SELECT DISTINCT product_id FROM TransactionDetails);

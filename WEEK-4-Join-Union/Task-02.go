-- Insert 5 operators
INSERT INTO Operators (operator_name) VALUES
('Operator A'),
('Operator B'),
('Operator C'),
('Operator D'),
('Operator E');

-- Insert 3 product types
INSERT INTO ProductTypes (product_type_name) VALUES
('Electronics'),
('Clothing'),
('Books');

-- Insert 2 products with product type id = 1 and operator id = 3
INSERT INTO Products (product_name, description, product_type_id, operator_id) VALUES
('Smartphone', 'Latest model smartphone', 1, 3),
('Laptop', 'High performance laptop', 1, 3);

-- Insert 3 products with product type id = 2 and operator id = 1
INSERT INTO Products (product_name, description, product_type_id, operator_id) VALUES
('T-Shirt', 'Cotton T-Shirt', 2, 1),
('Jeans', 'Denim jeans', 2, 1),
('Jacket', 'Winter jacket', 2, 1);

-- Insert 3 products with product type id = 3 and operator id = 4
INSERT INTO Products (product_name, description, product_type_id, operator_id) VALUES
('Fiction Book', 'Bestselling fiction book', 3, 4),
('Non-fiction Book', 'Popular non-fiction book', 3, 4),
('Science Book', 'Educational science book', 3, 4);

-- Insert 3 payment methods
INSERT INTO PaymentMethods (payment_method_name) VALUES
('Credit Card'),
('PayPal'),
('Bank Transfer');

-- Insert 5 users
INSERT INTO Users (username, email) VALUES
('Alice', 'alice@example.com'),
('Bob', 'bob@example.com'),
('Charlie', 'charlie@example.com'),
('Diana', 'diana@example.com'),
('Eve', 'eve@example.com');

-- Insert 3 transactions for each user
INSERT INTO Transactions (user_id, transaction_date, payment_method_id) VALUES
(1, '2024-08-05', 1),
(1, '2024-08-06', 2),
(1, '2024-08-07', 3),
(2, '2024-08-05', 2),
(2, '2024-08-06', 3),
(2, '2024-08-07', 1),
(3, '2024-08-05', 3),
(3, '2024-08-06', 1),
(3, '2024-08-07', 2),
(4, '2024-08-05', 1),
(4, '2024-08-06', 2),
(4, '2024-08-07', 3),
(5, '2024-08-05', 2),
(5, '2024-08-06', 3),
(5, '2024-08-07', 1);

-- Insert 3 products for each transaction
INSERT INTO TransactionProducts (transaction_id, product_id) VALUES
(1, 1), (1, 2), (1, 3),
(2, 2), (2, 3), (2, 4),
(3, 3), (3, 4), (3, 5),
(4, 4), (4, 5), (4, 6),
(5, 5), (5, 6), (5, 7),
(6, 6), (6, 7), (6, 8),
(7, 7), (7, 8), (7, 9),
(8, 8), (8, 9), (8, 10),
(9, 9), (9, 10), (9, 1),
(10, 10), (10, 1), (10, 2),
(11, 1), (11, 2), (11, 3),
(12, 2), (12, 3), (12, 4),
(13, 3), (13, 4), (13, 5),
(14, 4), (14, 5), (14, 6),
(15, 5), (15, 6), (15, 7);

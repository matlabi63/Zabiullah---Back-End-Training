
CREATE TABLE product_types (
		id INT(1),
    name VARCHAR(50),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE operators (
	name VARCHAR(50),
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE product (
		id INT(11)
    product_type_id INT(11),
    operator_id INT(11),
    code VARCHAR(50),
    name VARCHAR(100),
		status SMALLINT,
		created_at TIMESTAMP,
		updated_at TIMESTAMP
);

CREATE TABLE prodduct_description (
    id INT(11),
    description TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE transaction_details (
    transaction_id INT(11),
    product_id INT(11),
    status VARCHAR(10),
    qty INT(11),
    price NUMERIC(25.2),
		created_at TIMESTAMP,
		updated_at TIMESTAMP
);

CREATE TABLE transactions (
    id INT(11),
    user_id INT,(11)
		payment_method_id INT(11),
    status VARCHAR(10),
    total_price NUMERIC(25.2),
    created_at TIMESTAMP,
		updated_at TIMESTAMP
);

CREATE TABLE payment_methods (
	id INT(11),
	name VARCHAR(255),
	status SMALLINT,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE users (
	id INT(11),
	status SMALLINT,
	dob DATE,
	gender CHAR(1)
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);
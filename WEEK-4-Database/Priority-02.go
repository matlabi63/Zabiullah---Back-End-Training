
CREATE TABLE restaurant_types (
    id INT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);

CREATE TABLE restaurants (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    address TEXT,
    restaurant_types_id INT
);

 CREATE TABLE menus (
    id INT PRIMARY KEY,
    menu_type_id INT,
    restaurant_id INT,
    name VARCHAR(255),
		description TEXT,
		price INT
);

CREATE TABLE menu_types (
    id INT PRIMARY KEY,
   	name VARCHAR(255)
);

CREATE TABLE users (
    id INT PRIMARY KEY,
    username VARCHAR(100),
    email VARCHAR(255),
		password VARCHAR(255)
);

CREATE TABLE user_reviews (
    id INT PRIMARY KEY,
    restaurant_id INT,
		rating FLOAT,
		description TEXT,
		user_id INT
);

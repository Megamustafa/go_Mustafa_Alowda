CREATE TABLE restaurant_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE restaurants (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    restaurant_type_id INT,
    FOREIGN KEY (restaurant_type_id) REFERENCES restaurant_types(id)
);

CREATE TABLE menu_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE menus (
    id INT AUTO_INCREMENT PRIMARY KEY,
    menu_type_id INT,
    restaurant_id INT,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price INT,
    FOREIGN KEY (menu_type_id) REFERENCES menu_types(id),
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id)
);

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE user_reviews (
    id INT AUTO_INCREMENT PRIMARY KEY,
    restaurant_id INT,
    user_id INT,
    rating FLOAT,
    description TEXT,
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

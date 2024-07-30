-- Insert operators
INSERT INTO operators (id, name, created_at, updated_at) VALUES 
(1, 'Operator 1', NOW(), NOW()),
(2, 'Operator 2', NOW(), NOW()),
(3, 'Operator 3', NOW(), NOW()),
(4, 'Operator 4', NOW(), NOW()),
(5, 'Operator 5', NOW(), NOW());

-- Insert product types
INSERT INTO product_types (id, name, created_at, updated_at) VALUES
(1, 'Type 1', NOW(), NOW()),
(2, 'Type 2', NOW(), NOW()),
(3, 'Type 3', NOW(), NOW());

-- Insert products
INSERT INTO products (id, product_type_id, operator_id, code, name, status, created_at, updated_at) VALUES
(1, 1, 3, 'P001', 'Product 1', 1, NOW(), NOW()),
(2, 1, 3, 'P002', 'Product 2', 1, NOW(), NOW()),
(3, 2, 1, 'P003', 'Product 3', 1, NOW(), NOW()),
(4, 2, 1, 'P004', 'Product 4', 1, NOW(), NOW()),
(5, 2, 1, 'P005', 'Product 5', 1, NOW(), NOW()),
(6, 3, 4, 'P006', 'Product 6', 1, NOW(), NOW()),
(7, 3, 4, 'P007', 'Product 7', 1, NOW(), NOW()),
(8, 3, 4, 'P008', 'Product 8', 1, NOW(), NOW());

-- Insert product descriptions
INSERT INTO product_descriptions (id, description, created_at, updated_at) VALUES
(1, 'Description for Product 1', NOW(), NOW()),
(2, 'Description for Product 2', NOW(), NOW()),
(3, 'Description for Product 3', NOW(), NOW()),
(4, 'Description for Product 4', NOW(), NOW()),
(5, 'Description for Product 5', NOW(), NOW()),
(6, 'Description for Product 6', NOW(), NOW()),
(7, 'Description for Product 7', NOW(), NOW()),
(8, 'Description for Product 8', NOW(), NOW());

-- Insert payment methods
INSERT INTO payment_methods (id, name, status, created_at, updated_at) VALUES
(1, 'Credit Card', 1, NOW(), NOW()),
(2, 'PayPal', 1, NOW(), NOW()),
(3, 'Bank Transfer', 1, NOW(), NOW());

-- Insert users
INSERT INTO users (id, status, dob, gender, created_at, updated_at) VALUES
(1, 1, '1990-01-01', 'M', NOW(), NOW()),
(2, 1, '1991-02-02', 'F', NOW(), NOW()),
(3, 1, '1992-03-03', 'M', NOW(), NOW()),
(4, 1, '1993-04-04', 'F', NOW(), NOW()),
(5, 1, '1994-05-05', 'M', NOW(), NOW());

-- Insert transactions for each user
-- Using multiple insert statements for simplicity
INSERT INTO transactions (id, user_id, payment_method_id, status, total_qty, total_price, created_at, updated_at) VALUES
(1, 1, 1, 'completed', 3, 300.00, NOW(), NOW()),
(2, 1, 2, 'completed', 3, 300.00, NOW(), NOW()),
(3, 1, 3, 'completed', 3, 300.00, NOW(), NOW()),
(4, 2, 1, 'completed', 3, 300.00, NOW(), NOW()),
(5, 2, 2, 'completed', 3, 300.00, NOW(), NOW()),
(6, 2, 3, 'completed', 3, 300.00, NOW(), NOW()),
(7, 3, 1, 'completed', 3, 300.00, NOW(), NOW()),
(8, 3, 2, 'completed', 3, 300.00, NOW(), NOW()),
(9, 3, 3, 'completed', 3, 300.00, NOW(), NOW()),
(10, 4, 1, 'completed', 3, 300.00, NOW(), NOW()),
(11, 4, 2, 'completed', 3, 300.00, NOW(), NOW()),
(12, 4, 3, 'completed', 3, 300.00, NOW(), NOW()),
(13, 5, 1, 'completed', 3, 300.00, NOW(), NOW()),
(14, 5, 2, 'completed', 3, 300.00, NOW(), NOW()),
(15, 5, 3, 'completed', 3, 300.00, NOW(), NOW());

-- Insert products into each transaction
-- Using multiple insert statements for simplicity
INSERT INTO transaction_details (transaction_id, product_id, qty, price, created_at, updated_at) VALUES
(1, 1, 1, 100.00, NOW(), NOW()),
(1, 2, 1, 100.00, NOW(), NOW()),
(1, 3, 1, 100.00, NOW(), NOW()),
(2, 1, 1, 100.00, NOW(), NOW()),
(2, 2, 1, 100.00, NOW(), NOW()),
(2, 3, 1, 100.00, NOW(), NOW()),
(3, 1, 1, 100.00, NOW(), NOW()),
(3, 2, 1, 100.00, NOW(), NOW()),
(3, 3, 1, 100.00, NOW(), NOW()),
(4, 1, 1, 100.00, NOW(), NOW()),
(4, 2, 1, 100.00, NOW(), NOW()),
(4, 3, 1, 100.00, NOW(), NOW()),
(5, 1, 1, 100.00, NOW(), NOW()),
(5, 2, 1, 100.00, NOW(), NOW()),
(5, 3, 1, 100.00, NOW(), NOW()),
(6, 1, 1, 100.00, NOW(), NOW()),
(6, 2, 1, 100.00, NOW(), NOW()),
(6, 3, 1, 100.00, NOW(), NOW()),
(7, 1, 1, 100.00, NOW(), NOW()),
(7, 2, 1, 100.00, NOW(), NOW()),
(7, 3, 1, 100.00, NOW(), NOW()),
(8, 1, 1, 100.00, NOW(), NOW()),
(8, 2, 1, 100.00, NOW(), NOW()),
(8, 3, 1, 100.00, NOW(), NOW()),
(9, 1, 1, 100.00, NOW(), NOW()),
(9, 2, 1, 100.00, NOW(), NOW()),
(9, 3, 1, 100.00, NOW(), NOW()),
(10, 1, 1, 100.00, NOW(), NOW()),
(10, 2, 1, 100.00, NOW(), NOW()),
(10, 3, 1, 100.00, NOW(), NOW()),
(11, 1, 1, 100.00, NOW(), NOW()),
(11, 2, 1, 100.00, NOW(), NOW()),
(11, 3, 1, 100.00, NOW(), NOW()),
(12, 1, 1, 100.00, NOW(), NOW()),
(12, 2, 1, 100.00, NOW(), NOW()),
(12, 3, 1, 100.00, NOW(), NOW()),
(13, 1, 1, 100.00, NOW(), NOW()),
(13, 2, 1, 100.00, NOW(), NOW()),
(13, 3, 1, 100.00, NOW(), NOW()),
(14, 1, 1, 100.00, NOW(), NOW()),
(14, 2, 1, 100.00, NOW(), NOW()),
(14, 3, 1, 100.00, NOW(), NOW()),
(15, 1, 1, 100.00, NOW(), NOW()),
(15, 2, 1, 100.00, NOW(), NOW()),
(15, 3, 1, 100.00, NOW(), NOW());

-- Show users/customers name with gender Male/M
SELECT name FROM users WHERE gender = 'M';

-- Show products with id = 3
SELECT * FROM products WHERE id = 3;

-- Show customers data created_at within the range of past 7 days and having the letter 'a'
SELECT * FROM users WHERE created_at BETWEEN NOW() - INTERVAL 7 DAY AND NOW() AND name LIKE '%a%';

-- Count the amount of users/customers with the gender status of Female
SELECT COUNT(*) FROM users WHERE gender = 'F';

-- Show customers data with each of their names in alphabetical order
SELECT * FROM users ORDER BY name ASC;

-- Show 5 data at product data
SELECT * FROM products LIMIT 5;

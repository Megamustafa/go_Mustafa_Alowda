-- Change data product id 1 with the name of 'product dummy'
UPDATE products SET name = 'product dummy' WHERE id = 1;

-- Update qty = 3 on transaction detail with product id = 1
UPDATE transaction_details SET qty = 3 WHERE product_id = 1;

-- Combine transaction data from user id 1 and user id 2
SELECT * FROM transactions WHERE user_id IN (1, 2);

-- Show the total price from user id 1
SELECT SUM(total_price) FROM transactions WHERE user_id = 1;

-- Show the total transaction within product type 2
SELECT COUNT(*) FROM transactions 
JOIN transaction_details ON transactions.id = transaction_details.transaction_id
JOIN products ON transaction_details.product_id = products.id
WHERE products.product_type_id = 2;

-- Show all product table fields and name table product type fields that are correlated with each other
SELECT products.*, product_types.name AS product_type_name
FROM products
JOIN product_types ON products.product_type_id = product_types.id;

-- Show all transaction table fields, name table product field and name table user field
SELECT transactions.*, products.name AS product_name, users.name AS user_name
FROM transactions
JOIN transaction_details ON transactions.id = transaction_details.transaction_id
JOIN products ON transaction_details.product_id = products.id
JOIN users ON transactions.user_id = users.id;

-- Function to delete transaction details after deletion of transaction data
CREATE OR REPLACE FUNCTION delete_transaction_details()
RETURNS TRIGGER AS $$
BEGIN
    DELETE FROM transaction_details WHERE transaction_id = OLD.id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_delete_transaction
AFTER DELETE ON transactions
FOR EACH ROW
EXECUTE FUNCTION delete_transaction_details();

-- Function to update total_qty after deletion of transaction details
CREATE OR REPLACE FUNCTION update_total_qty()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE transactions
    SET total_qty = total_qty - OLD.qty
    WHERE id = OLD.transaction_id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_delete_transaction_detail
AFTER DELETE ON transaction_details
FOR EACH ROW
EXECUTE FUNCTION update_total_qty();

-- Show products data that has never been on the transaction_details table with sub-query
SELECT * FROM products
WHERE id NOT IN (SELECT DISTINCT product_id FROM transaction_details);

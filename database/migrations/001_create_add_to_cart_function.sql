CREATE OR REPLACE FUNCTION add_to_cart(
    in_cart_id INT, 
    in_product_id VARCHAR(255), 
    in_quantity INT
) RETURNS VOID AS $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM cart_items 
        WHERE cart_id = in_cart_id AND product_id = in_product_id
    ) THEN
        UPDATE cart_items
        SET quantity = quantity + in_quantity
        WHERE cart_id = in_cart_id AND product_id = in_product_id;
    ELSE 
        INSERT INTO cart_items (cart_id, product_id, quantity, added_at)
        VALUES (in_cart_id, in_product_id, in_quantity, CURRENT_TIMESTAMP);
    END IF;
END;
$$ LANGUAGE plpgsql;

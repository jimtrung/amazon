CREATE OR REPLACE FUNCTION update_cart(
    in_cart_id INT,
    in_product_id VARCHAR(255),
    new_quantity INT
) RETURNS VOID AS $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM cart_items WHERE cart_id = in_cart_id AND product_id = in_product_id
    ) THEN
        RAISE EXCEPTION 'Cannot find item in cart';
    END IF;

    IF new_quantity <= 0 THEN
        DELETE FROM cart_items WHERE cart_id = in_cart_id AND product_id = in_product_id;
    ELSE
        UPDATE cart_items
        SET quantity = new_quantity
        WHERE cart_id = in_cart_id AND product_id = in_product_id;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION delete_from_cart(
    in_cart_id INT,
    in_product_id VARCHAR(255)
) RETURNS VOID AS $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM cart_items WHERE cart_id = in_cart_id AND product_id = in_product_id
    ) THEN
        DELETE FROM cart_items WHERE cart_id = in_cart_id AND product_id = in_product_id;
    ELSE
        RAISE EXCEPTION 'Cannot find item in cart';
    END IF;
END;
$$ LANGUAGE plpgsql;

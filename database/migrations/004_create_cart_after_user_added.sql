CREATE OR REPLACE FUNCTION create_cart_after_user_insert()
RETURNS TRIGGER AS $$
BEGIN
    -- Insert a new cart for the new user
    INSERT INTO carts (user_id, created_at, updated_at)
    VALUES (NEW.user_id, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_create_cart
AFTER INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION create_cart_after_user_insert();

DELETE FROM users;
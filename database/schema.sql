-- Products Table
CREATE TABLE IF NOT EXISTS products (
  product_id VARCHAR(255) PRIMARY KEY,
  image VARCHAR(255) DEFAULT 'no_path',
  name VARCHAR(255) NOT NULL DEFAULT 'no_name',
  rating_stars FLOAT CHECK (rating_stars BETWEEN 0 AND 5), 
  rating_count INT CHECK (rating_count >= 0), 
  price_cents INT CHECK (price_cents >= 0), 
  keywords TEXT[] DEFAULT ARRAY[]::TEXT[],
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Users Table
CREATE TABLE IF NOT EXISTS users (
  user_id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE, 
  password VARCHAR(255) NOT NULL CHECK (LENGTH(password) >= 6), 
  email VARCHAR(255) NOT NULL UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'),
  phone VARCHAR(255) UNIQUE CHECK (phone ~* '^[0-9]+$'), 
  country VARCHAR(255),
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Carts Table
CREATE TABLE IF NOT EXISTS carts (
  cart_id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_cart_user FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Cart Items Table
CREATE TABLE IF NOT EXISTS cart_items (
  cart_item_id SERIAL PRIMARY KEY,
  cart_id INT NOT NULL,
  product_id VARCHAR(255) NOT NULL,
  quantity INT CHECK (quantity > 0),
  added_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_cart FOREIGN KEY (cart_id) REFERENCES carts(cart_id) ON DELETE CASCADE,
  CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE
);

-- Drop
DROP TABLE cart_items;
DROP TABLE carts;
DROP TABLE users;
DROP TABLE products;

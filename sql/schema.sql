CREATE TABLE products (
	id   		 UUID PRIMARY KEY,
	image		 VARCHAR(255),
	name 		 VARCHAR(255) NOT NULL,
	rating_stars FLOAT,
	rating_count INT,
	price_cents  INT,
	keywords     TEXT[]
);

CREATE TABLE cart (
	product_id UUID,
	quantity   INT,
	FOREIGN KEY (product_id) REFERENCES products(id)
);

DROP TABLE products;
DROP TABLE cart;
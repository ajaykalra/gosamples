-- Create a table named products
CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock_quantity INTEGER DEFAULT 0
);

-- Insert sample data into the products table
INSERT INTO products (product_name, price, stock_quantity) VALUES
    ('Laptop', 1200.00, 50),
    ('Mouse', 25.50, 200),
    ('Keyboard', 75.00, 100),
    ('Monitor', 300.00, 30);

-- Query the data to verify
SELECT * FROM products;
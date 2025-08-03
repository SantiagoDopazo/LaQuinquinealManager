CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    order_number VARCHAR(255) UNIQUE NOT NULL,
    client_name VARCHAR(255) NOT NULL,
    client_cuit VARCHAR(255),
    client_contact VARCHAR(255),
    client_address TEXT,
    product_type VARCHAR(255),
    material VARCHAR(255),
    thickness_mm DECIMAL(10,2),
    width_mm DECIMAL(10,2),
    length_mm DECIMAL(10,2),
    weight_kg DECIMAL(10,2),
    quantity INTEGER,
    order_date TIMESTAMP,
    delivery_date TIMESTAMP,
    status VARCHAR(100),
    notes TEXT,
    unit_price DECIMAL(10,2),
    total_price DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_orders_order_number ON orders(order_number);
CREATE INDEX idx_orders_status ON orders(status);
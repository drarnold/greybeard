-- Migration: Add non-nullable status column and index to 80M-row orders table
ALTER TABLE orders ADD COLUMN fulfillment_status VARCHAR(32) NOT NULL;
CREATE INDEX idx_orders_fulfillment_status ON orders(fulfillment_status);

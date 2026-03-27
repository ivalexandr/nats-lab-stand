CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    external_id TEXT NOT NULL UNIQUE,
    status TEXT NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    sent_at TIMESTAMP NULL
);

INSERT INTO orders (external_id, status, amount)
VALUES
    ('ord-1001', 'new', 1250.50),
    ('ord-1002', 'new', 300.00),
    ('ord-1003', 'new', 999.99);

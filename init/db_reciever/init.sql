CREATE TABLE IF NOT EXISTS orders_mirror (
    id BIGSERIAL PRIMARY KEY,
    source_id BIGINT NOT NULL,
    external_id TEXT NOT NULL UNIQUE,
    status TEXT NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    received_at TIMESTAMP NOT NULL DEFAULT NOW()
);

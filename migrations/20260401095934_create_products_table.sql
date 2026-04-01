-- +goose Up
CREATE TABLE products (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sku        VARCHAR(100) NOT NULL UNIQUE,
    name       VARCHAR(200) NOT NULL,
    unit       VARCHAR(20)  NOT NULL DEFAULT 'pcs',
    min_stock  INT          NOT NULL DEFAULT 0,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP    NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS products;
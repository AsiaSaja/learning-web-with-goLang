-- +goose Up
CREATE TYPE movement_type AS ENUM ('in', 'out');

CREATE TABLE stock_movements (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    product_id   UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    location_id  UUID NOT NULL REFERENCES locations(id),
    type         movement_type NOT NULL,
    quantity     INT NOT NULL CHECK (quantity > 0),
    notes        TEXT,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_stock_movements_product ON stock_movements(product_id);
CREATE INDEX idx_stock_movements_created ON stock_movements(created_at DESC);

-- +goose Down
DROP TABLE IF EXISTS stock_movements;
DROP TYPE IF EXISTS movement_type;
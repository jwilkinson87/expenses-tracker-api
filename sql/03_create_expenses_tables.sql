CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY,
    label VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS expenses (
    id UUID PRIMARY KEY,
    amount BIGINT NOT NULL,
    category_id UUID REFERENCES categories (id) ON DELETE SET NULL,
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    expense_date DATE NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
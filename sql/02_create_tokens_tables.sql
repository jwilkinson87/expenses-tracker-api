CREATE TABLE IF NOT EXISTS users_auth_tokens (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expiry_time TIMESTAMP,
    token VARCHAR(64) NOT NULL,
    user_id UUID REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS reset_tokens (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expiry_time TIMESTAMP,
    reset_token VARCHAR(64) NOT NULL,
    user_id UUID REFERNECES users (id) ON DELETE CASCADE
);
CREATE TABLE bills (
    id BIGSERIAL PRIMARY KEY,
    chat_id BIGINT NOT NULL,
    total_amount DECIMAL(10, 2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE participants (
    id BIGSERIAL PRIMARY KEY,
    bill_id BIGINT NOT NULL REFERENCES bills(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL,
    username VARCHAR(255),
    amount_owed DECIMAL(10, 2) DEFAULT 0,
    has_paid BOOLEAN DEFAULT FALSE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(bill_id, user_id) -- prevent same user from joining twice
);

CREATE TABLE food_items (
    id BIGSERIAL PRIMARY KEY,
    bill_id BIGINT NOT NULL REFERENCES bills(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

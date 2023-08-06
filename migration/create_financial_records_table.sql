CREATE TABLE IF NOT EXISTS financial_records (
    id SERIAL PRIMARY KEY,
    type VARCHAR(20),
    amount INT,
    user_id INT NOT NULL,
    created_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),
    notes VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
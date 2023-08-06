CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    password VARCHAR(255),
    created_at timestamp(0),
    total_expenses INT,
    total_income INT,
    BALANCE INT,
    UNIQUE(email)
);
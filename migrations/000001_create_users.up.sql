CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    balance FLOAT8 NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()

);
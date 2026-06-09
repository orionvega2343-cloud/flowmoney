CREATE TABLE IF NOT EXISTS transactions(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    amount FLOAT8 NOT NULL,
    type TEXT NOT NULL,
    date TIMESTAMP,
    category_id INT NOT NULL REFERENCES categories(id)
);
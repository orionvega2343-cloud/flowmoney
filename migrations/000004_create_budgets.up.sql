CREATE TABLE IF NOT EXISTS budgets(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    category_id INT NOT NULL REFERENCES categories(id),
    amount FLOAT8 NOT NULL,
    month INT NOT NULL,
    year INT NOT NULL
);
CREATE TABLE IF NOT EXISTS categories(
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    user_id INT NOT NULL REFERENCES users(id)
)
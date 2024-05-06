CREATE TABLE IF NOT EXISTS Posts(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id)
    ON DELETE CASCADE,
    title VARCHAR(255),
    content TEXT,
    likes INTEGER
)

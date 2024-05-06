CREATE TABLE IF NOT EXISTS Comments(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id)
    ON DELETE CASCADE,
    post_id INTEGER REFERENCES Posts(id),
    content TEXT
)

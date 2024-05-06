CREATE TABLE IF NOT EXISTS Users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(64),
    password VARCHAR(64)
)

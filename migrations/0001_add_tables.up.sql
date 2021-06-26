CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  phone_number TEXT NOT NULL UNIQUE,
  password TEXT,
  category_id INT
);

CREATE TABLE IF NOT EXISTS user_category (
  id serial PRIMARY KEY,
  category_name TEXT
);
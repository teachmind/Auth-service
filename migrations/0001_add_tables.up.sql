CREATE TABLE IF NOT EXISTS user_category (
  id serial PRIMARY KEY,
  category_name TEXT
);

CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  category_id INT,
  phone_number TEXT NOT NULL UNIQUE,
  password TEXT,
  CONSTRAINT category_id
      FOREIGN KEY(category_id) 
	      REFERENCES user_category(id)
          ON DELETE CASCADE
);

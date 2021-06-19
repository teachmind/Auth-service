CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  phone_number TEXT NOT NULL UNIQUE,
  password TEXT,
  category_id INT
);

CREATE TABLE IF NOT EXISTS carrier_request (
  id serial PRIMARY KEY,
  parcel_id INT,
  carrier_id INT,
  status INT
);

CREATE TABLE IF NOT EXISTS user_category (
  id serial PRIMARY KEY,
  category_name TEXT
);

CREATE TABLE IF NOT EXISTS carrier_request_status (
  id serial PRIMARY KEY,
  status_value TEXT
);

CREATE TABLE IF NOT EXISTS parcel_status (
  id serial PRIMARY KEY,
  status_value TEXT
);

CREATE TABLE IF NOT EXISTS parcel (
  id serial PRIMARY KEY,
  user_id INT,
  carrier_id INT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  status INT,
  source_address TEXT,
  destination_address TEXT,
  price FLOAT,
  carrier_fee FLOAT,
  company_fee FLOAT
);
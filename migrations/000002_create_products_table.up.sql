CREATE TABLE IF NOT EXISTS public.products (
id serial PRIMARY KEY,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
deleted_at TIMESTAMP NOT NULL,
name VARCHAR(100) NOT NULL,
description VARCHAR(255) NOT NULL,
image_url VARCHAR(100) NOT NULL,
price integer
)
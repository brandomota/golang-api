CREATE TABLE IF NOT EXISTS public.users (
id serial PRIMARY KEY,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
deleted_at TIMESTAMP NOT NULL,
name VARCHAR(100) NOT NULL UNIQUE ,
email VARCHAR(100) NOT NULL UNIQUE ,
age integer ,
user_type integer
)
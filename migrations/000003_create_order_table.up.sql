CREATE TABLE IF NOT EXISTS public.orders (
id serial PRIMARY KEY,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
deleted_at TIMESTAMP NOT NULL,
user_id INTEGER NOT NULL,
status VARCHAR(100) NOT NULL,
FOREIGN KEY (user_id) REFERENCES public.users(id)
)
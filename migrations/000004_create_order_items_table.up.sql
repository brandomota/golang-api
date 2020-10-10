CREATE TABLE IF NOT EXISTS public.orders (
id serial PRIMARY KEY,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
deleted_at TIMESTAMP NOT NULL,
order_id INTEGER NOT NULL,
product_id INTEGER NOT NULL,
quantity INTEGER NOT NULL,
FOREIGN KEY (order_id) REFERENCES public.orders(id),
FOREIGN KEY (product_id) REFERENCES public.products(id)
)
CREATE TABLE IF NOT EXISTS products_lists
(
  id         serial not null,
  user_id    int references users (user_id) on delete cascade not null,
  product_id int references products (id) not null,
  created_at timestamp NOT NULL DEFAULT NOW()
);
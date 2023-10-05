CREATE TABLE IF NOT EXISTS products_lists
(
  id         serial not null,
  user_id    int references users (id) on delete cascade not null,
  product_id int references products (id) on delete cascade not null,
  created_at timestamp NOT NULL DEFAULT NOW(),
  CONSTRAINT uc_user_product UNIQUE (user_id, product_id)
);

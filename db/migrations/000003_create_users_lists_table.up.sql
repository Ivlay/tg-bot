CREATE TABLE IF NOT EXISTS users_lists
(
  id      serial not null unique,
  user_id int references users (user_id) not null,
  product_id int references products (id) not null
);
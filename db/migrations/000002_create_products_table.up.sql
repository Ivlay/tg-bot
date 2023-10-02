CREATE TABLE IF NOT EXISTS products (
  id         serial not null unique,
  title      varchar(255) not null unique,
  price_id   varchar(255) not null,
  price      int not null,
  updated_at timestamp default current_timestamp not null,
  old_price  int not null
);

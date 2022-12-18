CREATE TABLE IF NOT EXISTS products (
  id serial not null unique,
  title varchar(255) not null,
  price_id varchar(255) unique not null,
  price int not null
);

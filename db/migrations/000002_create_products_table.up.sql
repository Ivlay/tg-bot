CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  NEW.old_price = OLD.price;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS products (
  id         serial not null unique,
  title      varchar(255) not null unique,
  price_id   varchar(255) not null,
  price      int not null,
  updated_at timestamp not null default now(),
  old_price  int not null
);

CREATE OR REPLACE TRIGGER set_timestamp
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

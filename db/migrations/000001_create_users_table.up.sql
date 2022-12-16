CREATE TABLE users
(
  id            serial       not null unique,
  firstname     varchar(255) not null,
  username      varchar(255) not null,
  chat_id       int primary key not null unique,
  user_id       int not null unique,
  created_at    timestamp NOT NULL DEFAULT NOW(),
  updated_at    timestamp NOT NULL DEFAULT NOW()
)
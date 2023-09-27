CREATE TABLE IF NOT EXISTS author(
  id serial primary key,
  email varchar not null,
  name varchar not null
);


CREATE TABLE IF NOT EXISTS post(
  id serial primary key,
  title varchar not null,
  content text,
  created_at timestamp default now(),
  updated_at timestamp default now(),
  author int not null,
  foreign key (author) references author(id)
);
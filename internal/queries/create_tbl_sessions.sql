use snippetbox;

create table sessions (
  token char(43) primary key,
  data blob not null,
  expiry timestamp(6) not null
);

create index sessions_expiry_idx on sessions (expiry);

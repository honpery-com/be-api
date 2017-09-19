-- tags
create table tags (
    id serial,
    name varchar(200) not null,
    description varchar(200),
    create_at timestamp not null,
    update_at timestamp not null
);

-- categories
create table categories (
    id serial,
    name varchar(200) not null,
    description varchar(200),
    create_at timestamp not null,
    update_at timestamp not null,
);

-- users
create table users (
    id serial,
    username varchar(20) not null,
    password varchar(100) not null,
    description varchar(200),
    create_at timestamp not null,
    update_at timestamp not null,
);

-- articles
create table articles (
    id serial,
    title varchar(100) not null,
    description varchar(200),
    body text not null,
    category int not null,
    tags int[] not null,
    author int not null,
    status int default 0,
    create_at timestamp not null,
    update_at timestamp not null,
    publish_at timestamp not null,
);
create table users (
    id serial primary key,
    age int not null,
    birthday date not null,
    career char(140) default '',
    email char(100) not null unique,
    start_age int not null,
    end_age int not null,
    gender char(25) not null,
    info char(500) default '',
    first_name char(50) not null,
    last_name char(50) not null,
    looking_for char(25) not null,
    school char(140) default '',
    password char(64) not null,
    salt char(32) not null, 
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);

create table decisions (
    from_user_id int not null references users(id),
    to_user_id int not null references users(id),
    likes bool not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint pk_decision primary key (from_user_id, to_user_id)
);

create table matches (
    id serial,
    user_1_id int not null references users(id),
    user_2_id int not null references users(id),
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    constraint pk_match primary key (user_1_id, user_2_id)
);

create table images (
    user_id int not null references users(id),
    index int not null,
    small bool default false,
    url char(2000) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    expires_at timestamp not null,
    constraint pk_image primary key (user_id, index)
);
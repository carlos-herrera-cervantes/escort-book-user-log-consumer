create table "user" (
    id varchar(100) not null,
    user_id varchar(100) not null unique,
    first_name varchar(100) null,
    last_name varchar(100) null,
    email varchar(100) not null,
    deleted boolean null default false,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key(id)
);

create table avatar (
    id varchar(100) not null,
    user_id varchar(100) not null,
    path varchar(1000) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key(id)
);

create table connection_log (
    id varchar(100) not null,
    user_id varchar(100) not null,
    last_connection timestamp not null,
    primary key(id)
);

create table request_log (
    id varchar(100) not null,
    user_id varchar(100) not null,
    component varchar(50) not null,
    path varchar(100) not null,
    method varchar(15) not null,
    payload varchar(100) not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key(id)
);

create table dictum (
    id varchar(100) not null,
    user_id varchar(100) not null,
    from_user varchar(100) not null,
    status_category_id varchar(100) not null,
    comment text not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key(id)
);

create table status_category (
    id varchar(100) not null,
    name varchar(100) not null,
    description varchar(200) not null,
    active boolean null default true,
    created_at timestamp not null,
    updated_at timestamp not null,
    primary key(id)
);

alter table avatar add foreign key(user_id) references "user"(user_id);
alter table connection_log add foreign key(user_id) references "user"(user_id);
alter table request_log add foreign key(user_id) references "user"(user_id);
alter table dictum add foreign key(user_id) references "user"(user_id);
alter table dictum add foreign key(from_user) references "user"(user_id);
alter table dictum add foreign key(status_category_id) references status_category(id);

insert into status_category values (
    'd2ed1136-27bf-41f9-a1bb-26735be303d6',
    'Active',
    'When the user is working',
    true,
    '2022-01-01',
    '2022-01-01'
);
insert into status_category values (
    'b086dacb-f5e9-4540-a40e-510c5f993299',
    'Inactive',
    'When the user goes to vacation or was suspended',
    true,
    '2022-01-01',
    '2022-01-01'
);
insert into status_category values (
    '8344df37-eb58-43a9-ba56-440b4c5d5637',
    'Locked',
    'When the user was locked for some reason',
    true,
    '2022-01-01',
    '2022-01-01'
);
insert into status_category values (
    'cc8bb902-0f75-4631-a2d0-6c48a26f5617',
    'Deleted',
    'When the user no longer works',
    true,
    '2022-01-01',
    '2022-01-01'
);
insert into status_category values (
    '95176cb6-f76a-4dba-9ac8-09e81ee07003',
    'Evaluation',
    'When the user is being monitored to check if he does not make mistakes',
    true,
    '2022-01-01',
    '2022-01-01'
);

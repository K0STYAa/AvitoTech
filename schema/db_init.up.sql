CREATE TABLE users
(
    id            serial       not null unique,
    balance       int          not null
);

CREATE TABLE history
(
    id          serial                                           not null unique,
    sender_id   int references users (id) on delete cascade      not null,
    receiver_id int references users (id) on delete cascade      not null,
    amount      int                                              not null
);
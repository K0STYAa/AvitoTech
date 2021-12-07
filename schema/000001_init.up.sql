CREATE TABLE users
(
    id            serial       CHECK (id >= 0) unique,
    balance       int          not null
);

CREATE TABLE history
(
    id             serial                                           not null unique,
    sender_id      int references users (id) on delete cascade      CHECK (sender_id >= 0),
    receiver_id    int references users (id) on delete cascade      CHECK (sender_id >= 0),
    amount         int                                              not null,
    departure_time timestamp                                        not null
);
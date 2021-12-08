CREATE TABLE users
(
    id            serial              CHECK (id >= 0) unique,
    balance       double precision    CHECK (balance >= 0)
);

CREATE TABLE history
(
    id             serial                                           not null unique,
    sender_id      int references users (id) on delete cascade      CHECK (sender_id >= 0),
    receiver_id    int references users (id) on delete cascade      CHECK (receiver_id >= 0),
    amount         double precision                                 CHECK (amount > 0),
    departure_time timestamp                                        not null
);
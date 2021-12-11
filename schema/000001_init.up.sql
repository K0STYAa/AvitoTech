CREATE TABLE users
(
    id            serial PRIMARY KEY  CHECK (id >= 0) unique,
    balance       int                 CHECK (balance >= 0)
);

CREATE TABLE history
(
    id             serial PRIMARY KEY                               not null unique,
    sender_id      int references users (id) on delete cascade      CHECK (sender_id >= 0),
    receiver_id    int references users (id) on delete cascade      CHECK (receiver_id >= 0),
    amount         int                                              CHECK (amount > 0),
    departure_time timestamp                                        not null
);

CREATE INDEX s_r_id ON history(sender_id, receiver_id);
CREATE INDEX amount ON history(amount);
CREATE INDEX departure_time ON history(departure_time);
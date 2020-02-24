create table if not exists item (
        id primary key,
        name text not null,
        person_id text not null,
        amount integer not null,
        date text not null
);

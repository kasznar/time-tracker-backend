create table date_day
(
    id   int auto_increment
        primary key,
    type int  not null,
    date date not null,
    constraint date_day_date_day_type_id_fk
        foreign key (type) references date_day_type (id)
);

INSERT INTO GO_DB.date_day (id, type, date) VALUES (1, 1, '2021-05-25');
INSERT INTO GO_DB.date_day (id, type, date) VALUES (2, 2, '2021-05-26');
INSERT INTO GO_DB.date_day (id, type, date) VALUES (3, 3, '2021-05-27');
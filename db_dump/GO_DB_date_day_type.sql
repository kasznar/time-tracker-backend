create table date_day_type
(
    id   int auto_increment
        primary key,
    name text null
);

INSERT INTO GO_DB.date_day_type (id, name) VALUES (1, 'workday');
INSERT INTO GO_DB.date_day_type (id, name) VALUES (2, 'holiday');
INSERT INTO GO_DB.date_day_type (id, name) VALUES (3, 'mandatory_vacation');
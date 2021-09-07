create table workday_type
(
    id   int auto_increment
        primary key,
    name text null
);

INSERT INTO GO_DB.workday_type (id, name) VALUES (1, 'workday');
INSERT INTO GO_DB.workday_type (id, name) VALUES (2, 'vacation');
INSERT INTO GO_DB.workday_type (id, name) VALUES (3, 'sick_leave');
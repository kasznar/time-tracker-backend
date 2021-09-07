create table user_status
(
    id   int auto_increment
        primary key,
    name text null
);

INSERT INTO GO_DB.user_status (id, name) VALUES (1, 'active');
INSERT INTO GO_DB.user_status (id, name) VALUES (2, 'inactive');
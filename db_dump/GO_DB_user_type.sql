create table user_type
(
    id   int auto_increment
        primary key,
    name text null
);

INSERT INTO GO_DB.user_type (id, name) VALUES (1, 'employee');
INSERT INTO GO_DB.user_type (id, name) VALUES (2, 'team_leader');
INSERT INTO GO_DB.user_type (id, name) VALUES (3, 'hr');
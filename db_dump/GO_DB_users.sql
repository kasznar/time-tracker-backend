create table users
(
    id                int auto_increment
        primary key,
    email             text not null,
    password          text not null,
    first_day         date not null,
    default_day_start time null,
    default_day_end   time null,
    work_hours        int  null,
    type              int  not null,
    status            int  null,
    last_day          date null,
    constraint users_user_status_id_fk
        foreign key (status) references user_status (id),
    constraint users_user_type_id_fk
        foreign key (type) references user_type (id)
);

INSERT INTO GO_DB.users (id, email, password, first_day, default_day_start, default_day_end, work_hours, type, status, last_day) VALUES (2, 'kecske@kecske.com', 'asd', '2021-05-01', '09:00:00', '17:00:00', 8, 1, 1, '2021-05-31');
INSERT INTO GO_DB.users (id, email, password, first_day, default_day_start, default_day_end, work_hours, type, status, last_day) VALUES (3, 'asd@asd.hu', 'asd2', '2021-05-12', '08:00:00', '16:00:00', 6, 2, 1, '2021-05-31');
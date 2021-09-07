create table workday
(
    id             int auto_increment
        primary key,
    type           int  not null,
    summary        text null,
    date_day       int  not null,
    user_id        int  not null,
    request_status int  not null,
    constraint workday_date_day_id_fk
        foreign key (date_day) references date_day (id),
    constraint workday_users_id_fk
        foreign key (user_id) references users (id),
    constraint workday_workday_type_id_fk
        foreign key (type) references workday_type (id)
);


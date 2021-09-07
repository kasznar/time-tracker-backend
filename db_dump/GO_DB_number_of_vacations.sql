create table number_of_vacations
(
    id      int auto_increment
        primary key,
    year    year not null,
    user_id int  not null,
    constraint number_of_vacations_users_id_fk
        foreign key (user_id) references users (id)
);


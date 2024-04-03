create table r_user_file_root
(
    id           bigint unsigned auto_increment
        primary key,
    user_id      bigint unsigned                     not null,
    tree_root_id bigint unsigned                     not null,
    create_at    timestamp default CURRENT_TIMESTAMP null,
    update_at    timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_at    timestamp                           null,
    constraint r_user_file_root_nk_user_id
        unique (user_id),
    constraint r_user_file_root_uindex_user_id
        unique (user_id)
)
    comment 'map user with file root id';

create table t_file_tree
(
    id        bigint unsigned auto_increment
        primary key,
    name      varchar(1024)                       not null,
    data_id   bigint unsigned                     null,
    parent    bigint unsigned                     null comment '0 root / else other',
    kind      tinyint   default 1                 not null comment '0 file 1 directory',
    share_url varchar(512)                        null,
    create_at timestamp default CURRENT_TIMESTAMP null,
    update_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_at timestamp                           null,
    constraint t_file_tree_nk_data_id
        unique (data_id),
    constraint t_file_tree_nk_parent
        unique (parent)
);

create table t_user
(
    ID        bigint unsigned auto_increment
        primary key,
    username  char(50)                            null,
    password  char(100)                           null,
    email     varchar(1024)                       null,
    phone     varchar(30)                         null,
    create_at timestamp default CURRENT_TIMESTAMP null,
    update_at timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    delete_at timestamp                           null,
    enable    tinyint   default 0                 null comment '0 启用 1 冻结',
    constraint t_user_k_delete_time
        unique (delete_at),
    constraint t_user_k_phone
        unique (phone),
    constraint t_user_username_index
        unique (username)
);


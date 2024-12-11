create schema lesson collate utf8mb4_uca1400_ai_ci;


-- auto-generated definition
create table apply_lists
(
    id         bigint unsigned auto_increment
        primary key,
    mid        varchar(255)    not null comment '申請人',
    lid        varchar(255)    not null comment '申請課程ID',
    payed      enum ('0', '1') not null comment '付款狀態',
    created_at timestamp       null,
    updated_at timestamp       null
)
    collate = utf8mb4_unicode_ci;

create table lessons
(
    id             bigint unsigned auto_increment
        primary key,
    lessonName     varchar(255) not null comment '課程名稱',
    lessonDescribe longtext     not null comment '課程簡介',
    mid            int          not null comment '建立者',
    teacher        varchar(255) not null comment '課程老師',
    lessonTime     varchar(255) not null comment '課程時間',
    lessonAddress  varchar(255) not null comment '課程地點',
    tuitionFee     varchar(255) not null comment '課程費用',
    email          varchar(255) not null comment '聯絡信箱',
    created_at     timestamp    null,
    updated_at     timestamp    null
)
    collate = utf8mb4_unicode_ci;


create table members
(
    id         bigint unsigned auto_increment
        primary key,
    account    varchar(255)               not null comment '帳號',
    firstName  varchar(255)               not null comment '姓',
    lastName   varchar(255)               not null comment '名',
    email      varchar(255)               not null comment '信箱',
    phone      text                       null comment '手機',
    password   varchar(255)               not null comment '密碼',
    memberType enum ('student', 'school') not null comment '會員類別',
    created_at timestamp                  null,
    updated_at timestamp                  null,
    constraint members_account_unique
        unique (account)
)
    collate = utf8mb4_unicode_ci;


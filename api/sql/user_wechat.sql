SET NAMES 'utf8mb4';

CREATE TABLE `user_wechat_open`
(
    `open_id`       varchar(128)    NOT NULL,
    `app_id`        varchar(128)    NOT NULL,
    `user_id`       int(11),
    `sex`           int(11)         NOT NULL,
    `state`         int(11)         NOT NULL,

    `nick_name`     varchar(200)    NOT NULL,
    `country`       varchar(100)    NOT NULL,
    `province`      varchar(100)    NOT NULL,
    `city`          varchar(100)    NOT NULL,
    `head_img_url`  varchar(1024)   NOT NULL,

    `created_at`    datetime        NOT NULL,
    `updated_at`    datetime        NOT NULL,
    PRIMARY KEY (`open_id`, `app_id`),
    KEY (`user_id`),
    KEY (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

CREATE TABLE `user_wechat_mini_app`
(
    `open_id`       varchar(128)    NOT NULL,
    `app_id`        varchar(128)    NOT NULL,
    `user_id`       int(11),
    `sex`           int(11)         NOT NULL,
    `state`         int(11)         NOT NULL,

    `nick_name`     varchar(200)    NOT NULL,
    `country`       varchar(100)    NOT NULL,
    `province`      varchar(100)    NOT NULL,
    `city`          varchar(100)    NOT NULL,
    `head_img_url`  varchar(1024)   NOT NULL,

    `created_at`    datetime        NOT NULL,
    `updated_at`    datetime        NOT NULL,
    PRIMARY KEY (`open_id`, `app_id`),
    KEY(`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

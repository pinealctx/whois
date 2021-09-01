SET NAMES 'utf8mb4';

CREATE TABLE `user`
(
    `id`            int(11)         NOT NULL,
    `state`         int(11)         NOT NULL,
    `nick_name`     varchar(200)    NOT NULL,
    `avatar`        varchar(1024)   NOT NULL,
    `created_at`    datetime        NOT NULL,
    `updated_at`    datetime        NOT NULL,
    PRIMARY KEY (`id`),
    KEY (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

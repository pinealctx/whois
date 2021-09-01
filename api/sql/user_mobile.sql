SET NAMES 'utf8mb4';

CREATE TABLE `user_mobile`
(
    `mobile`        varchar(30) NOT NULL,
    `area`          varchar(8)  NOT NULL,
    `user_id`       int(11),
    `state`         int(11)     NOT NULL,
    `created_at`    datetime    NOT NULL,
    `updated_at`    datetime    NOT NULL,
    PRIMARY KEY (`mobile`, `area`),
    KEY (`user_id`),
    KEY (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

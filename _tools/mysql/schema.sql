CREATE TABLE `user`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '사용자 식별자',
    `name`     varchar(20) NOT NULL COMMENT '사용자 이름',
    `password` VARCHAR(80) NOT NULL COMMENT '비밀번호 해시',
    `role`     VARCHAR(80) NOT NULL COMMENT '역할',
    `created`  DATETIME(6) NOT NULL COMMENT '레코드 생성 시각',
    `modified` DATETIME(6) NOT NULL COMMENT '레코드 수정 시각',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='사용자';

CREATE TABLE `task`
(
    `id`       BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '작업 식별자',
    `title`    VARCHAR(128) NOT NULL COMMENT '작업 제목',
    `status`   VARCHAR(20)  NOT NULL COMMENT '작업 상태',
    `created`  DATETIME(6) NOT NULL COMMENT '레코드 생성 시각',
    `modified` DATETIME(6) NOT NULL COMMENT '레코드 수정 시각',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='작업';

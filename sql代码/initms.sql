/*
这里在catms下面创建所有的表
CREATE SCHEMA catms
*/
CREATE TABLE user_info (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    gender VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    age INT UNSIGNED,
    avatar VARCHAR(255)
);

CREATE TABLE cat_info (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    gender VARCHAR(255) NOT NULL,
    date VARCHAR(255) NOT NULL,
    weight INT,
    sterilization TINYINT(1) NOT NULL,
    adapt TINYINT(1) NOT NULL,
    avatar VARCHAR(255),
    owner INT UNSIGNED,
    FOREIGN KEY (owner) REFERENCES user_info(id)
);

CREATE TABLE adapt_info (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    owner_id INT UNSIGNED,
    cat_id INT UNSIGNED,
    date VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES user_info(id),
    FOREIGN KEY (cat_id) REFERENCES cat_info(id)
);

CREATE TABLE home_info (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    cat_id INT UNSIGNED,
    cat_name VARCHAR(255),
    cat_avatar VARCHAR(255),
    tot_amount INT NOT NULL,
    cat_male INT NOT NULL,
    cat_female INT NOT NULL,
    cat_adapt INT NOT NULL,
    cat_sterilization INT NOT NULL,
    date VARCHAR(255) NOT NULL,
    FOREIGN KEY (cat_id) REFERENCES cat_info(id)
);

INSERT INTO catms.adapt_infos (id, owner_id, cat_id, date, status) VALUES (1, 1, 1, '2023/12/29', '已同意');

INSERT INTO catms.cat_infos (id, name, type, gender, date, weight, sterilization, adapt, avatar, owner) VALUES (1, '小橘', '橘猫', '公', '2022-04-09', 5, 0, 1, 'http://localhost:9090/uploads/catavatars/5b0b023a636fd9327083f82556086ccd.jpg', 1);
INSERT INTO catms.cat_infos (id, name, type, gender, date, weight, sterilization, adapt, avatar, owner) VALUES (2, '小狸花', '狸花猫', '母', '2021-11-24', 3, 1, 0, 'http://localhost:9090/uploads/catavatars/da74607c3c0231432c81b5a676eb7da7.jpg', 0);
INSERT INTO catms.cat_infos (id, name, type, gender, date, weight, sterilization, adapt, avatar, owner) VALUES (3, '小奶牛', '奶牛猫', '公', '2023-01-14', 1, 0, 0, 'http://localhost:9090/uploads/catavatars/087840630a4d00324ca271058a3d8830.jpg', 0);
INSERT INTO catms.cat_infos (id, name, type, gender, date, weight, sterilization, adapt, avatar, owner) VALUES (5, '小三花', '三花猫', '母', '2022-08-02', 3, 1, 0, 'http://localhost:9090/uploads/catavatars/2adc9262b6eb398d4b7e550ef7e4cb33.jpg', 0);
INSERT INTO catms.cat_infos (id, name, type, gender, date, weight, sterilization, adapt, avatar, owner) VALUES (8, '小奶牛', '长毛奶牛猫', '母', '2023-11-07', 2, 0, 0, 'http://localhost:9090/uploads/catavatars/762081b6636c39b7cbce88381583edf2.jpg', 0);

INSERT INTO catms.home_infos (id, cat_id, cat_name, cat_avatar, tot_amount, cat_male, cat_female, cat_adapt, cat_sterilization, date) VALUES (8, 1, '小橘', 'http://localhost:9090/uploads/catavatars/5b0b023a636fd9327083f82556086ccd.jpg', 5, 3, 2, 1, 0, '2023-12-29 21:17:08');
INSERT INTO catms.home_infos (id, cat_id, cat_name, cat_avatar, tot_amount, cat_male, cat_female, cat_adapt, cat_sterilization, date) VALUES (9, 1, '小橘', 'http://localhost:9090/uploads/catavatars/5b0b023a636fd9327083f82556086ccd.jpg', 5, 3, 2, 1, 0, '2023-12-29 21:18:55');
INSERT INTO catms.home_infos (id, cat_id, cat_name, cat_avatar, tot_amount, cat_male, cat_female, cat_adapt, cat_sterilization, date) VALUES (10, 1, '小橘', 'http://localhost:9090/uploads/catavatars/5b0b023a636fd9327083f82556086ccd.jpg', 5, 3, 2, 1, 0, '2023-12-29 21:19:02');
INSERT INTO catms.home_infos (id, cat_id, cat_name, cat_avatar, tot_amount, cat_male, cat_female, cat_adapt, cat_sterilization, date) VALUES (11, 1, '小橘', 'http://localhost:9090/uploads/catavatars/5b0b023a636fd9327083f82556086ccd.jpg', 5, 3, 2, 1, 2, '2023-12-29 21:19:16');
INSERT INTO catms.home_infos (id, cat_id, cat_name, cat_avatar, tot_amount, cat_male, cat_female, cat_adapt, cat_sterilization, date) VALUES (12, 1, '小橘', 'http://localhost:9090/uploads/catavatars/5b0b023a636fd9327083f82556086ccd.jpg', 5, 2, 3, 1, 2, '2023-12-29 22:13:53');

INSERT INTO catms.user_infos (id, name, password, gender, phone, age, avatar) VALUES (1, '小瓜', 'Gua#gua', '男', '13812345678', 20, 'http://localhost:9090/uploads/useravatars/5ced28e98064c5ee0c08fd6f6cf18e07.jpg');
INSERT INTO catms.user_infos (id, name, password, gender, phone, age, avatar) VALUES (2, '小汪', 'Woof#woof', '女', '13912345678', 25, 'http://localhost:9090/uploads/useravatars/180e01693cddaa64e86df3218d1d46a6.jpg');
INSERT INTO catms.user_infos (id, name, password, gender, phone, age, avatar) VALUES (3, '小熊', 'Bear#Bear', '男', '13612345678', 22, 'http://localhost:9090/uploads/useravatars/d71ca3a3e1f9c4687ef4583be4612578.jpg');
INSERT INTO catms.user_infos (id, name, password, gender, phone, age, avatar) VALUES (4, '小猪', 'Piggy#piggy', '女', '13712345678', 28, 'http://localhost:9090/uploads/useravatars/c4c2bc90f76246a96ff1e3ad6aa75dbe.jpg');



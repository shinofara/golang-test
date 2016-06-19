use wordpress;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  name varchar(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

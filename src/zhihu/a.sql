USE zhihu_user;
CREATE TABLE IF NOT EXISTS `user` (
  `my_id` int(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `is_followed` boolean DEFAULT FALSE,
  `avatar_url_template` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `user_type` varchar(255) DEFAULT NULL,
  `answer_count` int(255) DEFAULT 0,
  `url_token` varchar(255) DEFAULT NULL,
  `is_advertiser` boolean DEFAULT FALSE,
  `avatar_url` varchar(255) DEFAULT NULL,
  `is_following` boolean DEFAULT FALSE,
  `is_org` boolean DEFAULT FALSE,
  `headline` varchar(255) DEFAULT NULL,
  `follower_count` varchar(255) DEFAULT NULL,
  `the_type` varchar(255) DEFAULT NULL,
  `id` varchar(255) DEFAULT NULL,
  `articles_count` int(255) DEFAULT NULL,
  PRIMARY KEY (`my_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=3 ;
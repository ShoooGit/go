use sampledb;

CREATE TABLE users
(
	id char(28) NOT NULL COMMENT 'ID',
	name text NOT NULL COMMENT '名前',
	email text NOT NULL COMMENT 'メールアドレス',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
  deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ユーザ';


--
-- テーブルのデータのダンプ `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `created_at`, `updated_at`, `deleted_at`) VALUES
('1', 'hoge', 'hoge@gmail.com', '2018-12-22 20:04:31', '2018-12-22 20:04:31', NULL),
('2', 'fuga', 'fuga@gmail.com', '2018-12-31 10:08:01', '2018-01-01 22:30:20', NULL),
('3', 'yamada tarou', 'taro@xxx.com', '2019-04-05 10:08:01', '2019-04-05 10:08:01', NULL);
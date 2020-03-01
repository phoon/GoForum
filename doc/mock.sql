/**
    users table
*/
INSERT INTO `users`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `nickname`, `email`, `password`, `bio`, `role`, `avatar_url`, `status`) VALUES (1, '2020-03-01 17:47:12', '2020-03-01 17:47:12', NULL, 'admin', 'admin', 'admin@mail.com', '$2a$10$ZCZkSwGpDl61Ij2uEEBmZe8BHXPXcyFNtoG3l7TwyFVdk2meH3WPS', '我是管理员哦', 1, 'https://avatars1.githubusercontent.com/u/23262111?v=4', 0);
INSERT INTO `users`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `nickname`, `email`, `password`, `bio`, `role`, `avatar_url`, `status`) VALUES (2, '2020-03-01 17:49:18', '2020-03-01 17:49:18', NULL, 'user1', '大雄', 'dx@mail.com', '$2a$10$yH3jFYHyx53.UmsEIFY.yuSPQZfnHD7XUlIKR0lZVTMEBPWpptzQG', '我是大雄', 0, 'https://avatars1.githubusercontent.com/u/23262111?v=4', 0);
INSERT INTO `users`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `nickname`, `email`, `password`, `bio`, `role`, `avatar_url`, `status`) VALUES (5, '2020-03-01 17:49:42', '2020-03-01 17:49:42', NULL, 'user2', '静香', 'jx@mail.com', '$2a$10$kVxb0hbELxUxll/ow8C4JulAcN4FDySoJHvI5YjoXrybTbbCdE2fO', '这个用户很懒，什么都没留下', 0, 'https://avatars1.githubusercontent.com/u/23262111?v=4', 0);
INSERT INTO `users`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `nickname`, `email`, `password`, `bio`, `role`, `avatar_url`, `status`) VALUES (6, '2020-03-01 17:50:07', '2020-03-01 17:50:07', NULL, 'user3', '胖虎', 'ph@mail.com', '$2a$10$6lR4UxU.ybsloEJNPtdQ2OtcwP15/Hi6MsciMIroKUNPhtvxTelHi', 'Hey', 0, 'https://avatars1.githubusercontent.com/u/23262111?v=4', 0);

/**
    user_followers table
*/
INSERT INTO `user_followers`(`user_id`, `follower_id`) VALUES (2, 5);
INSERT INTO `user_followers`(`user_id`, `follower_id`) VALUES (5, 2);

/**
    user_following table
*/
INSERT INTO `user_following`(`user_id`, `following_id`) VALUES (2, 5);
INSERT INTO `user_following`(`user_id`, `following_id`) VALUES (5, 2);

/**
    categories table
*/
INSERT INTO `categories`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `description`, `status`) VALUES (1, '2020-03-01 17:55:13', '2020-03-01 18:02:15', NULL, '未分类', '未分类话题集合', 0);
INSERT INTO `categories`(`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `description`, `status`) VALUES (2, '2020-03-01 17:56:12', '2020-03-01 18:02:41', NULL, 'cate1', '随便写点', 0);

/**
    topics table
*/
INSERT INTO `topics`(`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `summary`, `user_id`, `category_id`, `raw_conent`, `html_content`, `status`) VALUES (2, '2020-03-01 18:04:54', '2020-03-01 18:04:54', NULL, '春江花月夜 (張若虛)', '春江潮水連海平，海上明月共潮生。\n灩灩隨波千萬里，何處春江無月明？\n江流宛轉遶芳甸，月照花林皆似霰。\n空裏流霜不覺飛，汀上白沙看不見。\n江天一色無纖塵，皎皎空中孤月輪。\n江畔何人初見月，江月何年初照人？\n人生代代無窮已，江月年年祇相似。', 2, 1, '春江潮水連海平，海上明月共潮生。\n灩灩隨波千萬里，何處春江無月明？\n江流宛轉遶芳甸，月照花林皆似霰。\n空裏流霜不覺飛，汀上白沙看不見。\n江天一色無纖塵，皎皎空中孤月輪。\n江畔何人初見月，江月何年初照人？\n人生代代無窮已，江月年年祇相似。\n不知江月待何人？但見長江送流水。\n白雲一片去悠悠，青楓浦上不勝愁。\n誰家今夜扁舟子，何處相思明月樓？\n可憐樓上月徘徊，應照離人妝鏡臺。\n玉戶簾中卷不去，擣衣砧上拂還來。\n此時相望不相聞，願逐月華流照君。\n鴻雁長飛光不度，魚龍潛躍水成文。\n昨夜閑潭夢落花，可憐春半不還家。\n江水流春去欲盡，江潭落月復西斜。\n斜月沉沉藏海霧，碣石瀟湘無限路。\n不知乘月幾人歸，落月搖情滿江樹。', '春江潮水連海平，海上明月共潮生。\n灩灩隨波千萬里，何處春江無月明？\n江流宛轉遶芳甸，月照花林皆似霰。\n空裏流霜不覺飛，汀上白沙看不見。\n江天一色無纖塵，皎皎空中孤月輪。\n江畔何人初見月，江月何年初照人？\n人生代代無窮已，江月年年祇相似。\n不知江月待何人？但見長江送流水。\n白雲一片去悠悠，青楓浦上不勝愁。\n誰家今夜扁舟子，何處相思明月樓？\n可憐樓上月徘徊，應照離人妝鏡臺。\n玉戶簾中卷不去，擣衣砧上拂還來。\n此時相望不相聞，願逐月華流照君。\n鴻雁長飛光不度，魚龍潛躍水成文。\n昨夜閑潭夢落花，可憐春半不還家。\n江水流春去欲盡，江潭落月復西斜。\n斜月沉沉藏海霧，碣石瀟湘無限路。\n不知乘月幾人歸，落月搖情滿江樹。', 0);
INSERT INTO `topics`(`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `summary`, `user_id`, `category_id`, `raw_conent`, `html_content`, `status`) VALUES (3, '2020-03-01 18:06:19', '2020-03-01 18:06:19', NULL, '测试文章', '啦啦啦啦啦', 5, 2, '啦啦啦啦啦的v上的v三', '的v上的vvvv王v额额为王v无法王v额发', 0);

/**
    user_liked_topics table
*/
INSERT INTO `user_liked_topics`(`user_id`, `topic_id`) VALUES (5, 2);

/**
    comments table
*/
INSERT INTO `comments`(`id`, `created_at`, `updated_at`, `deleted_at`, `content`, `user_id`, `topic_id`, `quote_id`, `status`) VALUES (1, '2020-03-01 18:25:52', '2020-03-01 18:25:52', NULL, '這詩太牛逼了！！！', 2, 2, 0, 0);
INSERT INTO `comments`(`id`, `created_at`, `updated_at`, `deleted_at`, `content`, `user_id`, `topic_id`, `quote_id`, `status`) VALUES (2, '2020-03-01 18:27:08', '2020-03-01 18:27:08', NULL, '就是就是，大雄你啥时候把作业还给我', 5, 2, 1, 0);

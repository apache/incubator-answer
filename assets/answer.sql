
-- --------------------------------------------------------

--
-- 表的结构 `activity`
--

CREATE TABLE `activity` (
  `id` bigint(20) NOT NULL COMMENT 'Activity ID autoincrement',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `user_id` bigint(20) NOT NULL COMMENT 'the user ID that generated the activity or affected by the activity',
  `trigger_user_id` bigint(20) NOT NULL DEFAULT '0',
  `object_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'the object ID that affected by the activity',
  `activity_type` int(11) NOT NULL COMMENT 'activity type, correspond to config id',
  `cancelled` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'mark this activity if cancelled or not,default 0(not cancelled)',
  `rank` int(11) NOT NULL DEFAULT '0' COMMENT 'rank of current operating user affected',
  `has_rank` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'this activity has rank or not'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='activity';

-- --------------------------------------------------------

--
-- 表的结构 `answer`
--

CREATE TABLE `answer` (
  `id` bigint(20) NOT NULL COMMENT 'answer id',
  `question_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'question id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'answer user id',
  `original_text` mediumtext NOT NULL COMMENT 'original text',
  `parsed_text` mediumtext NOT NULL COMMENT 'parsed text',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT ' answer status(available: 1; deleted: 10)',
  `adopted` int(11) NOT NULL DEFAULT '1' COMMENT 'adopted (1 failed 2 adopted)',
  `comment_count` int(11) NOT NULL DEFAULT '0' COMMENT 'comment count',
  `vote_count` int(11) NOT NULL DEFAULT '0' COMMENT 'vote count',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `revision_id` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='answer';

-- --------------------------------------------------------

--
-- 表的结构 `collection`
--

CREATE TABLE `collection` (
  `id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'collection id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'user id',
  `object_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'object id',
  `user_collection_group_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'user collection group id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='collection';

-- --------------------------------------------------------

--
-- 表的结构 `collection_group`
--

CREATE TABLE `collection_group` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL DEFAULT '0',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'the collection group name',
  `default_group` int(11) NOT NULL DEFAULT '1' COMMENT 'mark this group is default, default 1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='collection group';

-- --------------------------------------------------------

--
-- 表的结构 `comment`
--

CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL COMMENT 'comment id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'user id',
  `reply_user_id` bigint(20) DEFAULT NULL COMMENT 'reply user id',
  `reply_comment_id` bigint(20) DEFAULT NULL COMMENT 'reply comment id',
  `object_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'object id',
  `question_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'question id',
  `vote_count` int(11) NOT NULL DEFAULT '0' COMMENT 'user vote amount',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'comment status(available: 0; deleted: 10)',
  `original_text` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'original comment content',
  `parsed_text` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'parsed comment content'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='comment';

-- --------------------------------------------------------

--
-- 表的结构 `config`
--

CREATE TABLE `config` (
  `id` int(11) NOT NULL COMMENT 'config id',
  `key` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'the config key',
  `value` text COLLATE utf8mb4_unicode_ci COMMENT 'the config value, custom data structures and types'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='config';

-- --------------------------------------------------------

--
-- 表的结构 `meta`
--

CREATE TABLE `meta` (
  `id` int(10) UNSIGNED NOT NULL COMMENT 'id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
  `object_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'object id',
  `key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'key',
  `value` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'value'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='meta';

-- --------------------------------------------------------

--
-- 表的结构 `notification`
--

CREATE TABLE `notification` (
  `id` bigint(20) NOT NULL COMMENT 'notification id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'user_id',
  `object_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'object id',
  `content` text NOT NULL COMMENT 'notification content',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '1 inbox 2 achievements',
  `is_read` int(11) NOT NULL DEFAULT '1' COMMENT 'read status(unread: 1; read 2)',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT 'notification status(normal: 1; delete 2)',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='notification';

-- --------------------------------------------------------

--
-- 表的结构 `notification_read`
--

CREATE TABLE `notification_read` (
  `id` int(11) NOT NULL COMMENT 'id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'user id',
  `message_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'message id',
  `is_read` int(11) NOT NULL DEFAULT '1' COMMENT 'read status(unread: 1; read 2)'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='notification read record';

-- --------------------------------------------------------

--
-- 表的结构 `question`
--

CREATE TABLE `question` (
  `id` bigint(20) NOT NULL COMMENT 'question id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'user id',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'question title',
  `original_text` mediumtext NOT NULL COMMENT 'original text',
  `parsed_text` mediumtext NOT NULL COMMENT 'parsed text',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT ' question status(available: 1; deleted: 10)',
  `view_count` int(11) NOT NULL DEFAULT '0' COMMENT 'view count',
  `unique_view_count` int(11) NOT NULL DEFAULT '0' COMMENT 'unique view count',
  `vote_count` int(11) NOT NULL DEFAULT '0' COMMENT 'vote count',
  `answer_count` int(11) NOT NULL DEFAULT '0' COMMENT 'answer count',
  `collection_count` int(11) NOT NULL DEFAULT '0' COMMENT 'collection count',
  `follow_count` int(11) NOT NULL DEFAULT '0' COMMENT 'follow count',
  `accepted_answer_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'accepted answer id',
  `last_answer_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'last answer id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
  `post_update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'answer the last update time',
  `revision_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'revision id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='question';

-- --------------------------------------------------------

--
-- 表的结构 `report`
--

CREATE TABLE `report` (
  `id` bigint(20) NOT NULL COMMENT 'id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `user_id` bigint(20) NOT NULL COMMENT 'reporter user id',
  `object_id` bigint(20) NOT NULL COMMENT 'object id',
  `reported_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'reported user id',
  `object_type` int(11) NOT NULL DEFAULT '0' COMMENT 'revision type',
  `report_type` int(11) NOT NULL DEFAULT '0' COMMENT 'report type',
  `content` text NOT NULL COMMENT 'report content',
  `flaged_type` int(11) NOT NULL DEFAULT '0',
  `flaged_content` text,
  `status` int(11) NOT NULL DEFAULT '1' COMMENT 'status(normal: 1; pending:2; delete: 10)'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='report';

-- --------------------------------------------------------

--
-- 表的结构 `revision`
--

CREATE TABLE `revision` (
  `id` bigint(20) NOT NULL COMMENT 'id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'user id',
  `object_type` int(11) NOT NULL DEFAULT '0' COMMENT 'revision type(question: 1; answer 2)',
  `object_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'object id',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'title',
  `content` text NOT NULL COMMENT 'content',
  `log` varchar(255) DEFAULT NULL,
  `status` int(11) NOT NULL DEFAULT '1' COMMENT 'revision status(normal: 1; delete 2)'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='revision';

-- --------------------------------------------------------

--
-- 表的结构 `site_info`
--

CREATE TABLE `site_info` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
  `type` varchar(64) DEFAULT NULL,
  `content` mediumtext,
  `status` int(11) NOT NULL DEFAULT '1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- 表的结构 `tag`
--

CREATE TABLE `tag` (
  `id` bigint(20) NOT NULL COMMENT 'tag_id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `main_tag_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'main tag id',
  `main_tag_slug_name` varchar(50) NOT NULL DEFAULT '' COMMENT 'main tag slug name',
  `slug_name` varchar(50) NOT NULL DEFAULT '' COMMENT 'slug name',
  `display_name` varchar(50) NOT NULL DEFAULT '' COMMENT 'display name',
  `original_text` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'original comment content',
  `parsed_text` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'parsed comment content',
  `follow_count` int(11) NOT NULL DEFAULT '0' COMMENT 'follow count',
  `question_count` int(11) NOT NULL COMMENT 'question count',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT 'tag status(available: 1; deleted: 10)',
  `revision_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'revision id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='tag';

--
-- 转存表中的数据 `tag`
--

INSERT INTO `tag` (`id`, `created_at`, `updated_at`, `main_tag_id`, `main_tag_slug_name`, `slug_name`, `display_name`, `original_text`, `parsed_text`, `follow_count`, `question_count`, `status`, `revision_id`) VALUES
(10030000000000007, '2022-09-07 02:07:16', '2022-09-28 06:56:32', 10030000000000364, 'javascript', 'js', 'Js', 'js', 'Js', 7, 21, 1, 90),
(10030000000000008, '2022-09-07 02:07:16', '2022-09-19 08:40:36', 0, '', 'php', 'PHP', '2122121对外只有3001一个端口，群集内有负载策略会分发给不同实例2121211212121212121', '<p>2122121对外只有3001一个端口，群集内有负载策略会分发给不同实例2121211212121212121</p>\n', 2, 3, 10, 101),
(10030000000000009, '2022-09-07 02:07:16', '2022-09-28 08:00:48', 0, '', 'go', 'Go', 'abcbcsasasasasa', '<p>abcbcsasasasasa</p>\n', 7, 17, 1, 259),
(10030000000000010, '2022-09-07 02:07:16', '2022-09-28 05:16:00', 0, '', 'apple', 'Apple', '对外只有3001一个端口，群集内有负载策略会分发给不同实例', '<p>对外只有3001一个端口，群集内有负载策略会分发给不同实例</p>', 5, 3, 1, 0),
(10030000000000113, '2022-09-07 02:07:16', '2022-09-26 08:44:04', 0, '', 'lua', 'lua', '对外只有3001一个端口，群集内有负载策略会分发给不同实例', '<p>对外只有3001一个端口，群集内有负载策略会分发给不同实例</p>', 4, 4, 1, 0),
(10030000000000114, '2022-09-07 02:07:16', '2022-09-27 09:57:35', 0, '', 'dell', 'dell', '对外只有3001一个端口，群集内有负载策略会分发给不同实例', '<p>对外只有3001一个端口，群集内有负载策略会分发给不同实例</p>', 3, 0, 1, 0),
(10030000000000118, '2022-09-07 02:07:16', '2022-09-28 08:20:13', 0, '', 'dells', 'dells', '对外只有3001一个端口，群集内有负载策略会分发给不同实例', '<p>对外只有3001一个端口，群集内有负载策略会分发给不同实例</p>', 2, 2, 1, 0),
(10030000000000183, '2022-09-07 02:07:16', '2022-09-27 08:41:31', 0, '', 'string', 'string', '对外只有3001一个端口，群集内有负载策略会分发给不同实例\n\n好不错', '<p>对外只有3001一个端口，群集内有负载策略会分发给不同实例</p>\n<p>好不错</p>\n', 2, 22, 1, 186),
(10030000000000311, '2022-09-07 02:07:16', '2022-09-28 07:59:36', 10030000000000009, 'go', 'golang', 'golang', '', '', 4, 6, 1, 0),
(10030000000000312, '2022-09-07 02:12:06', '2022-09-15 08:56:04', 0, '', '算法', '算法', '', '', 1, 0, 1, 0),
(10030000000000314, '2022-09-07 02:14:08', '2022-09-26 06:54:36', 0, '', 'java', 'java', '', '', 1, 6, 1, 0),
(10030000000000316, '2022-09-07 02:14:28', '2022-09-28 08:20:13', 10030000000000009, 'go', 'golang2', 'golang2', 'golang2', 'golang2', 1, 2, 1, 28),
(10030000000000324, '2022-09-07 07:14:19', '2022-09-20 03:27:14', 0, '', 'python', 'python', '', '', 0, 2, 1, 41),
(10030000000000325, '2022-09-07 07:17:52', '2022-09-27 08:41:31', 0, '', 'spring boot', 'spring boot', '', '', 2, 2, 1, 43),
(10030000000000364, '2022-09-08 01:45:26', '2022-09-27 08:41:31', 0, '', 'javascript', 'javascript', 'JavaScript (JS) is a lightweight, interpreted, or just-in-time compiled programming language with first-class functions. While it is most well-known as the scripting language for Web pages, many non-browser environments also use it, such as Node.js, Apache CouchDB and Adobe Acrobat.\n\nJavaScript is a prototype-based, multi-paradigm, single-threaded, dynamic language, supporting object-oriented, imperative, and declarative (e.g. functional programming) styles. Read more about JavaScript. This section is dedicated to the JavaScript language itself, and not the parts that are specific to Web pages or other host environments. For information about APIs that are specific to Web pages, please see Web APIs and DOM.', '<p>JavaScript (JS) is a lightweight, interpreted, or just-in-time compiled programming language with first-class functions. While it is most well-known as the scripting language for Web pages, many non-browser environments also use it, such as Node.js, Apache CouchDB and Adobe Acrobat.</p>\n<p>JavaScript is a prototype-based, multi-paradigm, single-threaded, dynamic language, supporting object-oriented, imperative, and declarative (e.g. functional programming) styles. Read more about JavaScript. This section is dedicated to the JavaScript language itself, and not the parts that are specific to Web pages or other host environments. For information about APIs that are specific to Web pages, please see Web APIs and DOM.</p>\n', 3, 3, 1, 292),
(10030000000000401, '2022-09-08 07:26:38', '2022-09-28 08:20:13', 0, '', 'test', 'test', '111111\ngjsdghjfghjgjfhgasjdhgfhjasg[hello](https://www.baidu.com)\n\n22222222222\n\n3333333\n\n', '<p>111111\ngjsdghjfghjgjfhgasjdhgfhjasg<a href=\"https://www.baidu.com\">hello</a></p>\n<p>22222222222</p>\n<p>3333333</p>\n', 1, 3, 1, 303),
(10030000000000408, '2022-09-08 08:38:03', '2022-09-08 08:38:03', 0, '', 'code ', 'code ', 'code is code', '<p>code is code</p>\n', 0, 0, 1, 70),
(10030000000000428, '2022-09-08 11:44:10', '2022-09-26 06:54:31', 0, '', 'abc', 'abc', 'JavaScript 是一门弱类型的动态脚本语言，支持多种编程范式，包括面向对象和函数式编程，被广泛用于 Web 开发。\n\n一般来说，完整的JavaScript包括以下几个部分：\n- ECMAScript，描述了该语言的语法和基本对象\n- 文档对象模型（DOM），描述处理网页内容的方法和接口\n- 浏览器对象模型（BOM），描述与浏览器进行交互的方法和接口\n\n它的基本特点如下：\n- 是一种解释性脚本语言（代码不进行预编译）。\n- 主要用来向HTML页面添加交互行为。\n- 可以直接嵌入HTML页面，但写成单独的js文件有利于结构和行为的分离。\n\nJavaScript常用来完成以下任务：\n- 嵌入动态文本于HTML页面\n- 对浏览器事件作出响应\n- 读写HTML元素\n- 在数据被提交到服务器之前验证数据\n- 检测访客的浏览器信息\n\n![《 Javascript 优点在整个语言中占多大比例？][1]\n\n [1]: http://segmentfault.com/img/bVFXU', '<p>JavaScript 是一门弱类型的动态脚本语言，支持多种编程范式，包括面向对象和函数式编程，被广泛用于 Web 开发。</p>\n<p>一般来说，完整的JavaScript包括以下几个部分：</p>\n<ul>\n<li>ECMAScript，描述了该语言的语法和基本对象</li>\n<li>文档对象模型（DOM），描述处理网页内容的方法和接口</li>\n<li>浏览器对象模型（BOM），描述与浏览器进行交互的方法和接口</li>\n</ul>\n<p>它的基本特点如下：</p>\n<ul>\n<li>是一种解释性脚本语言（代码不进行预编译）。</li>\n<li>主要用来向HTML页面添加交互行为。</li>\n<li>可以直接嵌入HTML页面，但写成单独的js文件有利于结构和行为的分离。</li>\n</ul>\n<p>JavaScript常用来完成以下任务：</p>\n<ul>\n<li>嵌入动态文本于HTML页面</li>\n<li>对浏览器事件作出响应</li>\n<li>读写HTML元素</li>\n<li>在数据被提交到服务器之前验证数据</li>\n<li>检测访客的浏览器信息</li>\n</ul>\n<p><img src=\"http://segmentfault.com/img/bVFXU\" alt=\"《 Javascript 优点在整个语言中占多大比例？\"></p>\n', 0, 0, 1, 113),
(10030000000000507, '2022-09-16 06:32:36', '2022-09-16 08:14:27', 0, '', 'dddddd', 'dddddd', 'ddddddddddddddddd', 'ddddddddddddddddddd', 0, 1, 1, 144),
(10030000000000510, '2022-09-16 07:04:34', '2022-09-16 07:30:43', 0, '', 'vim', 'vim', 'vim', '<p>vim</p>\n', 0, 0, 1, 147),
(10030000000000592, '2022-09-20 10:23:28', '2022-09-27 08:41:31', 0, '', 'android-studio', 'Android-Studio', '', '', 3, 2, 1, 170),
(10030000000000644, '2022-09-23 09:25:39', '2022-09-27 08:41:31', 0, '', 'great', 'great', 'great', '<p>great</p>\n', 2, 1, 1, 272),
(10030000000000687, '2022-09-26 10:48:07', '2022-09-28 01:39:19', 0, '', 'qwer', 'qwerDis', 'qwert', '<p>qwert</p>\n', 0, 0, 1, 316),
(10030000000000688, '2022-09-26 10:49:39', '2022-09-28 08:20:13', 0, '', 'qwers', 'qwerDiss', 'qwerts', 'qwerts', 0, 1, 1, 300),
(10030000000000689, '2022-09-26 10:50:34', '2022-09-28 08:20:13', 10030000000000364, 'javascript', 'jss', '', '', '', 0, 0, 1, 302),
(10030000000000708, '2022-09-27 02:12:03', '2022-09-28 05:00:51', 0, '', 'flutter', 'Flutter', '', '<p>212121212121</p>\n', 3, 1, 1, 305),
(10030000000000712, '2022-09-27 02:32:53', '2022-09-27 10:09:37', 0, '', 'nextjs', 'Nextjs', '', '<p>2121212121</p>\n', 1, 1, 1, 307),
(10030000000000713, '2022-09-27 02:36:29', '2022-09-27 08:41:31', 0, '', 'umijs', 'Umijs', '212121212112121211', '<p>212121212112121211</p>\n', 1, 1, 1, 309),
(10030000000000739, '2022-09-28 01:45:09', '2022-09-28 01:45:18', 0, '', '', '', '', '', 0, 0, 1, 320);

-- --------------------------------------------------------

--
-- 表的结构 `tag_rel`
--

CREATE TABLE `tag_rel` (
  `id` bigint(20) NOT NULL COMMENT 'tag_list_id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `object_id` bigint(20) NOT NULL COMMENT 'object_id',
  `tag_id` bigint(20) NOT NULL COMMENT 'tag_id',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT 'tag_list_status(available: 1; deleted: 10)'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='tag relation';

-- --------------------------------------------------------

--
-- 表的结构 `uniqid`
--

CREATE TABLE `uniqid` (
  `id` bigint(20) NOT NULL COMMENT 'uniqid_id',
  `uniqid_type` int(11) NOT NULL DEFAULT '0' COMMENT 'uniqid_type'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='uniqid';

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `id` bigint(20) NOT NULL COMMENT 'user id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `suspended_at` timestamp NULL DEFAULT NULL COMMENT 'suspended time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'delete time',
  `last_login_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'last login date',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT 'username',
  `pass` varchar(255) NOT NULL DEFAULT '' COMMENT 'password',
  `e_mail` varchar(100) NOT NULL COMMENT 'email',
  `mail_status` tinyint(4) NOT NULL DEFAULT '2' COMMENT 'mail status(1 pass 2 to be verified)',
  `notice_status` int(11) NOT NULL DEFAULT '2' COMMENT 'notice status(1 on 2off)',
  `follow_count` int(11) NOT NULL DEFAULT '0' COMMENT 'follow count',
  `answer_count` int(11) NOT NULL DEFAULT '0' COMMENT 'answer_count',
  `question_count` int(11) NOT NULL DEFAULT '0' COMMENT 'question_count',
  `rank` int(11) NOT NULL DEFAULT '0' COMMENT 'rank',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT 'user status(available: 0; deleted: 10)',
  `authority_group` int(11) NOT NULL DEFAULT '1' COMMENT 'authority group',
  `display_name` varchar(50) NOT NULL DEFAULT '' COMMENT 'display name',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT 'avatar',
  `mobile` varchar(20) NOT NULL COMMENT 'mobile',
  `bio` text NOT NULL COMMENT 'bio markdown',
  `bio_html` text NOT NULL COMMENT 'bio html',
  `website` varchar(255) NOT NULL DEFAULT '' COMMENT 'website',
  `location` varchar(100) NOT NULL DEFAULT '' COMMENT 'location',
  `ip_info` varchar(255) NOT NULL DEFAULT '' COMMENT 'ip info',
  `is_admin` int(11) NOT NULL COMMENT 'admin flag'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user';

-- --------------------------------------------------------

--
-- 表的结构 `user_group`
--

CREATE TABLE `user_group` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT 'user group id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user group';

--
-- 转储表的索引
--

--
-- 表的索引 `activity`
--
ALTER TABLE `activity`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `post_id` (`object_id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `trigger_user_id` (`trigger_user_id`);

--
-- 表的索引 `answer`
--
ALTER TABLE `answer`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `collection`
--
ALTER TABLE `collection`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `collection_group`
--
ALTER TABLE `collection_group`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `comment`
--
ALTER TABLE `comment`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `config`
--
ALTER TABLE `config`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `key` (`key`);

--
-- 表的索引 `meta`
--
ALTER TABLE `meta`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `notification`
--
ALTER TABLE `notification`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD KEY `idx_objectid` (`object_id`);

--
-- 表的索引 `notification_read`
--
ALTER TABLE `notification_read`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `question`
--
ALTER TABLE `question`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `report`
--
ALTER TABLE `report`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `revision`
--
ALTER TABLE `revision`
  ADD PRIMARY KEY (`id`) USING BTREE;

--
-- 表的索引 `site_info`
--
ALTER TABLE `site_info`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `tag`
--
ALTER TABLE `tag`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD UNIQUE KEY `slug_name` (`slug_name`);

--
-- 表的索引 `tag_rel`
--
ALTER TABLE `tag_rel`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD UNIQUE KEY `idx_obj_tag_id` (`object_id`,`tag_id`) USING BTREE,
  ADD KEY `idx_questionid` (`object_id`),
  ADD KEY `idx_tagid` (`tag_id`) USING BTREE;

--
-- 表的索引 `uniqid`
--
ALTER TABLE `uniqid`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`) USING BTREE,
  ADD UNIQUE KEY `username` (`username`);

--
-- 表的索引 `user_group`
--
ALTER TABLE `user_group`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `activity`
--
ALTER TABLE `activity`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Activity ID autoincrement';

--
-- 使用表AUTO_INCREMENT `answer`
--
ALTER TABLE `answer`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'answer id';

--
-- 使用表AUTO_INCREMENT `collection_group`
--
ALTER TABLE `collection_group`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `comment`
--
ALTER TABLE `comment`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'comment id';

--
-- 使用表AUTO_INCREMENT `config`
--
ALTER TABLE `config`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'config id';

--
-- 使用表AUTO_INCREMENT `meta`
--
ALTER TABLE `meta`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id';

--
-- 使用表AUTO_INCREMENT `notification`
--
ALTER TABLE `notification`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'notification id';

--
-- 使用表AUTO_INCREMENT `notification_read`
--
ALTER TABLE `notification_read`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id';

--
-- 使用表AUTO_INCREMENT `report`
--
ALTER TABLE `report`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id';

--
-- 使用表AUTO_INCREMENT `revision`
--
ALTER TABLE `revision`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id';

--
-- 使用表AUTO_INCREMENT `site_info`
--
ALTER TABLE `site_info`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `tag_rel`
--
ALTER TABLE `tag_rel`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'tag_list_id';

--
-- 使用表AUTO_INCREMENT `uniqid`
--
ALTER TABLE `uniqid`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'uniqid_id';

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'user id';

--
-- 使用表AUTO_INCREMENT `user_group`
--
ALTER TABLE `user_group`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'user group id';
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

CREATE TABLE `activity` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'Activity ID autoincrement',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `user_id` bigint(20) NOT NULL COMMENT 'the user ID that generated the activity or affected by the activity',
  `trigger_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'trigger this activity user id',
  `object_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'the object ID that affected by the activity',
  `activity_type` int(11) NOT NULL COMMENT 'activity type, correspond to config id',
  `cancelled` tinyint(4) NOT NULL DEFAULT 0 COMMENT 'mark this activity if cancelled or not,default 0(not cancelled)',
  `rank` int(11) NOT NULL DEFAULT 0 COMMENT 'rank of current operating user affected',
  `has_rank` tinyint(4) NOT NULL DEFAULT 0 COMMENT 'this activity has rank or not',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `post_id` (`object_id`),
  KEY `user_id` (`user_id`),
  KEY `trigger_user_id` (`trigger_user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='activity';

CREATE TABLE `answer` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'answer id',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'create time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'update time',
  `question_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'question id',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'answer user id',
  `original_text` mediumtext NOT NULL COMMENT 'original text',
  `parsed_text` mediumtext NOT NULL COMMENT 'parsed text',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT 'answer status(available: 1; deleted: 10)',
  `adopted` int(11) NOT NULL DEFAULT 1 COMMENT 'adopted (1 failed 2 adopted)',
  `comment_count` int(11) NOT NULL DEFAULT 0 COMMENT 'comment count',
  `vote_count` int(11) NOT NULL DEFAULT 0 COMMENT 'vote count',
  `revision_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'revision id',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='answer';

CREATE TABLE `collection` (
  `id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'collection id',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'create time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT 'update time',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'user id',
  `object_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'object id',
  `user_collection_group_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'user collection group id',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='collection';

CREATE TABLE `collection_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'create time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT 'update time',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'user id',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'the collection group name',
  `default_group` int(11) NOT NULL DEFAULT 1 COMMENT 'mark this group is default, default 1',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='collection group';

CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'comment id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'user id',
  `reply_user_id` bigint(20) DEFAULT NULL COMMENT 'reply user id',
  `reply_comment_id` bigint(20) DEFAULT NULL COMMENT 'reply comment id',
  `object_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'object id',
  `question_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'question id',
  `vote_count` int(11) NOT NULL DEFAULT 0 COMMENT 'user vote amount',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT 'comment status(available: 0; deleted: 10)',
  `original_text` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'original comment content',
  `parsed_text` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'parsed comment content',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `object_id` (`object_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='comment';

CREATE TABLE `config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'config id',
  `key` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'the config key',
  `value` text COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'the config value, custom data structures and types',
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`key`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='config';

CREATE TABLE `meta` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'created time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT 'updated time',
  `object_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'object id',
  `key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'key',
  `value` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'value',
  PRIMARY KEY (`id`),
  KEY `object_id` (`object_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='meta';

CREATE TABLE `notification` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'notification id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'user_id',
  `object_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'object id',
  `content` text NOT NULL COMMENT 'notification content',
  `type` int(11) NOT NULL DEFAULT 0 COMMENT '1 inbox 2 achievements',
  `is_read` int(11) NOT NULL DEFAULT 1 COMMENT 'read status(unread: 1; read 2)',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT 'notification status(normal: 1; delete 2)',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_objectid` (`object_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='notification';

CREATE TABLE `notification_read` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'user id',
  `message_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'message id',
  `is_read` int(11) NOT NULL DEFAULT 1 COMMENT 'read status(unread: 1; read 2)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='notification read record';

CREATE TABLE `question` (
  `id` bigint(20) NOT NULL COMMENT 'question id',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'create time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'update time',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'user id',
  `title` varchar(150) NOT NULL DEFAULT '' COMMENT 'question title',
  `original_text` mediumtext NOT NULL COMMENT 'original text',
  `parsed_text` mediumtext NOT NULL COMMENT 'parsed text',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT ' question status(available: 1; deleted: 10)',
  `view_count` int(11) NOT NULL DEFAULT 0 COMMENT 'view count',
  `unique_view_count` int(11) NOT NULL DEFAULT 0 COMMENT 'unique view count',
  `vote_count` int(11) NOT NULL DEFAULT 0 COMMENT 'vote count',
  `answer_count` int(11) NOT NULL DEFAULT 0 COMMENT 'answer count',
  `collection_count` int(11) NOT NULL DEFAULT 0 COMMENT 'collection count',
  `follow_count` int(11) NOT NULL DEFAULT 0 COMMENT 'follow count',
  `accepted_answer_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'accepted answer id',
  `last_answer_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'last answer id',
  `post_update_time` timestamp NULL DEFAULT current_timestamp() COMMENT 'answer the last update time',
  `revision_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'revision id',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='question';

CREATE TABLE `report` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `user_id` bigint(20) NOT NULL COMMENT 'reporter user id',
  `object_id` bigint(20) NOT NULL COMMENT 'object id',
  `reported_user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'reported user id',
  `object_type` int(11) NOT NULL DEFAULT 0 COMMENT 'revision type',
  `report_type` int(11) NOT NULL DEFAULT 0 COMMENT 'report type',
  `content` text NOT NULL COMMENT 'report content',
  `flaged_type` int(11) NOT NULL DEFAULT 0 COMMENT 'flag type',
  `flaged_content` text DEFAULT NULL COMMENT 'flag content',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT 'status(normal: 1; pending:2; delete: 10)',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='report';

CREATE TABLE `revision` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'user id',
  `object_type` int(11) NOT NULL DEFAULT 0 COMMENT 'revision type(question: 1; answer 2)',
  `object_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'object id',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT 'title',
  `content` text NOT NULL COMMENT 'content',
  `log` varchar(255) DEFAULT NULL COMMENT 'log',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT 'revision status(normal: 1; delete 2)',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `object_id` (`object_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='revision';

CREATE TABLE `site_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'create time',
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'update time',
  `type` varchar(64) DEFAULT NULL COMMENT 'type',
  `content` mediumtext DEFAULT NULL COMMENT 'content',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT 'status',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='site info';

CREATE TABLE `tag` (
  `id` bigint(20) NOT NULL COMMENT 'tag_id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `main_tag_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'main tag id',
  `main_tag_slug_name` varchar(35) NOT NULL DEFAULT '' COMMENT 'main tag slug name',
  `slug_name` varchar(35) NOT NULL DEFAULT '' COMMENT 'slug name',
  `display_name` varchar(35) NOT NULL DEFAULT '' COMMENT 'display name',
  `original_text` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'original comment content',
  `parsed_text` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'parsed comment content',
  `follow_count` int(11) NOT NULL DEFAULT 0 COMMENT 'follow count',
  `question_count` int(11) NOT NULL COMMENT 'question count',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT 'tag status(available: 1; deleted: 10)',
  `revision_id` bigint(20) NOT NULL DEFAULT 0 COMMENT 'revision id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `slug_name` (`slug_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='tag';

CREATE TABLE `tag_rel` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'tag_list_id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `object_id` bigint(20) NOT NULL COMMENT 'object_id',
  `tag_id` bigint(20) NOT NULL COMMENT 'tag_id',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT 'tag_list_status(available: 1; deleted: 10)',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_obj_tag_id` (`object_id`,`tag_id`) USING BTREE,
  KEY `idx_questionid` (`object_id`),
  KEY `idx_tagid` (`tag_id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='tag relation';

CREATE TABLE `uniqid` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'uniqid_id',
  `uniqid_type` int(11) NOT NULL DEFAULT 0 COMMENT 'uniqid_type',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='uniqid';

CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'user id',
  `created_at` timestamp NULL DEFAULT NULL COMMENT 'create time',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT 'update time',
  `suspended_at` timestamp NULL DEFAULT NULL COMMENT 'suspended time',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'delete time',
  `last_login_date` timestamp NOT NULL DEFAULT current_timestamp() COMMENT 'last login date',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT 'username',
  `pass` varchar(255) NOT NULL DEFAULT '' COMMENT 'password',
  `e_mail` varchar(100) NOT NULL COMMENT 'email',
  `mail_status` tinyint(4) NOT NULL DEFAULT 2 COMMENT 'mail status(1 pass 2 to be verified)',
  `notice_status` int(11) NOT NULL DEFAULT 2 COMMENT 'notice status(1 on 2off)',
  `follow_count` int(11) NOT NULL DEFAULT 0 COMMENT 'follow count',
  `answer_count` int(11) NOT NULL DEFAULT 0 COMMENT 'answer_count',
  `question_count` int(11) NOT NULL DEFAULT 0 COMMENT 'question_count',
  `rank` int(11) NOT NULL DEFAULT 0 COMMENT 'rank',
  `status` int(11) NOT NULL DEFAULT 1 COMMENT 'user status(available: 0; deleted: 10)',
  `authority_group` int(11) NOT NULL DEFAULT 1 COMMENT 'authority group',
  `display_name` varchar(30) NOT NULL DEFAULT '' COMMENT 'display name',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT 'avatar',
  `mobile` varchar(20) NOT NULL COMMENT 'mobile',
  `bio` text NOT NULL COMMENT 'bio markdown',
  `bio_html` text NOT NULL COMMENT 'bio html',
  `website` varchar(255) NOT NULL DEFAULT '' COMMENT 'website',
  `location` varchar(100) NOT NULL DEFAULT '' COMMENT 'location',
  `ip_info` varchar(255) NOT NULL DEFAULT '' COMMENT 'ip info',
  `is_admin` int(11) NOT NULL COMMENT 'admin flag',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='user';

INSERT INTO `config` (`id`, `key`, `value`) VALUES
(1, 'answer.accepted', '15'),
(2, 'answer.voted_up', '10'),
(3, 'question.voted_up', '10'),
(4, 'tag.edit_accepted', '2'),
(5, 'answer.accept', '2'),
(6, 'answer.voted_down_cancel', '2'),
(7, 'question.voted_down_cancel', '2'),
(8, 'answer.vote_down_cancel', '1'),
(9, 'question.vote_down_cancel', '1'),
(10, 'user.activated', '1'),
(11, 'edit.accepted', '2'),
(12, 'answer.vote_down', '-1'),
(13, 'question.voted_down', '-2'),
(14, 'answer.voted_down', '-2'),
(15, 'answer.accept_cancel', '-2'),
(16, 'answer.deleted', '-5'),
(17, 'question.voted_up_cancel', '-10'),
(18, 'answer.voted_up_cancel', '-10'),
(19, 'answer.accepted_cancel', '-15'),
(20, 'object.reported', '-100'),
(21, 'edit.rejected', '-2'),
(22, 'daily_rank_limit', '200'),
(23, 'daily_rank_limit.exclude', '[\"answer.accepted\"]'),
(24, 'user.follow', '0'),
(25, 'comment.vote_up', '0'),
(26, 'comment.vote_up_cancel', '0'),
(27, 'question.vote_down', '0'),
(28, 'question.vote_up', '0'),
(29, 'question.vote_up_cancel', '0'),
(30, 'answer.vote_up', '0'),
(31, 'answer.vote_up_cancel', '0'),
(32, 'question.follow', '0'),
(33, 'email.config', '{\"email_web_name\":\"answer\",\"email_from\":\"\",\"email_from_pass\":\"\",\"email_from_hostname\":\"\",\"email_from_smtp\":\"\",\"email_from_name\":\"Answer Team\",\"email_register_title\":\"[{{.SiteName}}] Confirm your new account\",\"email_register_body\":\"Welcome to {{.SiteName}}<br><br>\\n\\nClick the following link to confirm and activate your new account:<br>\\n{{.RegisterUrl}}<br><br>\\n\\nIf the above link is not clickable, try copying and pasting it into the address bar of your web browser.\\n\",\"email_pass_reset_title\":\"[{{.SiteName }}] Password reset\",\"email_pass_reset_body\":\"Somebody asked to reset your password on [{{.SiteName}}].<br><br>\\n\\nIf it was not you, you can safely ignore this email.<br><br>\\n\\nClick the following link to choose a new password:<br>\\n{{.PassResetUrl}}\\n\",\"email_change_title\":\"[{{.SiteName}}] Confirm your new email address\",\"email_change_body\":\"Confirm your new email address for {{.SiteName}}  by clicking on the following link:<br><br>\\n\\n{{.ChangeEmailUrl}}<br><br>\\n\\nIf you did not request this change, please ignore this email.\\n\"}'),
(35, 'tag.follow', '0'),
(36, 'rank.question.add', '0'),
(37, 'rank.question.edit', '0'),
(38, 'rank.question.delete', '0'),
(39, 'rank.question.vote_up', '0'),
(40, 'rank.question.vote_down', '0'),
(41, 'rank.answer.add', '0'),
(42, 'rank.answer.edit', '0'),
(43, 'rank.answer.delete', '0'),
(44, 'rank.answer.accept', '0'),
(45, 'rank.answer.vote_up', '0'),
(46, 'rank.answer.vote_down', '0'),
(47, 'rank.comment.add', '0'),
(48, 'rank.comment.edit', '0'),
(49, 'rank.comment.delete', '0'),
(50, 'rank.report.add', '0'),
(51, 'rank.tag.add', '0'),
(52, 'rank.tag.edit', '0'),
(53, 'rank.tag.delete', '0'),
(54, 'rank.tag.synonym', '0'),
(55, 'rank.link.url_limit', '0'),
(56, 'rank.vote.detail', '0'),
(57, 'reason.spam', '{\"name\":\"spam\",\"description\":\"This post is an advertisement, or vandalism. It is not useful or relevant to the current topic.\"}'),
(58, 'reason.rude_or_abusive', '{\"name\":\"rude or abusive\",\"description\":\"A reasonable person would find this content inappropriate for respectful discourse.\"}'),
(59, 'reason.something', '{\"name\":\"something else\",\"description\":\"This post requires staff attention for another reason not listed above.\",\"content_type\":\"textarea\"}'),
(60, 'reason.a_duplicate', '{\"name\":\"a duplicate\",\"description\":\"This question has been asked before and already has an answer.\",\"content_type\":\"text\"}'),
(61, 'reason.not_a_answer', '{\"name\":\"not a answer\",\"description\":\"This was posted as an answer, but it does not attempt to answer the question. It should possibly be an edit, a comment, another question, or deleted altogether.\",\"content_type\":\"\"}'),
(62, 'reason.no_longer_needed', '{\"name\":\"no longer needed\",\"description\":\"This comment is outdated, conversational or not relevant to this post.\"}'),
(63, 'reason.community_specific', '{\"name\":\"a community-specific reason\",\"description\":\"This question doesn’t meet a community guideline.\"}'),
(64, 'reason.not_clarity', '{\"name\":\"needs details or clarity\",\"description\":\"This question currently includes multiple questions in one. It should focus on one problem only.\",\"content_type\":\"text\"}'),
(65, 'reason.normal', '{\"name\":\"normal\",\"description\":\"A normal post available to everyone.\"}'),
(66, 'reason.normal.user', '{\"name\":\"normal\",\"description\":\"A normal user can ask and answer questions.\"}'),
(67, 'reason.closed', '{\"name\":\"closed\",\"description\":\"A closed question can’t answer, but still can edit, vote and comment.\"}'),
(68, 'reason.deleted', '{\"name\":\"deleted\",\"description\":\"All reputation gained and lost will be restored.\"}'),
(69, 'reason.deleted.user', '{\"name\":\"deleted\",\"description\":\"Delete profile, authentication associations.\"}'),
(70, 'reason.suspended', '{\"name\":\"suspended\",\"description\":\"A suspended user can\'t log in.\"}'),
(71, 'reason.inactive', '{\"name\":\"inactive\",\"description\":\"An inactive user must re-validate their email.\"}'),
(72, 'reason.looks_ok', '{\"name\":\"looks ok\",\"description\":\"This post is good as-is and not low quality.\"}'),
(73, 'reason.needs_edit', '{\"name\":\"needs edit, and I did it\",\"description\":\"Improve and correct problems with this post yourself.\"}'),
(74, 'reason.needs_close', '{\"name\":\"needs close\",\"description\":\"A closed question can’t answer, but still can edit, vote and comment.\"}'),
(75, 'reason.needs_delete', '{\"name\":\"needs delete\",\"description\":\"All reputation gained and lost will be restored.\"}'),
(76, 'question.flag.reasons', '[\"reason.spam\",\"reason.rude_or_abusive\",\"reason.something\",\"reason.a_duplicate\"]'),
(77, 'answer.flag.reasons', '[\"reason.spam\",\"reason.rude_or_abusive\",\"reason.something\",\"reason.not_a_answer\"]'),
(78, 'comment.flag.reasons', '[\"reason.spam\",\"reason.rude_or_abusive\",\"reason.something\",\"reason.no_longer_needed\"]'),
(79, 'question.close.reasons', '[\"reason.a_duplicate\",\"reason.community_specific\",\"reason.not_clarity\",\"reason.something\"]'),
(80, 'question.status.reasons', '[\"reason.normal\",\"reason.closed\",\"reason.deleted\"]'),
(81, 'answer.status.reasons', '[\"reason.normal\",\"reason.deleted\"]'),
(82, 'comment.status.reasons', '[\"reason.normal\",\"reason.deleted\"]'),
(83, 'user.status.reasons', '[\"reason.normal.user\",\"reason.suspended\",\"reason.deleted.user\",\"reason.inactive\"]'),
(84, 'question.review.reasons', '[\"reason.looks_ok\",\"reason.needs_edit\",\"reason.needs_close\",\"needs_delete\"]'),
(85, 'answer.review.reasons', '[\"reason.looks_ok\",\"reason.needs_edit\",\"reason.needs_delete\"]'),
(86, 'comment.review.reasons', '[\"reason.looks_ok\",\"reason.needs_edit\",\"reason.needs_delete\"]');

INSERT INTO `site_info` (`id`, `created_at`, `updated_at`, `type`, `content`, `status`) VALUES
(1, '2022-09-16 12:08:27', '2022-09-29 12:34:59', 'interface', '{\"logo\":\"\",\"theme\":\"black\",\"language\":\"en_US\"}', 0);

INSERT INTO `user` (`id`, `created_at`, `updated_at`, `suspended_at`, `deleted_at`, `last_login_date`, `username`, `pass`, `e_mail`, `mail_status`, `notice_status`, `follow_count`, `answer_count`, `question_count`, `rank`, `status`, `authority_group`, `display_name`, `avatar`, `mobile`, `bio`, `bio_html`, `website`, `location`, `ip_info`, `is_admin`) VALUES
(1, '2022-09-01 20:58:47', '2022-09-28 13:14:48', NULL, NULL, '2022-09-28 13:14:48', 'admin', '$2a$10$.gnUnpW.8ssRNaEvx.XwvOR2NuPsGzFLWWX2rqSIVAdIvLNZZYs5y', 'admin@admin.com', 1, 1, 0, 0, 0, 0, 1, 0, 'admin', '', '', '', '', '', '', '', 1);


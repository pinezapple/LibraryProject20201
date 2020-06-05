drop table if exists borrowform_cache;

create table borrowform_cache 
(
	`id_borrow` bigint(20) unsigned NOT NULL,
	`id_doc` bigint(20) unsigned NOT NULL,
	`doc_name` varchar(100) DEFAULT "",
	`id_cus` bigint(20) unsigned NOT NULL,
	`id_lib` bigint(20) unsigned NOT NULL,
	`status` int NOT NULL DEFAULT 0,
	`start_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`end_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`id_borrow`)
) Engine=InnoDB;

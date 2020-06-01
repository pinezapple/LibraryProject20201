drop table if exists doc;

create table doc 
(
	`id_doc` bigint(20) unsigned NOT NULL,
	`doc_name` varchar(100) DEFAULT "",
	`doc_author` varchar(30) DEFAULT "",
	`doc_type` varchar(30) DEFAULT "",
	`doc_description` varchar(100) DEFAULT "", 
	`status` int DEFAULT 0,
	`id_borrow` bigint(20) unsigned DEFAULT 0,
	`fee` bigint(20) DEFAULT 0,
	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`id_doc`)
) Engine=InnoDB;

drop table if exists documents;
drop table if exists document_version;
drop table if exists categories;
drop table if exists authors;
drop table if exists black_list;
drop table if exists barcode_cache;

create table document_version
(
	`document_version_id` bigint(20),
	`document_version` varchar(1000),
	`doc_id` bigint(20),
	`doc_description` varchar(1000),
	`publisher` varchar(1000),
	`author_id` bigint(20),
	`fee` bigint(20),
	`price` bigint(20),
	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

create table documents
(
	`doc_id` bigint(20),
	`doc_name` varchar(255),
	`category_id` bigint(20),
	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

create table categories
(
	`category_id` bigint(20),
	`category_name` varchar(255),
	`doc_description` varchar(1000),
        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

create table authors
(
	`author_id` bigint(20),
	`author_name` varchar(1000),
	`description` varchar(1000),
        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

create table barcode_cache
(
  `barcode_id` bigint(20),
  `document_version_id` bigint(20)
) Engine=InnoDB;

create table black_list 
(
  `user_id` bigint(20),
  `borrow_form_id` bigint(20),
  `money` varchar(255),
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;





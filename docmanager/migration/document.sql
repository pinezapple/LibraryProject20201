drop table if exists documents;
drop table if exists documents_version;
drop table if exists categories;
drop table if exists sale_bill;
drop table if exists barcodes;
drop table if exists authors;
drop table if exists borrow_form;
drop table if exists payments;

CREATE TABLE `documents` (
  `doc_id` bigint(20),
  `doc_name` bigint(20),
  `category_id` bigint(20),
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

CREATE TABLE `documents_version` (
  `documents_version` varchar(255),
  `doc_id` bigint(20),
  `version` bigint(20),
  `doc_description` varchar(255),
  `author_id` bigint(20),
  `fee` bigint(20),
  `price` bigint(20),
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

CREATE TABLE `barcodes` (
  `barcode_id` bigint(20),
  `documents_version` varchar(255),
  `status` bigint(20),
  `sale_bill_id` bigint(20), 
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

CREATE TABLE `borrow_form` (
  `borrow_form_id` bigint(20),
  `librarian_id` bigint(20),
  `reader_id` bigint(20),
  `barcode_id` BLOB,
  `status` bigint(20),
  `borrow_start_time` timestamp DEFAULT CURRENT_TIMESTAMP,
  `borrow_end_time` timestamp,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

CREATE TABLE `payments` (
  `payment_id` bigint(20),
  `borrow_form_id` bigint(20),
  `reader_id` bigint(20),
  `barcode_id` BLOB, 
  `barcode_status` BLOB,
  `money` BLOB, 
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

CREATE TABLE `categories` (
  `category_id` bigint(20),
  `category_name` varchar(255),
  `doc_description` varchar(255),
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

CREATE TABLE `authors` (
  `author_id` bigint(20),
  `author_name` varchar(255),
  `description` varchar(255),
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

CREATE TABLE `sale_bill` (
  `sale_bill_id` bigint(20),
  `barcode_id` BLOB,
  `price` BLOB,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

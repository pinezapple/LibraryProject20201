drop table if exists sale_bill;
drop table if exists barcodes;
drop table if exists borrow_form;
drop table if exists payments;

CREATE TABLE `barcodes` (
  `barcode_id` bigint(20),
  `documents_version_id` bigint(20),
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

CREATE TABLE `sale_bill` (
  `sale_bill_id` bigint(20),
  `barcode_id` BLOB,
  `price` BLOB,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) Engine=InnoDB;

/*
 Navicat Premium Data Transfer

 Source Server         : PostgresLocal
 Source Server Type    : PostgreSQL
 Source Server Version : 100014
 Source Host           : localhost:5432
 Source Catalog        : pos
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 100014
 File Encoding         : 65001

 Date: 13/07/2021 15:38:26
*/


-- ----------------------------
-- Type structure for transaction_type
-- ----------------------------
DROP TYPE IF EXISTS "public"."transaction_type";
CREATE TYPE "public"."transaction_type" AS ENUM (
  'in',
  'out'
);
ALTER TYPE "public"."transaction_type" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for product_outlets_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."product_outlets_id_seq";
CREATE SEQUENCE "public"."product_outlets_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."product_outlets_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for roles_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."roles_id_seq";
CREATE SEQUENCE "public"."roles_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."roles_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Sequence structure for units_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."units_id_seq";
CREATE SEQUENCE "public"."units_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;
ALTER SEQUENCE "public"."units_id_seq" OWNER TO "postgres";

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS "public"."categories";
CREATE TABLE "public"."categories" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."categories" OWNER TO "postgres";

-- ----------------------------
-- Records of categories
-- ----------------------------
BEGIN;
INSERT INTO "public"."categories" VALUES ('e15a9796-2b7b-4534-a73f-93bd76179f8f', 'mie instant', '2021-07-10 06:28:09', '2021-07-10 06:28:09', NULL);
INSERT INTO "public"."categories" VALUES ('d2ddbdce-06e6-47fb-b39a-84a6d54664b8', 'es krim', '2021-07-10 06:28:09', '2021-07-10 06:28:09', NULL);
INSERT INTO "public"."categories" VALUES ('6ad439ea-6e27-4cea-8d3b-44d5c1d0ca7a', 'beras', '2021-07-10 06:28:09', '2021-07-10 06:28:09', NULL);
INSERT INTO "public"."categories" VALUES ('2480c81b-fbe5-4f33-b542-838f498acca4', 'sabun cuci piring', '2021-07-10 06:28:09', '2021-07-10 06:28:09', NULL);
INSERT INTO "public"."categories" VALUES ('3dd17261-f8b5-4753-be9e-f44bf97130ba', 'sabun cuci tangan', '2021-07-10 06:28:09', '2021-07-10 06:28:09', NULL);
INSERT INTO "public"."categories" VALUES ('ced93eb2-b361-41aa-a882-d08369c14090', 'sampo', '2021-07-10 06:28:09', '2021-07-10 06:28:09', NULL);
COMMIT;

-- ----------------------------
-- Table structure for customers
-- ----------------------------
DROP TABLE IF EXISTS "public"."customers";
CREATE TABLE "public"."customers" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "merchant_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "mobile_phone_number" char(15) COLLATE "pg_catalog"."default",
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."customers" OWNER TO "postgres";

-- ----------------------------
-- Records of customers
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS "public"."files";
CREATE TABLE "public"."files" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."files" OWNER TO "postgres";

-- ----------------------------
-- Records of files
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for merchants
-- ----------------------------
DROP TABLE IF EXISTS "public"."merchants";
CREATE TABLE "public"."merchants" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "logo" char(36) COLLATE "pg_catalog"."default",
  "address" text COLLATE "pg_catalog"."default",
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."merchants" OWNER TO "postgres";

-- ----------------------------
-- Records of merchants
-- ----------------------------
BEGIN;
INSERT INTO "public"."merchants" VALUES ('7123db18-acbf-4cbc-b26e-9ebba5eef36d', 'Warung Asep', NULL, 'Jalan Buntu Nomor 10', '2021-07-10 08:07:32', '2021-07-10 08:07:32', NULL);
COMMIT;

-- ----------------------------
-- Table structure for order_details
-- ----------------------------
DROP TABLE IF EXISTS "public"."order_details";
CREATE TABLE "public"."order_details" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "order_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "product_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "price" numeric NOT NULL,
  "discount" numeric,
  "quantity" int2,
  "sub_total" numeric,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."order_details" OWNER TO "postgres";

-- ----------------------------
-- Records of order_details
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS "public"."orders";
CREATE TABLE "public"."orders" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "transaction_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "customer_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."orders" OWNER TO "postgres";

-- ----------------------------
-- Records of orders
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for outlets
-- ----------------------------
DROP TABLE IF EXISTS "public"."outlets";
CREATE TABLE "public"."outlets" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "merchant_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "address" text COLLATE "pg_catalog"."default",
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."outlets" OWNER TO "postgres";

-- ----------------------------
-- Records of outlets
-- ----------------------------
BEGIN;
INSERT INTO "public"."outlets" VALUES ('5df9fb2d-7df6-4c95-9c3b-e9b161964627', '7123db18-acbf-4cbc-b26e-9ebba5eef36d', 'outlet a', 'pinggir jalan', '2021-07-10 12:36:51', '2021-07-10 12:36:52', NULL);
INSERT INTO "public"."outlets" VALUES ('5461d5b3-109f-40d4-8660-4658837f1649', '7123db18-acbf-4cbc-b26e-9ebba5eef36d', 'outlet b', 'pinggir perempatan', '2021-07-10 12:37:10', '2021-07-10 12:37:12', NULL);
COMMIT;

-- ----------------------------
-- Table structure for product_outlets
-- ----------------------------
DROP TABLE IF EXISTS "public"."product_outlets";
CREATE TABLE "public"."product_outlets" (
  "id" int4 NOT NULL DEFAULT nextval('product_outlets_id_seq'::regclass),
  "product_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "outlet_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "price" numeric,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."product_outlets" OWNER TO "postgres";

-- ----------------------------
-- Records of product_outlets
-- ----------------------------
BEGIN;
INSERT INTO "public"."product_outlets" VALUES (13, 'c2d6dc0d-025e-4cea-93bd-6f43670fb3fe', '5df9fb2d-7df6-4c95-9c3b-e9b161964627', 3000, '0001-01-01 00:00:00', '0001-01-01 00:00:00', NULL);
INSERT INTO "public"."product_outlets" VALUES (14, 'c2d6dc0d-025e-4cea-93bd-6f43670fb3fe', '5461d5b3-109f-40d4-8660-4658837f1649', 3000, '0001-01-01 00:00:00', '0001-01-01 00:00:00', NULL);
COMMIT;

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS "public"."products";
CREATE TABLE "public"."products" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "sku" char(10) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "unit_id" int4 NOT NULL,
  "category_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "image_id" char(36) COLLATE "pg_catalog"."default",
  "merchant_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6),
  "stock" int4
)
;
ALTER TABLE "public"."products" OWNER TO "postgres";

-- ----------------------------
-- Records of products
-- ----------------------------
BEGIN;
INSERT INTO "public"."products" VALUES ('c2d6dc0d-025e-4cea-93bd-6f43670fb3fe', 'm-123     ', 'mie indome goreng', 1, 'e15a9796-2b7b-4534-a73f-93bd76179f8f', '                                    ', '7123db18-acbf-4cbc-b26e-9ebba5eef36d', '2021-07-10 06:22:02.24336', '2021-07-10 06:30:51.172518', NULL, 0);
INSERT INTO "public"."products" VALUES ('f1d39b95-c8cb-4098-a5f1-f0c417bdd7ad', 'm-124     ', 'mie indome ayam geprek', 1, 'e15a9796-2b7b-4534-a73f-93bd76179f8f', '                                    ', '7123db18-acbf-4cbc-b26e-9ebba5eef36d', '2021-07-10 06:31:22.772122', '2021-07-10 06:31:46.119718', '2021-07-10 06:31:46.119718', 0);
COMMIT;

-- ----------------------------
-- Table structure for purchase_details
-- ----------------------------
DROP TABLE IF EXISTS "public"."purchase_details";
CREATE TABLE "public"."purchase_details" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "purchase_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "item_name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "price" numeric NOT NULL,
  "discount" numeric,
  "quantity" int2,
  "sub_total" numeric,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."purchase_details" OWNER TO "postgres";

-- ----------------------------
-- Records of purchase_details
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for purchases
-- ----------------------------
DROP TABLE IF EXISTS "public"."purchases";
CREATE TABLE "public"."purchases" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "transaction_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "supplier_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."purchases" OWNER TO "postgres";

-- ----------------------------
-- Records of purchases
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS "public"."roles";
CREATE TABLE "public"."roles" (
  "id" int4 NOT NULL DEFAULT nextval('roles_id_seq'::regclass),
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "slug" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."roles" OWNER TO "postgres";

-- ----------------------------
-- Records of roles
-- ----------------------------
BEGIN;
INSERT INTO "public"."roles" VALUES (1, 'kasir', 'kasir');
INSERT INTO "public"."roles" VALUES (2, 'pemilik', 'pemilik');
INSERT INTO "public"."roles" VALUES (3, 'super admin', 'super-admin');
COMMIT;

-- ----------------------------
-- Table structure for suppliers
-- ----------------------------
DROP TABLE IF EXISTS "public"."suppliers";
CREATE TABLE "public"."suppliers" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "merchant_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "phone_number" char(15) COLLATE "pg_catalog"."default",
  "address" text COLLATE "pg_catalog"."default",
  "description" text COLLATE "pg_catalog"."default",
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."suppliers" OWNER TO "postgres";

-- ----------------------------
-- Records of suppliers
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for transactions
-- ----------------------------
DROP TABLE IF EXISTS "public"."transactions";
CREATE TABLE "public"."transactions" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "merchat_id" char(36) COLLATE "pg_catalog"."default" NOT NULL,
  "transaction_number" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "transaction_type" "public"."transaction_type",
  "sub_total" numeric,
  "discount" numeric,
  "grand_total" numeric,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."transactions" OWNER TO "postgres";

-- ----------------------------
-- Records of transactions
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for units
-- ----------------------------
DROP TABLE IF EXISTS "public"."units";
CREATE TABLE "public"."units" (
  "id" int4 NOT NULL DEFAULT nextval('units_id_seq'::regclass),
  "name" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."units" OWNER TO "postgres";

-- ----------------------------
-- Records of units
-- ----------------------------
BEGIN;
INSERT INTO "public"."units" VALUES (1, 'pcs', '2021-07-10 12:28:42', '2021-07-10 12:28:45', NULL);
INSERT INTO "public"."units" VALUES (2, 'lusin', '2021-07-10 12:29:04', '2021-07-10 12:29:06', NULL);
INSERT INTO "public"."units" VALUES (3, 'kardus', '2021-07-10 12:29:19', '2021-07-10 12:29:21', NULL);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" char(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT uuid_generate_v4(),
  "email" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "role_id" int4 NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL,
  "deleted_at" timestamp(6),
  "merchant_id" char(36) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."users" OWNER TO "postgres";

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO "public"."users" VALUES ('9c8a2bab-e97c-4347-9ba6-c7c449f2da13', 'kasir1@mailinator.com', '$2a$04$x3dIS782mCchMC97PDW0Ge57pKPHWd0lmkPSh/9ZC03JcAu.yoeYe', 1, '2021-07-10 02:09:17.554668', '2021-07-10 02:09:17.554668', NULL, '7123db18-acbf-4cbc-b26e-9ebba5eef36d');
INSERT INTO "public"."users" VALUES ('3f14ab91-a762-417f-ad1d-42b5b3196dbc', 'kasir2@mailinator.com', '$2a$04$LsRzLY2AUQAN.bFIZc4Yb.9Nf3bD9m77H2FMgdaBdRBbXjkNxceVG', 1, '2021-07-10 02:10:35.950618', '2021-07-10 03:48:00.477924', NULL, '7123db18-acbf-4cbc-b26e-9ebba5eef36d');
INSERT INTO "public"."users" VALUES ('7f6f8391-655e-4965-940d-e85a4abea736', 'kasir3@mailinator.com', '$2a$04$pQFlof.sCIeDFchgxcWboefTbLrEs0Sh4qtpxLEd2m6vmV./BOuW.', 1, '2021-07-10 03:24:24.868078', '2021-07-10 03:49:57.552118', '2021-07-10 03:49:57.552118', '7123db18-acbf-4cbc-b26e-9ebba5eef36d');
COMMIT;

-- ----------------------------
-- Function structure for uuid_generate_v1
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v1"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v1"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v1'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_generate_v1"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_generate_v1mc
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v1mc"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v1mc"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v1mc'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_generate_v1mc"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_generate_v3
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v3"("namespace" uuid, "name" text);
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v3"("namespace" uuid, "name" text)
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v3'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_generate_v3"("namespace" uuid, "name" text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_generate_v4
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v4"();
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v4"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v4'
  LANGUAGE c VOLATILE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_generate_v4"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_generate_v5
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_generate_v5"("namespace" uuid, "name" text);
CREATE OR REPLACE FUNCTION "public"."uuid_generate_v5"("namespace" uuid, "name" text)
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_generate_v5'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_generate_v5"("namespace" uuid, "name" text) OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_nil
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_nil"();
CREATE OR REPLACE FUNCTION "public"."uuid_nil"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_nil'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_nil"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_ns_dns
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_dns"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_dns"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_dns'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_ns_dns"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_ns_oid
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_oid"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_oid"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_oid'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_ns_oid"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_ns_url
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_url"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_url"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_url'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_ns_url"() OWNER TO "postgres";

-- ----------------------------
-- Function structure for uuid_ns_x500
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."uuid_ns_x500"();
CREATE OR REPLACE FUNCTION "public"."uuid_ns_x500"()
  RETURNS "pg_catalog"."uuid" AS '$libdir/uuid-ossp', 'uuid_ns_x500'
  LANGUAGE c IMMUTABLE STRICT
  COST 1;
ALTER FUNCTION "public"."uuid_ns_x500"() OWNER TO "postgres";

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."product_outlets_id_seq"
OWNED BY "public"."product_outlets"."id";
SELECT setval('"public"."product_outlets_id_seq"', 17, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."roles_id_seq"
OWNED BY "public"."roles"."id";
SELECT setval('"public"."roles_id_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."units_id_seq"
OWNED BY "public"."units"."id";
SELECT setval('"public"."units_id_seq"', 2, true);

-- ----------------------------
-- Primary Key structure for table categories
-- ----------------------------
ALTER TABLE "public"."categories" ADD CONSTRAINT "categories_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table customers
-- ----------------------------
ALTER TABLE "public"."customers" ADD CONSTRAINT "customers_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table files
-- ----------------------------
ALTER TABLE "public"."files" ADD CONSTRAINT "files_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table merchants
-- ----------------------------
ALTER TABLE "public"."merchants" ADD CONSTRAINT "merchants_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table order_details
-- ----------------------------
ALTER TABLE "public"."order_details" ADD CONSTRAINT "order_details_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table orders
-- ----------------------------
ALTER TABLE "public"."orders" ADD CONSTRAINT "orders_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table outlets
-- ----------------------------
ALTER TABLE "public"."outlets" ADD CONSTRAINT "outlets_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table product_outlets
-- ----------------------------
ALTER TABLE "public"."product_outlets" ADD CONSTRAINT "product_outlets_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table products
-- ----------------------------
ALTER TABLE "public"."products" ADD CONSTRAINT "products_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table purchase_details
-- ----------------------------
ALTER TABLE "public"."purchase_details" ADD CONSTRAINT "purchase_details_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table purchases
-- ----------------------------
ALTER TABLE "public"."purchases" ADD CONSTRAINT "purchases_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table roles
-- ----------------------------
ALTER TABLE "public"."roles" ADD CONSTRAINT "roles_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table suppliers
-- ----------------------------
ALTER TABLE "public"."suppliers" ADD CONSTRAINT "suppliers_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table transactions
-- ----------------------------
ALTER TABLE "public"."transactions" ADD CONSTRAINT "transactions_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table units
-- ----------------------------
ALTER TABLE "public"."units" ADD CONSTRAINT "units_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table customers
-- ----------------------------
ALTER TABLE "public"."customers" ADD CONSTRAINT "customers_merchant_id_fkey" FOREIGN KEY ("merchant_id") REFERENCES "public"."merchants" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table merchants
-- ----------------------------
ALTER TABLE "public"."merchants" ADD CONSTRAINT "merchants_logo_fkey" FOREIGN KEY ("logo") REFERENCES "public"."files" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table order_details
-- ----------------------------
ALTER TABLE "public"."order_details" ADD CONSTRAINT "order_details_order_id_fkey" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."order_details" ADD CONSTRAINT "order_details_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table orders
-- ----------------------------
ALTER TABLE "public"."orders" ADD CONSTRAINT "orders_customer_id_fkey" FOREIGN KEY ("customer_id") REFERENCES "public"."customers" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."orders" ADD CONSTRAINT "orders_transaction_id_fkey" FOREIGN KEY ("transaction_id") REFERENCES "public"."transactions" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table outlets
-- ----------------------------
ALTER TABLE "public"."outlets" ADD CONSTRAINT "outlets_merchant_id_fkey" FOREIGN KEY ("merchant_id") REFERENCES "public"."merchants" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table product_outlets
-- ----------------------------
ALTER TABLE "public"."product_outlets" ADD CONSTRAINT "product_outlets_outlet_id_fkey" FOREIGN KEY ("outlet_id") REFERENCES "public"."outlets" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."product_outlets" ADD CONSTRAINT "product_outlets_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "public"."products" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table products
-- ----------------------------
ALTER TABLE "public"."products" ADD CONSTRAINT "products_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."products" ADD CONSTRAINT "products_merchat_id_fkey" FOREIGN KEY ("merchant_id") REFERENCES "public"."merchants" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."products" ADD CONSTRAINT "products_unit_id_fkey" FOREIGN KEY ("unit_id") REFERENCES "public"."units" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table purchase_details
-- ----------------------------
ALTER TABLE "public"."purchase_details" ADD CONSTRAINT "purchase_details_purchase_id_fkey" FOREIGN KEY ("purchase_id") REFERENCES "public"."purchases" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table purchases
-- ----------------------------
ALTER TABLE "public"."purchases" ADD CONSTRAINT "purchases_supplier_id_fkey" FOREIGN KEY ("supplier_id") REFERENCES "public"."suppliers" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."purchases" ADD CONSTRAINT "purchases_transaction_id_fkey" FOREIGN KEY ("transaction_id") REFERENCES "public"."transactions" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table suppliers
-- ----------------------------
ALTER TABLE "public"."suppliers" ADD CONSTRAINT "suppliers_merchant_id_fkey" FOREIGN KEY ("merchant_id") REFERENCES "public"."merchants" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table transactions
-- ----------------------------
ALTER TABLE "public"."transactions" ADD CONSTRAINT "transactions_merchat_id_fkey" FOREIGN KEY ("merchat_id") REFERENCES "public"."merchants" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_role_id_fkey" FOREIGN KEY ("role_id") REFERENCES "public"."roles" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;

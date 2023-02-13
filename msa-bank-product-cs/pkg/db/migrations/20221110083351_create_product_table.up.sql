create table if not exists msa_bank_product_cs_schema.product
(
    id         uuid
        constraint table_name_pk
            primary key,
    name varchar not null,
    description  varchar
);
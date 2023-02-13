create table if not exists msa_bank_client_cs_schema.client
(
    id         uuid
        constraint table_name_pk
            primary key,
    first_name varchar not null,
    last_ha_name  varchar not null,
    birth_date varchar not null
);
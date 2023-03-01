create table if not exists msa_bank_account_cs_schema.account
(
    id             uuid
        constraint table_name_pk
            primary key,
    account_number varchar     not null,
    amount         float8      not null,
    client_id      uuid        not null,
    start_date     varchar(10) not null,
    end_date       varchar(10) not null
);

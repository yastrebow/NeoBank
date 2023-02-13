create table if not exists msa_bank_credit_cs_schema.credit
(
    id uuid constraint table_name_pk primary key,
    amount NUMERIC(10, 2) not null,
    rate  NUMERIC(3, 1) not null,
    months integer not null,
    total_amount NUMERIC(10, 2) not null
);

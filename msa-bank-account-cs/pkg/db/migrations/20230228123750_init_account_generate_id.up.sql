 ALTER TABLE msa_bank_account_cs_schema.account
 ALTER COLUMN id SET DATA TYPE UUID USING (gen_random_uuid()),
 ALTER COLUMN id SET DEFAULT gen_random_uuid();
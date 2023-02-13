 ALTER TABLE msa_bank_credit_cs_schema.credit
 ALTER COLUMN id SET DATA TYPE UUID USING (gen_random_uuid()),
 ALTER COLUMN id SET DEFAULT gen_random_uuid();
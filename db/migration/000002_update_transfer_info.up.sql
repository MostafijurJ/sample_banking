ALTER table transfers ADD COLUMN "status" varchar NOT NULL DEFAULT 'PENDING';
ALTER table transfers ADD COLUMN "from_account_number" varchar NOT NULL;
ALTER table transfers ADD COLUMN "to_account_number" varchar NOT NULL;
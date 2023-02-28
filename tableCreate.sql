CREATE TABLE customers (
	customer_number int4 PRIMARY KEY,
	customer_type varchar(256) NULL,
	customer_name varchar(256) NULL,
	customer_country_code bpchar(20) NULL,
	customer_state varchar(100) NULL,
	customer_postal_code bpchar(10) NULL,
	customer_city varchar(85) NULL,
	customer_address varchar(256) NULL,
	customer_phone varchar(16) NULL
);

CREATE TABLE accounts (
	account_number int4 PRIMARY KEY,
	account_type varchar(30) NULL,
	account_postal_code varchar(10) NULL,
	account_country_code bpchar(10) NULL,
	customer_number int4 not NULL, 
    CONSTRAINT fk_customer_number
       FOREIGN KEY(customer_number)
          REFERENCES customers(customer_number) ON DELETE CASCADE
);

CREATE TABLE cards (
	payment_card_number varchar(20) PRIMARY KEY,
	credit_limit numeric NULL,
	account_number int4,
    CONSTRAINT fk_account_number
       FOREIGN KEY(account_number)
          REFERENCES accounts(account_number) ON DELETE CASCADE
);


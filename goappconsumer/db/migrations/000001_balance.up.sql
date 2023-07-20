CREATE TABLE balance (
	id varchar(255) not null,
	account_id_from varchar(255) not null,
    account_id_to varchar(255) not null,
    balance_account_id_from decimal(19,2),
    balance_account_id_to decimal(19,2),
	created_at datetime,
	update_at datetime,
	PRIMARY KEY (id)
)
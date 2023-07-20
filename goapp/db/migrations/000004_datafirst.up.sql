BEGIN;
INSERT INTO clients (id, name, email, created_at) VALUES ('9b5da235-4b0a-42df-be2a-8bccd41d232a','John Doe','john@j.com','2023-07-20');
INSERT INTO clients (id, name, email, created_at) VALUES ('bf3cf371-4739-472a-9e9c-34bcd16703f1','Jane Doe','jane@j.com','2023-07-20');
INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('4b9633e7-21bf-413c-be9c-4cded7468652','9b5da235-4b0a-42df-be2a-8bccd41d232a',1000,'2023-07-20');
INSERT INTO accounts (id, client_id, balance, created_at) VALUES ('c990b127-9c96-4ea9-a996-c70c3d9bd44c','bf3cf371-4739-472a-9e9c-34bcd16703f1',2000,'2023-07-20');
COMMIT;
CREATE TABLE IF NOT EXISTS registers
(
    id SMALLSERIAL PRIMARY KEY, 
    first_name             VARCHAR(255) NOT NULL,
    last_name             VARCHAR(255) NOT NULL,
    address             VARCHAR(255) NOT NULL,
    phonenumber             VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    title VARCHAR NOT NULL
);

INSERT INTO employees(first_name, last_name, title) VALUES('Daniil', 'Blagiy', 'intern');
INSERT INTO employees(first_name, last_name, title) VALUES('Ivan', 'Thespacebiker', 'scientist');
INSERT INTO employees(first_name, last_name, title) VALUES('Ostap', 'Rodrigez', 'bender');

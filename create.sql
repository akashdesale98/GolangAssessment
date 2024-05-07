CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Employee (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    salary DECIMAL(10, 2),
    position VARCHAR(100)
);

INSERT INTO Employee ( name, salary, position) 
VALUES 
    ( 'John Doe', 50000, 'Manager'),
    ( 'Jane Smith', 60000, 'Developer'),
    ( 'Michael Johnson', 55000, 'Analyst'),
    ( 'Emily Brown', 70000, 'Designer');

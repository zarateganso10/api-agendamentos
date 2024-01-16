CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users_type (
    name VARCHAR(255) UNIQUE PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS companies (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    type VARCHAR(255),
    company_id uuid,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_type FOREIGN KEY(type) REFERENCES users_type(name) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT fk_company FOREIGN KEY(company_id) REFERENCES companies(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS appointments (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    client_id uuid DEFAULT NULL,
    employee_id uuid,
    start_date TIMESTAMPTZ NOT NULL,
    finish_date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_employee FOREIGN KEY(employee_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_client FOREIGN KEY(client_id) REFERENCES users(id) ON DELETE SET NULL
);








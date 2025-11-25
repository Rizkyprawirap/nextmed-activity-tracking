CREATE TABLE admins (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);

CREATE TABLE clients (
    client_id   UUID not null PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    email       VARCHAR(255) UNIQUE NOT NULL,
    api_key     VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP NULL
);


CREATE TABLE logs (
    log_id      UUID NOT null PRIMARY KEY,
    client_id   UUID NOT NULL REFERENCES clients(client_id),
    api_key     VARCHAR(255) NOT NULL,
    ip          VARCHAR(45),
    endpoint    TEXT NOT NULL,
    timestamp   TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP NULL
);



CREATE INDEX idx_logs_api_key_timestamp
    ON logs (api_key, timestamp DESC);

CREATE INDEX idx_logs_api_key
    ON logs (api_key);

CREATE INDEX idx_logs_timestamp
    ON logs (timestamp DESC);

CREATE INDEX idx_logs_client_timestamp
    ON logs (client_id, timestamp DESC);

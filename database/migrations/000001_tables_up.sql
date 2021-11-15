CREATE TYPE user_auth_type AS ENUM('local', 'vk', 'google');
CREATE TYPE stock_type AS ENUM('stock', 'bond', 'fund');

CREATE TABLE if not exists users
(
    user_id        SERIAL              NOT NULL,
    user_auth_type user_auth_type      NOT NULL,
    user_name      VARCHAR(255)        NOT NULL,
    lastname       VARCHAR(255)        NOT NULL,
    email          VARCHAR(255)        UNIQUE,
    password_hash  VARCHAR(255),
    created_at     TIMESTAMPTZ         NOT NULL DEFAULT current_timestamp,
    modified_at    TIMESTAMPTZ,
    PRIMARY KEY (user_id, user_auth_type)
);
-- CREATE TABLE if not exists vkusers
-- (
--     vkuser_id     INTEGER             PRIMARY KEY,
--     vkuser_name   VARCHAR(255)        NOT NULL,
--     lastname      VARCHAR(255)        NOT NULL,
--     email         VARCHAR(255)        UNIQUE
-- );
CREATE TABLE if not exists portfolios
(
    portfolio_id   SERIAL             PRIMARY KEY,
    user_id        INTEGER,
    user_auth_type user_auth_type,
--     vkuser_id      INTEGER            REFERENCES vkusers(vkuser_id),
    portfolio_name VARCHAR(255)       NOT NULL,
    description    VARCHAR(255),
    is_public      BOOL               NOT NULL,
    created_at     TIMESTAMPTZ        NOT NULL DEFAULT current_timestamp,
    modified_at    TIMESTAMPTZ,
    FOREIGN KEY (user_id, user_auth_type) REFERENCES users(user_id, user_auth_type)
);
CREATE TABLE if not exists currencies
(
    currency_id   SERIAL              PRIMARY KEY,
    currency_name VARCHAR(255)        NOT NULL UNIQUE,
    ticker        VARCHAR(255)        NOT NULL UNIQUE
);
CREATE TABLE if not exists balances
(
    balance_id    SERIAL              PRIMARY KEY,
    portfolio_id  INTEGER             REFERENCES portfolios(portfolio_id),
    currency_id   INTEGER             REFERENCES currencies(currency_id),
    amount        MONEY               DEFAULT 0
);
-- CREATE TABLE if not exists type_stocks
-- (
--     type_stocks_id SERIAL             PRIMARY KEY,
--     name_stock_type VARCHAR(255)      NOT NULL UNIQUE
-- );
CREATE TABLE if not exists stocks
(
    stock_id      SERIAL              PRIMARY KEY,
    stock_name    VARCHAR(255)        NOT NULL UNIQUE,
    description   VARCHAR(255),
    ticker        VARCHAR(255)        NOT NULL UNIQUE,
    stock_type    stock_type          NOT NULL,
    cost          MONEY               NOT NULL,
    currency      INTEGER             REFERENCES currencies(currency_id)
);
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
CREATE TABLE if not exists portfolios
(
    portfolio_id   SERIAL             PRIMARY KEY,
    user_id        INTEGER,
    user_auth_type user_auth_type,
    portfolio_name VARCHAR(255)       NOT NULL,
    description    VARCHAR(255),
    is_public      BOOL               NOT NULL,
    created_at     TIMESTAMPTZ        NOT NULL DEFAULT current_timestamp,
    modified_at    TIMESTAMPTZ,
    FOREIGN KEY (user_id, user_auth_type) REFERENCES users(user_id, user_auth_type) ON DELETE CASCADE
);
CREATE TABLE if not exists currencies
(
    currency_id     SERIAL              PRIMARY KEY,
    currency_name   VARCHAR(255)        NOT NULL UNIQUE,
    currency_ticker VARCHAR(255)        NOT NULL UNIQUE
);
CREATE TABLE if not exists balances
(
    balance_id    SERIAL              PRIMARY KEY,
    portfolio_id  INTEGER             REFERENCES portfolios(portfolio_id) ON DELETE CASCADE,
    currency_id   INTEGER             REFERENCES currencies(currency_id),
    money_value   FLOAT               NOT NULL DEFAULT 0

);
CREATE TABLE if not exists stocks
(
    stock_id      SERIAL,
    stock_name    VARCHAR(255)        NOT NULL UNIQUE,
    description   VARCHAR(255),
    ticker        VARCHAR(255)        NOT NULL UNIQUE,
    stock_type    stock_type          NOT NULL,
    cost          FLOAT               NOT NULL,
    currency      INTEGER             REFERENCES currencies(currency_id),
    PRIMARY KEY (stock_id, cost, currency)
);
CREATE TABLE if not exists stocks_items
(
    stocks_item_id SERIAL,
    portfolio      INTEGER            REFERENCES portfolios(portfolio_id) ON DELETE CASCADE,
    stock_item     INTEGER,
    stock_cost     FLOAT,
    stock_currency INTEGER,
    amount         INTEGER,
    stock_value    FLOAT              GENERATED ALWAYS AS (amount * stock_cost) STORED,
    created_at     TIMESTAMPTZ        NOT NULL DEFAULT current_timestamp,
    FOREIGN KEY (stock_item, stock_cost, stock_currency) REFERENCES stocks (stock_id, cost, currency)
);

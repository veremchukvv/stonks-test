-- insert into portfolios(user_id, user_auth_type, portfolio_name, description, is_public) values (1, 'local', 'test1', 'no description', true);
-- insert into portfolios(user_id, user_auth_type, portfolio_name, description, is_public) values (1, 'local', 'test2', 'no description', true);
-- insert into portfolios(user_id, user_auth_type, portfolio_name, description, is_public) values (1, 'local', 'test3', 'no description', true);

insert into currencies(currency_name, currency_ticker) values ('Russian Rouble', 'RUR');
insert into currencies(currency_name, currency_ticker) values ('United States Dollar', 'USD');
insert into currencies(currency_name, currency_ticker) values ('Euro', 'EUR');

insert into stocks(stock_name, description, ticker, stock_type, cost, currency) values ('NULL_RUB', 'NULL_RUB', 'NULL_RUB', 'stock', 0, 1);
insert into stocks(stock_name, description, ticker, stock_type, cost, currency) values ('NULL_USD', 'NULL_USD', 'NULL_USD', 'stock', 0, 2);
insert into stocks(stock_name, description, ticker, stock_type, cost, currency) values ('NULL_EUR', 'NULL_EUR', 'NULL_EUR', 'stock', 0, 3);
insert into stocks(stock_name, description, ticker, stock_type, cost, currency) values ('Yandex', 'Russian IT Company', 'YNDX', 'stock', 100, 1);
insert into stocks(stock_name, description, ticker, stock_type, cost, currency) values ('Kamaz', 'Russian automaker company', 'KMAZ', 'stock', 50, 1);
insert into stocks(stock_name, description, ticker, stock_type, cost, currency) values ('Facebook', 'US IT Company', 'FB', 'stock', 200, 2);
insert into stocks(stock_name, description, ticker, stock_type, cost, currency) values ('Tesla', 'US automaker company', 'TSLA', 'stock', 1000, 2);
insert into stocks(stock_name, description, ticker, stock_type, cost, currency) values ('Siemens', 'European Company', 'SIE', 'stock', 300, 3);
insert into stocks(stock_name, description, ticker, stock_type, cost, currency) values ('Volkswagen', 'European automaker company', 'VOW', 'stock', 500, 3);

insert into closed_deals(portfolio, stock_item, stock_cost, stock_currency, amount, closed_at, buy_cost, sell_cost, income_money, income_percent) values (7, 6, 200, 2, 3, NOW(), 100, 300, 200, 20);
insert into closed_deals(portfolio, stock_item, stock_cost, stock_currency, amount, closed_at, buy_cost, sell_cost, income_money, income_percent) values (7, 6, 200, 2, 3, NOW(), 100, 300, 200, 20);
insert into closed_deals(portfolio, stock_item, stock_cost, stock_currency, amount, closed_at, buy_cost, sell_cost, income_money, income_percent) values (7, 4, 100, 1, 2, NOW(), 200, 100, -100, -20);
insert into closed_deals(portfolio, stock_item, stock_cost, stock_currency, amount, closed_at, buy_cost, sell_cost, income_money, income_percent) values (7, 3, 0, 3, 1, NOW(), 200, 100, -100, -20);

-- insert into balances(portfolio_id, currency_id, money_value) values (1, 1, 1000);
-- insert into balances(portfolio_id, currency_id, money_value) values (1, 2, 1500);
-- insert into balances(portfolio_id, currency_id, money_value) values (1, 3, 2000);

-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (1, 1, 0, 1, 1);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (1, 1, 0, 2, 1);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (1, 1, 0, 3, 1);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (2, 1, 0, 1, 1);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (2, 1, 0, 2, 1);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (2, 1, 0, 3, 1);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (3, 1, 0, 1, 1);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (3, 1, 0, 2, 1);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (3, 1, 0, 3, 1);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (5, 1, 100, 1, 10);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (5, 2, 50, 1, 20);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (5, 3, 200, 2, 30);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (5, 4, 1000, 2, 40);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (5, 5, 300, 3, 50);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (5, 6, 500, 3, 60);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (2, 1, 100, 1, 5);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (2, 2, 50, 1, 3);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (3, 1, 100, 1, 10);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (3, 2, 50, 1, 20);
--
--
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (7, 1, 100, 1, 10);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (7, 2, 50, 1, 20);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (7, 3, 200, 2, 30);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (7, 4, 1000, 2, 40);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (7, 5, 300, 3, 50);
-- insert into deals(portfolio, stock_item, stock_cost, stock_currency, amount) values (7, 6, 500, 3, 60);
insert into portfolios(portfolio_id, user_id, user_auth_type, portfolio_name, description, is_public) values (1, 1, 'local', 'test1', 'no description', true);
insert into portfolios(portfolio_id, user_id, user_auth_type, portfolio_name, description, is_public) values (2, 1, 'local', 'test2', 'no description', true);
insert into portfolios(portfolio_id, user_id, user_auth_type, portfolio_name, description, is_public) values (3, 1, 'local', 'test3', 'no description', true);

insert into currencies(currency_id, currency_name, ticker) values (1, 'Russian Rouble', 'RUR');
insert into currencies(currency_id, currency_name, ticker) values (2, 'United States Dollar', 'USD');
insert into currencies(currency_id, currency_name, ticker) values (3, 'Euro', 'EUR');

insert into balances(balance_id, portfolio_id, currency_id, money_value) values (1, 1, 1, 1000);
insert into balances(balance_id, portfolio_id, currency_id, money_value) values (2, 1, 2, 1500);
insert into balances(balance_id, portfolio_id, currency_id, money_value) values (3, 1, 3, 2000);

insert into stocks(stock_id, stock_name, description, ticker, stock_type, cost, currency) values (1, 'NULL', 'NULL', 'NULL', 'stock', 0, 1);
insert into stocks(stock_id, stock_name, description, ticker, stock_type, cost, currency) values (2, 'Yandex', 'Russian IT Company', 'YNDX', 'stock', 100, 1);
insert into stocks(stock_id, stock_name, description, ticker, stock_type, cost, currency) values (3, 'Kamaz', 'Russian automaker company', 'KMAZ', 'stock', 50, 1);
insert into stocks(stock_id, stock_name, description, ticker, stock_type, cost, currency) values (4, 'Facebook', 'US IT Company', 'FB', 'stock', 200, 2);
insert into stocks(stock_id, stock_name, description, ticker, stock_type, cost, currency) values (5, 'Tesla', 'US automaker company', 'TSLA', 'stock', 1000, 2);
insert into stocks(stock_id, stock_name, description, ticker, stock_type, cost, currency) values (6, 'Siemens', 'European Company', 'SIE', 'stock', 300, 3);
insert into stocks(stock_id, stock_name, description, ticker, stock_type, cost, currency) values (7, 'Volkswagen', 'European automaker company', 'VOW', 'stock', 500, 3);

insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (1, 1, 1, 0, 1, 1);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (1, 1, 2, 100, 1, 10);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (2, 1, 3, 50, 1, 20);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (3, 1, 4, 200, 2, 30);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (4, 1, 5, 1000, 2, 40);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (5, 1, 6, 300, 3, 50);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (6, 1, 7, 500, 3, 60);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (1, 2, 1, 0, 1, 1);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (7, 2, 2, 100, 1, 5);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (8, 2, 3, 50, 1, 3);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (9, 3, 2, 100, 1, 10);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (10, 3, 3, 50, 1, 20);
insert into stocks_items(stocks_item_id, portfolio, stock_item, stock_cost, stock_currency, amount) values (11, 4, 1, 0, 1, 1);
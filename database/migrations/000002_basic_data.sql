insert into portfolios(portfolio_id, user_id, user_auth_type, portfolio_name, description, is_public) values (1, 1, 'local', 'test1', 'no description', true);
insert into portfolios(portfolio_id, user_id, user_auth_type, portfolio_name, description, is_public) values (2, 1, 'local', 'test2', 'no description', true);
insert into portfolios(portfolio_id, user_id, user_auth_type, portfolio_name, description, is_public) values (3, 1, 'local', 'test3', 'no description', true);

insert into currencies(currency_id, currency_name, ticker) values (1, 'Russian Rouble', 'RUR');
insert into currencies(currency_id, currency_name, ticker) values (2, 'United States Dollar', 'USD');
insert into currencies(currency_id, currency_name, ticker) values (3, 'Euro', 'EUR');

insert into balances(balance_id, portfolio_id, currency_id, amount) values (1, 1, 1, 1000);
insert into balances(balance_id, portfolio_id, currency_id, amount) values (2, 1, 2, 1500);
insert into balances(balance_id, portfolio_id, currency_id, amount) values (3, 1, 3, 2000);

insert into stocks(stock_id, stock_name, description, ticker, stock_type, cost, currency) values (1, 'Yandex', 'Russian IT Company', 'YNDX', 'stock', 100, 1);
insert into stocks(stock_id, stock_name, description, ticker, stock_type, cost, currency) values (2, 'Facebook', 'US IT Company', 'FB', 'stock', 200, 2);
insert into stocks(stock_id, stock_name, description, ticker, stock_type, cost, currency) values (3, 'Siemens', 'European Company', 'SIE', 'stock', 300, 3);

insert into stocks_items(stocks_item_id, portfolio, stock_item, amount) values (1, 1, 1, 10);
insert into stocks_items(stocks_item_id, portfolio, stock_item, amount) values (2, 1, 2, 20);
insert into stocks_items(stocks_item_id, portfolio, stock_item, amount) values (3, 1, 3, 30);
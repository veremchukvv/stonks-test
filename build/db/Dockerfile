FROM bitnami/postgresql:14.1.0
ADD ./database/migrations/000000_create_database.sql /docker-entrypoint-initdb.d/
ADD ./database/migrations/*up.sql /docker-entrypoint-initdb.d/
ENV POSTGRES_DB stonks
ENV POSTGRES_USER stonks_admin
ENV POSTGRES_PASSWORD postgres123
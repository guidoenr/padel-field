# PostgreSQL
FROM postgres:14.5 as db

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
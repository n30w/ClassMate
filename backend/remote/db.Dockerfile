FROM postgres:16
LABEL authors="Neo"

# Uses dev-init.sql to initialize the database for development,
# with appropriate tables and data.
COPY development/dev-init.sql /docker-entrypoint-initdb.d/
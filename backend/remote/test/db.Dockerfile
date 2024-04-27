FROM postgres:16
LABEL authors="Neo"

# Uses dev-init.sql to initialize the database for testing,
# with appropriate tables and data.
COPY test-init.sql /docker-entrypoint-initdb.d/
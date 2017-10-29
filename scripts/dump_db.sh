docker exec core_database_1 pg_dump -U postgres lowside-dev > $1-dump.sql

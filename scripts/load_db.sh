cat $1 | docker exec -i core_database_1 su postgres -c 'psql lowside-dev'

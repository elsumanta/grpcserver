#!/bin/sh
# do not use this script on production
# this script is for development phase only

dbname="grpc-server"
port=${PGPORT:-5656}
user="server"
password="server"
host=${PGHOST:-localhost}

echo "host $host port: $port"
PGPASSWORD=$password psql -U $user -d "postgres" -h $host -p $port -c "CREATE DATABASE $dbname" 2> /dev/null
for filename in schema/*.sql; do
    PGPASSWORD=$password psql -h $host -p $port -d $dbname -U $user -f "$filename"
done

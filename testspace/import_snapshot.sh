#!/bin/bash

# Function to check if a command exists
command_exists() {
  command -v "$1" >/dev/null 2>&1
}

# Check if unzip, wget, and psql commands exist
if ! command_exists unzip; then
  echo "Error: 'unzip' command not found. Please install unzip and try again."
  exit 1
fi

if ! command_exists wget; then
  echo "Error: 'wget' command not found. Please install wget and try again."
  exit 1
fi

if ! command_exists psql; then
  echo "Error: 'psql' command not found. Please install PostgreSQL client and try again."
  exit 1
fi

# Ask for variables with _EXTRACT and export them
read -p "Enter Postgres Server: " PGHOST_EXTRACT
read -p "Enter Postgres Port: " PGPORT_EXTRACT
read -p "Enter Postgres User: " PGUSER_EXTRACT
read -p "Enter Postgres Password: " PGPASSWORD_EXTRACT
read -p "Enter Postgres Database: " PGDB_EXTRACT

export PGHOST=${PGHOST_EXTRACT}
export PGUSER=${PGUSER_EXTRACT}
export PGDB=${PGDB_EXTRACT}
export PGPASSWORD=${PGPASSWORD_EXTRACT}
export PGPORT=${PGPORT_EXTRACT}

# Download and unzip the file from rest.diadata.org
wget -qO snapshot.zip http://rest.diadata.org/snapshot/latest
unzip -o snapshot.zip

# Run the psql commands
psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-blockchain.sql
psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-exchange.sql
psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-asset.sql
psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-pool.sql
psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-poolasset.sql
psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-exchangepair.sql

# Delete the output-*.sql files
rm output-blockchain.sql

echo "Data import and processing completed successfully."

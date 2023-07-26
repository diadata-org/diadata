#!/usr/bin/env bash

set -euo pipefail

echo "Extracting a snapshot of remote data ..."
export PGHOST=${PGHOST_EXTRACT} PGUSER=${PGUSER_EXTRACT} PGDB=${PGDB_EXTRACT} PGPASSWORD=${PGPASSWORD_EXTRACT}; \
psql --host ${PGHOST} --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/dump-asset.sql --output /tmp/dump-asset.csv; \
psql --host ${PGHOST} --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/dump-blockchain.sql --output /tmp/dump-blockchain.csv; \
psql --host ${PGHOST} --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/dump-exchange_cex.sql --output /tmp/dump-exchange_cex.csv; \
psql --host ${PGHOST} --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/dump-exchange_dex.sql --output /tmp/dump-exchange_dex.csv; \
psql --host ${PGHOST} --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/dump-exchangepair.sql --output /tmp/dump-exchangepair.csv; \
psql --host ${PGHOST} --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/dump-pool.sql --output /tmp/dump-pool.csv; \
psql --host ${PGHOST} --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/dump-poolasset.sql --output /tmp/dump-poolasset.csv
echo "Snapshot completed"

echo "Loading data ..."
export PGHOST= PGUSER=${PGUSER_TEMP} PGDB=${PGDB_TEMP} PGPASSWORD=${PGPASSWORD_TEMP}
psql --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/load-blockchain.sql
psql --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/load-exchange.sql
psql --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/load-asset.sql
psql --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/load-pool.sql
psql --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/load-poolasset.sql
psql --port 5432 --username ${PGUSER} --dbname ${PGDB} --file /mnt/env-context/scripts/load-exchangepair.sql
echo "Load completed"

echo "Dumping data ..."
pg_dump -p 5432 -U $PGUSER --format c --blobs --verbose --dbname $PGDB --schema public --file /mnt/env-workdir/pg_dump.backup
pg_restore --data-only --file /mnt/env-workdir/pg_dump.sql /mnt/env-workdir/pg_dump.backup
echo "Dump completed"
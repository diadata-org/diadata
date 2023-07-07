\COPY exchange FROM '/tmp/dump-exchange_cex.csv' WITH (FORMAT csv, DELIMITER ';');
\COPY exchange FROM '/tmp/dump-exchange_dex.csv' WITH (FORMAT csv, DELIMITER ';');
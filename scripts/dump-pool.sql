COPY (
    SELECT * FROM pool WHERE pool_id IN (
        SELECT DISTINCT pool_id FROM poolasset WHERE "time_stamp" >= now() - interval '7 days'
    )
) TO STDOUT WITH (format csv, delimiter ';');
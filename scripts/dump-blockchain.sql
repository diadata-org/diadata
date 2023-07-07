COPY (
    SELECT * FROM blockchain WHERE "name" IN (
        SELECT DISTINCT blockchain FROM asset WHERE asset_id IN (
            SELECT DISTINCT asset_id FROM poolasset WHERE "time_stamp" >= now() - interval '7 days'
        )
    )
) TO STDOUT WITH (format csv, delimiter ';');
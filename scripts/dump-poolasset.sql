COPY (
    SELECT * FROM poolasset WHERE "time_stamp" >= now() - interval '7 days'
) TO STDOUT WITH (format csv, delimiter ';');
COPY (
    SELECT * FROM asset
) TO STDOUT WITH (format csv, delimiter ';');
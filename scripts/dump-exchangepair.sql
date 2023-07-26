COPY (
    SELECT * FROM exchangepair
) TO STDOUT WITH (format csv, delimiter ';');
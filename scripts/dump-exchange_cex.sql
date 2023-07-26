COPY (
    SELECT * FROM exchange WHERE centralized = true
) TO STDOUT WITH (format csv, delimiter ';');
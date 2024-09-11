CREATE OR REPLACE FUNCTION ensure_mandatory_read_write() 
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'DELETE' AND OLD.access_level = 'read_write' THEN
        IF NOT EXISTS (
            SELECT 1 
            FROM wallet_public_keys 
            WHERE customer_id = OLD.customer_id 
            AND access_level = 'read_write'
            AND key_id != OLD.key_id 
        ) THEN
            RAISE EXCEPTION 'Customer % must have at least one read_write access.', OLD.customer_id;
        END IF;
    END IF;

    IF TG_OP = 'UPDATE' AND OLD.access_level = 'read_write' AND NEW.access_level != 'read_write' THEN
        IF NOT EXISTS (
            SELECT 1 
            FROM wallet_public_keys 
            WHERE customer_id = OLD.customer_id 
            AND access_level = 'read_write'
            AND key_id != OLD.key_id
        ) THEN
            RAISE EXCEPTION 'Customer % must have at least one read_write access.', OLD.customer_id;
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER check_delete_read_write
BEFORE DELETE ON wallet_public_keys
FOR EACH ROW
EXECUTE FUNCTION ensure_mandatory_read_write();



CREATE TRIGGER check_update_read_write
BEFORE UPDATE ON wallet_public_keys
FOR EACH ROW
EXECUTE FUNCTION ensure_mandatory_read_write();
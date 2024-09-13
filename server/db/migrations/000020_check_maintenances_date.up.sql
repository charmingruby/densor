CREATE OR REPLACE FUNCTION check_maintenance_date()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.start_date < NOW() THEN
        RAISE EXCEPTION 'A data de início da manutenção não pode ser anterior à data atual.';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_check_maintenance_date
BEFORE INSERT OR UPDATE ON maintenances
FOR EACH ROW
EXECUTE FUNCTION check_maintenance_date();
CREATE OR REPLACE FUNCTION prevent_admin_deletion()
RETURNS TRIGGER AS $$
BEGIN
    IF OLD.role = 'admin' THEN
        RAISE EXCEPTION 'Não é permitido excluir um administrador.';
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_prevent_admin_deletion
BEFORE DELETE ON users
FOR EACH ROW
EXECUTE FUNCTION prevent_admin_deletion();


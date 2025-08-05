DROP TRIGGER IF EXISTS update_clients_updated_at ON clients;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP INDEX IF EXISTS idx_clients_cuit;
DROP TABLE IF EXISTS clients;
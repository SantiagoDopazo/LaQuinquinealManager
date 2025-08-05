-- Script para verificar que la migración funcionó correctamente

-- 1. Ver todas las tablas
SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public' 
ORDER BY table_name;

-- 2. Ver estructura de clients
SELECT column_name, data_type, is_nullable, column_default
FROM information_schema.columns 
WHERE table_name = 'clients'
ORDER BY ordinal_position;

-- 3. Ver estructura de orders (debe tener client_id)
SELECT column_name, data_type, is_nullable, column_default
FROM information_schema.columns 
WHERE table_name = 'orders'
ORDER BY ordinal_position;

-- 4. Ver algunos clientes creados
SELECT id, name, cuit, created_at FROM clients LIMIT 5;

-- 5. Ver órdenes con client_id
SELECT id, order_number, client_id FROM orders LIMIT 5;

-- 6. Ver join entre órdenes y clientes
SELECT 
    o.id as order_id,
    o.order_number,
    c.name as client_name,
    c.cuit as client_cuit
FROM orders o
JOIN clients c ON o.client_id = c.id
LIMIT 5;
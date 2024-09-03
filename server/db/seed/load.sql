-- Data load to users table
INSERT INTO users (id, name, email, password, role, is_active, created_at)
VALUES 
('user_001', 'John Doe', 'john.doe@example.com', 'password123', 'admin', TRUE, NOW()),
('user_002', 'Jane Smith', 'jane.smith@example.com', 'password456', 'user', TRUE, NOW()),
('user_003', 'Alice Johnson', 'alice.johnson@example.com', 'password789', 'user', FALSE, NOW());

-- Data load to sectors table
INSERT INTO sectors (id, name, localization, responsible_id, created_at)
VALUES 
('sector_001', 'Electrical', 'Building A', 'user_001', NOW()),
('sector_002', 'Mechanical', 'Building B', 'user_002', NOW());

-- Data load to equipments table
INSERT INTO equipments (id, name, model, manufacturer, sector_id, created_at)
VALUES 
('equip_001', 'Generator', 'GXT2000', 'PowerCo', 'sector_001', NOW()),
('equip_002', 'HVAC', 'HVAC-X500', 'ClimateTech', 'sector_002', NOW());

-- Data load to sensor_categories table
INSERT INTO sensor_categories (id, name, description, created_at)
VALUES 
('cat_001', 'Temperature', 'Temperature sensors for climate control', NOW()),
('cat_002', 'Pressure', 'Pressure sensors for hydraulic systems', NOW());

-- Data load to sensors table
INSERT INTO sensors (id, name, description, sensor_category_id, equipment_id, status, observation, sector_id, created_at)
VALUES 
('sensor_001', 'Sensor de Temp 1', 'Sensor para monitoramento de temperatura ambiente', 'cat_001', 'equip_002', 'ATIVO', 'Funciona corretamente', 'sector_002', NOW()),
('sensor_002', 'Sensor de Pressão 1', 'Sensor para pressão hidráulica', 'cat_002', 'equip_001', 'ATIVO', 'Sem problemas', 'sector_001', NOW()),
('sensor_003', 'Sensor de Temp 2', 'Sensor para monitoramento de temperatura externa', 'cat_001', 'equip_001', 'ATIVO', 'Calibrado recentemente', 'sector_001', NOW()),
('sensor_004', 'Sensor de Pressão 2', 'Sensor para medir pressão de fluidos', 'cat_002', 'equip_002', 'DEFEITUOSO', 'Necessita manutenção', 'sector_002', NOW()),
('sensor_005', 'Sensor de Temp 3', 'Sensor para monitoramento de temperatura de armazém', 'cat_001', 'equip_001', 'ATIVO', 'Funciona corretamente', 'sector_001', NOW()),
('sensor_006', 'Sensor de Pressão 3', 'Sensor para monitoramento de pressão de gás', 'cat_002', 'equip_002', 'ATIVO', 'Sem problemas detectados', 'sector_002', NOW()),
('sensor_007', 'Sensor de Temp 4', 'Sensor para monitoramento de temperatura da sala de servidores', 'cat_001', 'equip_002', 'ATIVO', 'Funciona como esperado', 'sector_002', NOW()),
('sensor_008', 'Sensor de Pressão 4', 'Sensor para monitoramento de pressão hidráulica', 'cat_002', 'equip_001', 'MANUTENÇÃO', 'Requer calibração', 'sector_001', NOW()),
('sensor_009', 'Sensor de Temp 5', 'Sensor para monitoramento de temperatura da fábrica', 'cat_001', 'equip_001', 'ATIVO', 'Operacional', 'sector_001', NOW()),
('sensor_010', 'Sensor de Pressão 5', 'Sensor para medir pressão de ar', 'cat_002', 'equip_002', 'ATIVO', 'Em boas condições', 'sector_002', NOW()),
('sensor_011', 'Sensor de Temp 6', 'Sensor para monitoramento de temperatura do laboratório', 'cat_001', 'equip_002', 'DEFEITUOSO', 'Aguardando reparo', 'sector_002', NOW()),
('sensor_012', 'Sensor de Pressão 6', 'Sensor para pressão do sistema hidráulico', 'cat_002', 'equip_001', 'ATIVO', 'Recentemente revisado', 'sector_001', NOW()),
('sensor_013', 'Sensor de Temp 7', 'Sensor para monitoramento de temperatura ambiental', 'cat_001', 'equip_001', 'ATIVO', 'Sem problemas', 'sector_001', NOW()),
('sensor_014', 'Sensor de Pressão 7', 'Sensor para medição de pressão em tubulações', 'cat_002', 'equip_002', 'ATIVO', 'Funciona corretamente', 'sector_002', NOW()),
('sensor_015', 'Sensor de Temp 8', 'Sensor para monitoramento de temperatura crítica', 'cat_001', 'equip_001', 'MANUTENÇÃO', 'Necessita inspeção', 'sector_001', NOW()),
('sensor_016', 'Sensor de Pressão 8', 'Sensor para monitoramento de pressão do sistema', 'cat_002', 'equip_002', 'ATIVO', 'Operacional', 'sector_002', NOW()),
('sensor_017', 'Sensor de Temp 9', 'Sensor para monitoramento de temperatura industrial', 'cat_001', 'equip_001', 'ATIVO', 'Funcionando bem', 'sector_001', NOW()),
('sensor_018', 'Sensor de Pressão 9', 'Sensor para medição de alta pressão', 'cat_002', 'equip_002', 'MANUTENÇÃO', 'Requer calibração', 'sector_002', NOW()),
('sensor_019', 'Sensor de Temp 10', 'Sensor para controle de temperatura em armazenamento', 'cat_001', 'equip_001', 'ATIVO', 'Calibrado', 'sector_001', NOW()),
('sensor_020', 'Sensor de Pressão 10', 'Sensor para monitoramento de níveis de pressão hidráulica', 'cat_002', 'equip_002', 'ATIVO', 'Sem problemas', 'sector_002', NOW()),
('sensor_021', 'Sensor de Temp 11', 'Sensor para monitoramento de temperatura ambiente', 'cat_001', 'equip_001', 'MANUTENÇÃO', 'Necessita calibração', 'sector_001', NOW()),
('sensor_022', 'Sensor de Pressão 11', 'Sensor para monitoramento de pressão em tanques', 'cat_002', 'equip_002', 'ATIVO', 'Em bom estado', 'sector_002', NOW());
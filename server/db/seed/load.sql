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
('sensor_001', 'Temp Sensor 1', 'Sensor for monitoring room temperature', 'cat_001', 'equip_002', 'active', 'Works properly', 'sector_002', NOW()),
('sensor_002', 'Pressure Sensor 1', 'Sensor for hydraulic pressure', 'cat_002', 'equip_001', 'active', 'No issues', 'sector_001', NOW());

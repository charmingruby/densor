CREATE TABLE IF NOT EXISTS supplier_stock_materials
(
    id varchar PRIMARY KEY NOT NULL,
    stock_quantity int,
    minimum_quantity int,
    supplier_id varchar,
    stock_material_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,

    FOREIGN KEY (supplier_id) REFERENCES suppliers(id),
    FOREIGN KEY (stock_material_id) REFERENCES stock_materials(id)
);
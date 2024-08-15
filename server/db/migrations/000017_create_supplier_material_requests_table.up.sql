CREATE TABLE IF NOT EXISTS supplier_material_requests
(
    id varchar PRIMARY KEY NOT NULL,
    necessary_at timestamp,
    status varchar,
    quantity integer,
    supplier_stock_material_id varchar,
    created_at timestamp DEFAULT now() NOT NULL,

    FOREIGN KEY (supplier_stock_material_id) REFERENCES supplier_stock_materials(id)
);
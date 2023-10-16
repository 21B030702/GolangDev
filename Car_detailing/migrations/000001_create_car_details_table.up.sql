CREATE TABLE IF NOT EXISTS car_details (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    description text NOT NULL,
    dateofproduction text NOT NULL,
    weight float NOT NULL,
    material text NOT NULL,
    price int NOT NULL,
    version integer NOT NULL DEFAULT 1
);
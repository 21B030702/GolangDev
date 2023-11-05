ALTER TABLE car_details ADD CONSTRAINT car_details_weight_check CHECK (weight > 0);
ALTER TABLE car_details ADD CONSTRAINT car_details_price_length_check CHECK (price > 0);
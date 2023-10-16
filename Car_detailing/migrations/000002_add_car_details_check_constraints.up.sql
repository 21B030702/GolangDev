ALTER TABLE car_details ADD CONSTRAINT car_details_weight_check CHECK (weight > 0);
ALTER TABLE car_details ADD CONSTRAINT car_details_date_of_production_check CHECK (year BETWEEN 1888 AND date_part('year', now()));
ALTER TABLE car_details ADD CONSTRAINT car_details_price_length_check CHECK (array_length(genres, 1) BETWEEN 1 AND 5);
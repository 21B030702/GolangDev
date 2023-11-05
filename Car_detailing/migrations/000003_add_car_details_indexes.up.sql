CREATE INDEX IF NOT EXISTS car_details_title_idx ON car_details USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS car_details_genres_idx ON car_details USING GIN (material);
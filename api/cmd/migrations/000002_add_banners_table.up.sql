CREATE TABLE IF NOT EXISTS banners
(
  banner_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  image_id VARCHAR(255),
  mobile_image_id VARCHAR(255),
  collection_id INT REFERENCES collections(collection_id) ON DELETE SET NULL,
  created_at TIMESTAMP DEFAULT NOW()
);
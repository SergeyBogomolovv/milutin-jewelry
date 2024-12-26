CREATE TABLE IF NOT EXISTS collections
(
	collection_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT,
	image_id VARCHAR(255),
	created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS collection_items
(
	item_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	collection_id INT REFERENCES collections(collection_id) ON DELETE CASCADE,
	title TEXT,
	description TEXT,
	image_id VARCHAR(255),
	created_at TIMESTAMP DEFAULT NOW()
);

DROP TABLE IF EXISTS feeds CASCADE;
CREATE TABLE feeds (
	id  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        name TEXT NOT NULL,
	url TEXT UNIQUE NOT NULL,
	user_id UUID REFERENCES users(id) ON DELETE CASCADE
);


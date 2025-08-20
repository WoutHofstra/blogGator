
DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
        id  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        title TEXT NOT NULL,
        url TEXT UNIQUE NOT NULL,
	description TEXT,
	published_at TIMESTAMP,
        feed_id UUID REFERENCES feeds(id) ON DELETE CASCADE
);

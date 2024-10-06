-- file_metadata table
CREATE TABLE IF NOT EXISTS file_metadata (
    file_id UUID PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    part_count INT NOT NULL
);

-- file_parts table
CREATE TABLE IF NOT EXISTS file_parts (
    file_id UUID NOT NULL,
    part_number INT NOT NULL,
    data BYTEA NOT NULL,
    PRIMARY KEY (file_id, part_number),
    FOREIGN KEY (file_id) REFERENCES file_metadata(file_id) ON DELETE CASCADE
);

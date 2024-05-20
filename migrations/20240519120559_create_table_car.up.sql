-- Create manufacturer table
CREATE TABLE manufacturer (
                              id SERIAL PRIMARY KEY,
                              name VARCHAR(255) NOT NULL,
                              hsn VARCHAR(255) NOT NULL
);

-- Create car table with a foreign key referencing manufacturer
CREATE TABLE cars (
                     id SERIAL PRIMARY KEY,
                     name VARCHAR(255) NOT NULL,
                     tsn VARCHAR(255) NOT NULL,
                     manufacturer_id INTEGER NOT NULL,
                     CONSTRAINT fk_manufacturer FOREIGN KEY (manufacturer_id) REFERENCES manufacturer(id)
);

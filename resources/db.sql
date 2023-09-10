CREATE DATABASE thuggeeks;

CREATE TABLE thug_geeks (
                            id serial PRIMARY KEY,
                            title varchar(255) NOT NULL,
                            image varchar(255) NOT NULL,
                            description_short text NOT NULL,
                            description_long text NOT NULL,
                            tag varchar(50) NOT NULL,
                            like_count int NOT NULL
);
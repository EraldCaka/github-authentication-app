CREATE TABLE users (
                       id bigserial PRIMARY KEY,
                       name varchar,
                       username varchar,
                       img_url varchar
);
CREATE TABLE starred_repos (
                               id bigserial PRIMARY KEY,
                               user_id bigint REFERENCES users(id),
                               repo_name varchar,
                               repo_owner varchar
);
CREATE TABLE commits (
                         id bigserial PRIMARY KEY,
                         commit_sha varchar NOT NULL,
                         date timestamp with time zone,
                         repo_id bigint REFERENCES starred_repos(id)
);


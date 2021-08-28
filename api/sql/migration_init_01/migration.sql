DROP TABLE component_data;
DROP TABLE components;
DROP TABLE project_tags;
DROP TABLE projects;
DROP TABLE priorities;
DROP TABLE tags;
DROP TABLE users;

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL
);

create TABLE tags (
    id BIGSERIAL PRIMARY KEY,
    "user_id" BIGSERIAL NOT NULL references users(id),
    "name" VARCHAR(255) NOT NULL
);

create TABLE priorities (
    id BIGSERIAL PRIMARY KEY,
    "name" VARCHAR(20) NOT NULL UNIQUE
);


create TABLE projects (
    id BIGSERIAL PRIMARY KEY,
    "user_id" BIGSERIAL NOT NULL references users(id),
    title VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "priority_id"   BIGSERIAL NOT NULL references priorities(id)
);

create TABLE project_tags (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGSERIAL NOT NULL references projects(id),
    tag_id BIGSERIAL NOT NULL references tags(id)
);

create TABLE components (
    id BIGSERIAL PRIMARY KEY,
    "project_id" BIGSERIAL NOT NULL references projects(id),
    "title" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL
);

create TABLE component_data (
    id BIGSERIAL PRIMARY KEY,
    "component_id" BIGSERIAL NOT NULL references components(id),
    "key" VARCHAR(255) NOT NULL,
    "value" TEXT NOT NULL
);


INSERT INTO priorities ("name")
VALUES
    ('Low'),
    ('Medium'),
    ('High');
-- DDL
CREATE TABLE t_name_parent (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    parent_id INT
);


-- DML
INSERT INTO t_name_parent (name, parent_id)
VALUES ('Zaki', 2),
       ('Ilham', NULL),
       ('Irwan', 2),
       ('Arka', 3);
-- DML
SELECT t1.id, t1.name, t2.name AS parent_name
FROM t_name_parent t1
LEFT JOIN t_name_parent t2 ON t1.parent_id = t2.id;
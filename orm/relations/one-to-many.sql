CREATE TABLE a (
    id INT PRIMARY KEY
);
CREATE TABLE b (
    a_id int,
    FOREIGN KEY (a_id) REFERENCES a(id)
);

SELECT * from a LEFT JOIN b ON a.id=b.a_id;

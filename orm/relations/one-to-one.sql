CREATE TABLE b (
    id INT PRIMARY KEY
);
CREATE TABLE a (
    b_id int,
    FOREIGN KEY (b_id) REFERENCES b(id)
);

SELECT * from a LEFT JOIN b ON a.b_id=b.id;

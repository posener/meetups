CREATE TABLE a (
    id INT PRIMARY KEY
);
CREATE TABLE b (
    id INT PRIMARY KEY
);
CREATE TABLE ab_relation (
    a_id INT,
    b_id INT,
    FOREIGN KEY (a_id) REFERENCES a(id),
    FOREIGN KEY (b_id) REFERENCES b(id)
);

SELECT * from a
    LEFT JOIN ab_relation ON a.id=ab_relation.a_id
    LEFT JOIN b ON ab_relation.b_id=b.id;

SELECT * from b
    LEFT JOIN ab_relation ON b.id=ab_relation.b_id
    LEFT JOIN a ON ab_relation.a_id=a.id;

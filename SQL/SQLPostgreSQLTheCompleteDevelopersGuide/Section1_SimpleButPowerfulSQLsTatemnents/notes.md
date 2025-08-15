## 4. Database Design
- What kind of thing are we storing?
- What properties does this thing have?
- What type of data does each of those properties contain?

## 5. Creating Tables
```sql
CREATE TABLE cities (
  name VARCHAR(50),
  country VARCHAR(50),
  population INTEGER,
  area INTEGER  
);
```

## 6. Analyzing CREATE TABLE
Keywords: Tells the DB what you want to do. Write these in uppercase (ex: `CREATE TABLE`)
Identifier: Identifies the thing you want to work on. Always written in lower case (ex: `cities`).

## 7. Inserting Data Into a Table
```sql
INSERT INTO cities (name, country, population, area)
VALUES ('Tokyo', 'Japan', 38000000, 8223);
```

```sql
INSERT INTO cities (name, country, population, area)
VALUES 
    ('Delhi', 'India', 28000000, 2240),
    ('Shanghai', 'China', 22000000, 1250),
    ('Sao Paulo', 'Brazil', 22000000, 3043);
```

## 8. Retrieving Data with Select
```sql
SELECT * FROM cities;
SELECT name, country FROM cities;
```

## 9. Calculated Columns
- You can process data before pulling it from the db, uh-oh
```sql
SELECT name, population / area AS population_density FROM cities;
```
Dang noice. There are math operators!

## 12. String Operators and Functions
- Ooooh there are string operators and functions, noice!
```sql
SELECT name || ', ' || country AS location FROM cities;
SELECT CONCAT(name, ', ',country) AS location FROM cities;
SELECT UPPER (CONCAT(name, ', ',country)) AS location FROM cities;


```
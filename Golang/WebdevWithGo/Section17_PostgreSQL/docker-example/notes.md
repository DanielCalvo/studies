
```sql
CREATE DATABASE employees;
```

- \c employees
- SELECT current_database();

## 132
```sql
DROP DATABASE employees;
```
```sql
CREATE DATABASE company;
```


## 133
- \c company
- \d

```sql
CREATE TABLE employees (
   ID INT PRIMARY KEY     NOT NULL,
   NAME           TEXT    NOT NULL,
   RANK           INT     NOT NULL,
   ADDRESS        CHAR(50),
   SALARY         REAL DEFAULT 25500.00,
   BDAY			  DATE DEFAULT '1900-01-01'
);
```

- \d
- \d employees
- DROP TABLE employees;

## 134
- `psql -h localhost -U postgres`
- `\c company`

```sql
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (1, 'Mark', 7, '1212 E. Lane, Someville, AK, 57483', 43000.00 ,'1992-01-13');
SELECT * FROM employees;
```

Omitted values will have the default value
```sql
INSERT INTO employees (ID,NAME,RANK,ADDRESS,BDAY) VALUES (2, 'Marian', 8, '7214 Wonderlust Ave, Lost Lake, KS, 22897', '1989-11-21');
```

You can also use the DEFAULT instead of specifying a value: 
```sql
INSERT INTO employees (ID,NAME,RANK,ADDRESS,SALARY,BDAY) VALUES (3, 'Maxwell', 6, '7215 Jasmine Place, Corinda, CA 98743', 87500.00, DEFAULT);
```

## 135
```sql
CREATE TABLE phonenumbers(
	ID  SERIAL PRIMARY KEY,
	PHONE           TEXT      NOT NULL
);
```
```sql
INSERT INTO phonenumbers (PHONE) VALUES ( '234-432-5234'), ('543-534-6543'), ('312-123-5432');
```

## 136
- DROP TABLE employees;
```sql
CREATE TABLE employees (
    ID  SERIAL PRIMARY KEY NOT NULL,
    NAME TEXT NOT NULL,
    SCORE INT DEFAULT 10 NOT NULL,
    SALARY REAL
);
```
- INSERT INTO employees (NAME, SCORE, SALARY) VALUES ('Daniel', '10', 1);
- SELECT * FROM employees;

## 139 Cross join
- A cross join will produce ros which combine each row from the first table with each row from the second table

```sql
CREATE TABLE person (
                        ID  SERIAL PRIMARY KEY NOT NULL,
                        NAME           CHAR(50) NOT NULL
);
INSERT INTO person (NAME) VALUES ('Shen'), ('Daniel'), ('Juan'), ('Arin'), ('McLeod');
```

```sql
CREATE TABLE sport (
   ID  SERIAL PRIMARY KEY NOT NULL,
   NAME           CHAR(50) NOT NULL,
   P_ID         INT      references person(ID)
);
INSERT INTO sport (NAME, P_ID) VALUES ('Surf',1),('Soccer',3),('Ski',3),('Sail',3),('Bike',3);
```

## 140 Inner join
- The ID of a record in one table is connected with the ID of a record in another table, or something like that
```sql
SELECT person.name, sport.name FROM person INNER JOIN sport ON person.id = sport.p_id;
```

## 141 Three table inner join, also 142
```sql
create database blockbuster;
\c blockbuster;

create table customers (cid serial primary key not null, cfirst char(50) not null);
create table movies (mid serial primary key not null, mmovie char(50) not null);
create table rentals (rid serial primary key not null, cid int references customers(cid), mid int references movies(mid));

insert into customers (cfirst) values ('James Bond'), ('Miss Moneypenny'), ('Q'), ('M'), ('Fleming');
insert into movies (mmovie) values ('Jaws'), ('Alien'), ('Never Say Never'), ('Skyfall'), ('Highlander');
insert into rentals (cid, mid) values (1,3), (2,5), (4,1), (3,2), (5,4), (3,2), (1,3), (2,4), (5,4), (2,1), (2,3), (4,5), (5,2), (2,1), (3,2), (3,3), (2,3), (1,4), (3,2), (2,3), (3,3), (2,4), (2,3), (1,2), (3,5), (3,4), (1,5);

select customers.cfirst, movies.mmovie from customers inner join rentals on customers.cid = rentals.cid inner join movies on rentals.mid = movies.mid;
```
- There are also a few more join examples in the video that are not typed out in the docs
- But if you wanna get good with these join shenanigans you have to practice a lot!

## 143 clauses
```sql
SELECT * FROM employees WHERE salary > 60000;
SELECT * FROM employees WHERE salary > 60000 AND score = 26;
SELECT * FROM employees WHERE score IN (25,26);
SELECT * FROM employees WHERE score NOT IN (25,26);
SELECT * FROM employees WHERE score BETWEEN 23 AND 26;
SELECT * FROM employees WHERE score IS NOT NULL;
SELECT * FROM employees WHERE name LIKE '%an%';
SELECT * FROM employees WHERE score <= 24 OR salary < 50000;
SELECT * FROM employees LIMIT 4;
SELECT * FROM employees ORDER BY id LIMIT 4;
```

## 144 update
```sql
UPDATE table
SET col1 = val1, col2 = val2, ..., colN = valN
WHERE <condition>;

UPDATE employees SET score = 99 WHERE ID = 3;
```

## 145 delete
```sql
DELETE FROM table WHERE <condition>;
DELETE FROM sport WHERE id = 6;
DELETE FROM sport; <- Careful, deletes all records!
```

## 146 users
```sql
select current_user;
\du
CREATE USER dummy WITH PASSWORD 'password';
GRANT ALL PRIVILEGES ON DATABASE company to dummy;
REVOKE ALL PRIVILEGES ON DATABASE company from dummy;
ALTER USER dummy WITH SUPERUSER;
ALTER USER dummy WITH NOSUPERUSER;
DROP USER dummy;
```

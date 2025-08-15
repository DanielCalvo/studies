## 13. Filtering Rows with "Where"
```sql
SELECT name, area FROM cities WHERE area > 4000; 
```

Internal order of execution in PgSQL, neat:
`SELECT name, area FROM cities WHERE area > 4000;` 
    Third           First              Second

## 14. More on the "Where" Keyword
- Wow, there are quite a few math operators available!
- `SELECT name, area FROM cities WHERE area = 8223;`
- `SELECT name, area FROM cities WHERE area != 8223;`

## 15. Compound "Where" Clauses
- `SELECT name, area FROM cities WHERE area BETWEEN 2000 AND 4000;`
- `SELECT name, area FROM cities WHERE name IN ('Delhi', 'Shanghai');`
- `SELECT name, area FROM cities WHERE name NOT IN ('Delhi', 'Shanghai');`
- `SELECT name, area FROM cities WHERE area NOT IN (3043, 8223) AND name = 'Delhi';` 
- `SELECT name, area FROM cities WHERE area NOT IN (3043, 8223) OR name = 'Delhi';` 

## 20. Calculations in "Where" Clauses
```sql
SELECT name, population / area AS population_density FROM cities WHERE population / area > 6000;
```

## 22. Updating Rows
```sql
UPDATE cities SET population = 39000000 WHERE name = 'Tokyo';
```

## 23. Deleting Rows
```sql
DELETE FROM cities WHERE name = 'Tokyo';
```
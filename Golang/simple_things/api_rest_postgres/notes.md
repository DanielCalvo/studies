Next actions
- Create the database schema for the book dataset
- Put the csv data inside the schema!

```sql
CREATE DATABASE startups;
\c startups;

CREATE TABLE company (
   ID             SERIAL PRIMARY KEY,
   NAME           TEXT    NOT NULL,
   VALUATION      REAL,
   DATEJOINED	  DATE DEFAULT '1900-01-01'
);
```

Fields:
- Company-Valuation ($B)
- Date Joined
- Countr 
- City
- Industry
- Select Inverstors 
- Founded Year 
- Total Raise 
- Financial Stage 
- Investors Count 
- Deal Terms 
- Portfolio Exits
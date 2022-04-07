## Ah
I was going to do a full CRUD api, but dealing with postgreSQL got kinda boring, there's not much of a challenge here

---

psql -h localhost -U postgres
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
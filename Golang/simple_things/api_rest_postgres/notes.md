## Ah
I was going to do a full CRUD api, but dealing with postgreSQL got kinda boring, there's not much of a challenge here
- How do you do pagination if someone wants a list of all the companies?
    - Can you return all the companies that start with the letter "H" or have a valuation bigger than X billion?
- How do you do authentication so that not anyone can delete companies?
- What if the DB stops responding?

Dang this rabbit hole goes deep

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

## Other ideas
- How do apply the "Cloud native Go" book to this?
- Can you implement some sort of authentication? 
- Everyone uses Postman these days, I should probably learn to use that instead of curl and browser...
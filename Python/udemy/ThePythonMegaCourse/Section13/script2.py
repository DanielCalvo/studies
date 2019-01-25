import psycopg2

#https://hub.docker.com/_/postgres
#docker run -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
#psql -h127.0.0.1 -Upostgres -W
#Password as specified above is mysecretpassword
#create database database1;

def create_table():
    conn = psycopg2.connect("dbname='database1' user='postgres' password='mysecretpassword' host='localhost' port='5432'")
    cur = conn.cursor()
    cur.execute("CREATE TABLE IF NOT EXISTS store (item TEXT, quantity INTEGER, price REAL)")
    cur.execute("INSERT INTO store VALUES ('Wine Glass', 8, 10.5)")
    conn.commit()
    conn.close()

def insert(item, quantity, price):
    conn = psycopg2.connect("dbname='database1' user='postgres' password='mysecretpassword' host='localhost' port='5432'")
    cur = conn.cursor()
    cur.execute("INSERT INTO store VALUES (%s,%s,%s)", (item, quantity, price))
    conn.commit()
    conn.close()

def view():
    conn = psycopg2.connect("dbname='database1' user='postgres' password='mysecretpassword' host='localhost' port='5432'")
    cur = conn.cursor()
    cur.execute("SELECT * FROM store")
    rows = cur.fetchall()
    conn.close()
    return rows

def delete(item):
    conn = psycopg2.connect("dbname='database1' user='postgres' password='mysecretpassword' host='localhost' port='5432'")
    cur = conn.cursor()
    cur.execute("DELETE FROM store WHERE item=%s", (item,))
    conn.commit()
    conn.close()

def update(quantity, price, item):
    conn = psycopg2.connect("dbname='database1' user='postgres' password='mysecretpassword' host='localhost' port='5432'")
    cur = conn.cursor()
    cur.execute("UPDATE store SET quantity=%s, price=%s WHERE item=%s",(quantity, price, item))
    conn.commit()
    conn.close()

# create_table()
# insert("Water Glass", 10, 5)
# delete("Wine Glass")
# update(11, 6, "Water Glass")
print(view())
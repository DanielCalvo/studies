import yaml
import MySQLdb

myfile = open("parameters.yaml", "r")
myvar = yaml.load(myfile)

db = MySQLdb.connect(host=myvar["parameters"]["database"]["host"],    # your host, usually localhost
                     user=myvar["parameters"]["database"]["user"],         # your username
                     passwd=myvar["parameters"]["database"]["password"],  # your password
                     db=myvar["parameters"]["database"]["name"])

print(myvar["parameters"]["database"]["host"])
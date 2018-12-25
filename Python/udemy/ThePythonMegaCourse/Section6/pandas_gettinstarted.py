
import pandas

df1 = pandas.DataFrame([[2,3,4],[10,20,30]],columns=["Price","Age","Value"],index=["First", "Second"])
#print(df1)

df2 = pandas.DataFrame([{"Name":"John","Surname":"Johns"},{"Name":"Jack"}])

#print(df2)

type(df2)
#print(dir(df2)) #Woah

print(df1.mean())
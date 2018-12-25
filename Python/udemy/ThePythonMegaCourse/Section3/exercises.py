# from datetime import datetime
#Lecture 3.64
# delta = datetime.now() - datetime(1900, 12, 31)
#
##Lecture 3.65
# print(delta.days)
# print(delta.seconds)
# print(datetime.now())
# now = datetime.now()
# print(now)
# then = datetime(1900, 12, 31, 20, 12, 59, 83845)
# print(then - now)
# whenever = datetime.strptime("2018-04-02", "%Y-%m-%d")
# print(whenever)
# print(whenever.strftime("%Y-%m-%d"))
#
##Lecture 3.66
# a = ['a','b','c']
# b = [1,2,3]
#
# for i, j in zip(a,b):
#     print(i,j)
#
#Lecture 3.67
# with open("example.txt", "w") as myfile:
#     myfile.write("Something!")

#The file is closed as soon as the "with" part is completed. Neat!

#Lecture 3.68
# temperatures = [10, -20, -289, 100]
# def c_to_f(c):
#     if c < -273.15:
#         return "That temperature doesn't make sense!"
#     else:
#         f = c* 9/5 + 32
#         return f
#
# for t in temperatures:
#     if type(c_to_f(t)) == float:
#         with open("c_to_f.txt", "a+") as myfile:
#             myfile.write(str(c_to_f(t)) + "\n")

#Lecture 3.70

import glob
from datetime import datetime
timestamp = datetime.now()

with open(timestamp.strftime("%Y-%m-%d-%H-%M-%S-%f"), "w") as timestamp_file:
    for myfile in glob.glob('file*.txt'):
        with open(myfile, "r") as open_file:
            for line in open_file.readlines():
                timestamp_file.write("Something!")



import cv2
import numpy
img = cv2.imread("smallgray.png",0)
#
# print(type(img))
# print(img)
#
# cv2.imwrite("newsmallgray.png", img)

#slicing arrays

#print(img)
#print(img.shape)
#print(img[0:2,2:4])

# for i in img:
# #     print(i)

# for i in img.T:
#     print(i)

# for i in img.flat:
#     print(i)

#ims = numpy.hstack((img, img, img))

ims = numpy.vstack((img, img, img))

#lst = numpy.hsplit(ims, 5)
lst = numpy.vsplit(ims, 3)

print(lst[0])

#Simple and crude program that takes images from one directory, resizes and saves them into another directory

from PIL import Image
import os

image_directory = 'E:\Daniel O RETORNO\Imagens\mobile pix'
resized_image_directory = image_directory + '\\resized'
image_list = [f for f in os.listdir(image_directory) if f.endswith('.jpg') or f.endswith('jpg')]

try:
    os.stat(resized_image_directory)
except FileNotFoundError:
    os.mkdir(resized_image_directory)

image_list.reverse()

for image_file in image_list:
    try:
        os.stat(resized_image_directory+'\\'+image_file)
        print(image_file+' is already resized. Skipping')
        continue
    except FileNotFoundError:
        pass
    try:
        im = Image.open(image_directory+'\\'+image_file)
    except OSError:
        print('Could not open'+image_directory+'\\'+image_file+'. Is it a valid image?')
        continue
    if im.size == ((4864, 2736)) or im.size == ((2592, 1456)):
        im_resized = im.resize((1920,1080))
        im_resized.save(resized_image_directory+'\\'+image_file)
        print(image_file+' was resized')
    if im.size == ((4048, 3036)):
        im_resized = im.resize((2080,1542))
        im_resized.save(resized_image_directory+'\\'+image_file)
        print(image_file+' was resized')
    if im.size == ((3036, 4048)):
        im_resized = im.resize((1542, 2080))
        im_resized.save(resized_image_directory+'\\'+image_file)
        print(image_file+' was resized')


# Maybe I'll get around finishing this later. This is a function to get rid of that horrendous copying and pasting up above
# def get_resized_size(x, y):
#
#     if x > y:
#         print("This is an image in landscape mode")
#         #4040, 3036
#
#     print(x / y)
#
#     if y > x:
#         print("This is an image in portrait mode")
#
# get_resized_size(4040, 3036)
#
# get_resized_size(1920, 1080)

#if image size is part of the sizes that we want to resize
    #resize the image
#else
    #do nothing

#What images size do you want to resize?



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
    if im.size == ((4864, 2736)) or im.size == ((2592, 1456)): #I'm only interested in resizing images that match these two sizes.
        im_resized = im.resize((1920,1080))
        im_resized.save(resized_image_directory+'\\'+image_file)
        print(image_file+' was resized')

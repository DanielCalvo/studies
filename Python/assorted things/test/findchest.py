import pyautogui
import subprocess
import logging
import coloredlogs
import random
import time


#Action list:
#1 Close everything
#2 Open backpack
#3 See what comes next but try to not think much about it right now

#Actually first of all maybe you should look for a chest

#How are you going to access the things inside the chest? Color? Position relative to chest location on UI bar?

#Let's access things by position and then have:
#Brown backpack: Food
#Grey backpack: Blank runes
#Blue backpack: Completed runes
#Green backpack: Life rings

#Oooooh pyautogui allows you

#Step 1: Close everything!
#Step 2: Open backpack
#Backpack cannot be of the same color as the rest of the backpacks (purple it is then!)
#Backpack must contain blank runes

#Step 3: Look for chest
#Step 4: Open chest

#Close everything before opening the chest!

def get_image_xy(image):
    image_pos = pyautogui.locateOnScreen(image)
    if image_pos is not None:
        image_pos_list = list(image_pos)
        image_x, image_y = image_pos_list[0], image_pos_list[1]
        print("Image " + image + " found at " + str(image_x) + " " + str(image_y))
        return image_x, image_y
    else:
        print("Image " + image + " not found on screen")
        return None, None

def get_image_list_xy(image):
    image_pos_box = pyautogui.locateAllOnScreen(image)
    image_pos_list = list(image_pos_box) ##I don't want a box, I want a list of (x, y)
    image_list_xy = []
    for i in image_pos_list:
        image_list_xy.append([i[0], i[1]])
    return image_list_xy

close_button_list = get_image_list_xy("images/close_button.png")

#The loop below doesn't work: When you close one menu, the others move and previous coordinates become invalid
for close_button in close_button_list:
    print(close_button)
    pyautogui.click(close_button[0], close_button[1], button='left', duration=random.uniform(0.1, 1.2))
    time.sleep(0.1)


#Try using the The locateCenterOnScreen() function for this line
character_x, character_y = 2330, 500

# chest_x, chest_y = get_image_xy("images/chest.png")
# if chest_x is not None and chest_y is not None:
#     print("Found chest!")
#     #Now what do we do again?
#     pyautogui.click(chest_x, chest_y, button='left', duration=random.uniform(0.1, 1.2))
#     exit(0)
# else:
#     print("Chest not found :(")
#
# pyautogui.click(character_x + 75, character_y, button='right', duration=random.uniform(0.1, 1.2))
# time.sleep(0.2)
# chest_x, chest_y = get_image_xy("images/chest.png")
# if chest_x is not None and chest_y is not None:
#     print("Found chest!")
#     exit(0)
# else:
#     print("Chest not found :(")
#
# pyautogui.click(character_x, character_y + 75, button='right', duration=random.uniform(0.1, 1.2))
# time.sleep(0.2)
# chest_x, chest_y = get_image_xy("images/chest.png")
# if chest_x is not None and chest_y is not None:
#     print("Found chest!")
#     exit(0)
# else:
#     print("Chest not found :(")
#
# pyautogui.click(character_x - 75, character_y, button='right', duration=random.uniform(0.1, 1.2))
# time.sleep(0.2)
# chest_x, chest_y = get_image_xy("images/chest.png")
# if chest_x is not None and chest_y is not None:
#     print("Found chest!")
#     exit(0)
# else:
#     print("Chest not found :(")
#
# pyautogui.click(character_x, character_y - 75, button='right', duration=random.uniform(0.1, 1.2))
# time.sleep(0.2)
# chest_x, chest_y = get_image_xy("images/chest.png")
# if chest_x is not None and chest_y is not None:
#     print("Found chest!")
#     exit(0)
# else:
#     print("Chest not found :(")


# if chest_x is not None and chest_y is not None:
#     print("Chest found!")
# else:
#     print("Chest not found :(")

import pyautogui
import time
import random
import logging
import coloredlogs


      #value_generated_session = 0
      #value_generated_session = rune['value']


runes = {
            'Great Fireball':  {'words': 'adori mas flam', 'manacost': 530, 'value': 57, 'number_made': 4},
            'Explosion': {'words': 'adevo mas hur', 'manacost': 570, 'value': 31, 'number_made': 6},
            'Sudden Death': {'words': 'adori gran mort', 'manacost': 570, 'value': 135, 'number_made': 3},
            }
counter = 0


im = pyautogui.screenshot()

life_ring_pos = pyautogui.locateOnScreen('life_ring.png')
empty_ring_slot_pos = pyautogui.locateOnScreen('empty_ring_slot.png')

if life_ring_pos is not None:
    life_ring_pos_list = list(life_ring_pos)
    life_ring_x, life_ring_y = life_ring_pos_list[0], life_ring_pos_list[1]
    print(life_ring_x, life_ring_y)
    print("Life ring found at") #If you make a list you can count how many life rings you have!
else:
    print("Life ring not found")

if empty_ring_slot_pos is not None:
    empty_ring_slot_pos_list = list(empty_ring_slot_pos)
    empty_ring_slot_x, empty_ring_slot_y = empty_ring_slot_pos_list[0], empty_ring_slot_pos_list[1]
    print(empty_ring_slot_x, empty_ring_slot_y)
else:
    print("Ring slot in use")

if life_ring_pos is not None and empty_ring_slot_pos is not None:
    pyautogui.click(life_ring_x, life_ring_y, button='left', duration=random.uniform(0.2, 1.2))
    pyautogui.dragTo(empty_ring_slot_x, empty_ring_slot_y, duration=random.uniform(0.2, 1.2), button='left')

    # pyautogui.dragTo()

# empty_ring_slot_pos = pyautogui.locateOnScreen('empty_ring_slot.png')
# if empty_ring_slot_pos is not None:
#     print(list(empty_ring_slot_pos))
# else:
#     print(empty_ring_slot_pos)

#pyautogui.click(3129, 503, button='left', duration=random.uniform(0.2, 1.2))

# rune = runes['Explosion']
#
#
# print(rune['value'])
# print(rune['manacost'])

# First slot of backpack on top most position: X: 3140, Y: 428

#This replaces the click for the eat function:

#This replaces the click for the click on the menubar when typing the rune atributes
#pyautogui.click(2040, 1063, button='left')

#pyautogui.click(2100 + (random.randint(-500, 500)), 1063 + (random.randint(-15, 15)), button='left', duration=random.uniform(0.2, 1.2))
#pyautogui.typewrite("adori mas flam", random.uniform(0.1, 0.2))

#Anything from 1600 to 2600

#average is 2100

#Improve this, randomize the sleep
# time.sleep(1)
# pyautogui.click()

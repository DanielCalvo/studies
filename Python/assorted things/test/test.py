import pyautogui
import time
import random
import logging
import coloredlogs

def create_rune(rune):
    logging.info("You called create rune with the following argument: " + str(rune))

    # First slot of backpack on top most position: X: 3140, Y: 428 -- 1770 455
    backpack_food_x = 1770
    backpack_food_y = 455

    pyautogui.click(backpack_food_x + random.randint(-9, 9), backpack_food_y + random.randint(-9, 9), button='right', duration=random.uniform(0.1, 1.2))
    time.sleep(random.uniform(0.05, 0.3))
    pyautogui.click(backpack_food_x + random.randint(-9, 9), backpack_food_y + random.randint(-9, 9), button='right', duration=random.uniform(0.1, 1.2))
    time.sleep(random.uniform(0.05, 0.3))
    pyautogui.click(backpack_food_x + random.randint(-9, 9), backpack_food_y + random.randint(-9, 9), button='right', duration=random.uniform(0.1, 1.2))
    time.sleep(random.uniform(0.05, 0.3))
    pyautogui.click(backpack_food_x + random.randint(-9, 9), backpack_food_y + random.randint(-9, 9), button='right', duration=random.uniform(0.1, 1.2))

    time.sleep(random.randint(1, 3))
    pyautogui.click(800 + (random.randint(-500, 500)), 1063 + (random.randint(-15, 15)),
                    button='left', duration=random.uniform(0.2, 1.2))

    pyautogui.typewrite(rune['words'], random.uniform(0.1, 0.2))
    pyautogui.typewrite('\n')
    logging.info("Rune created")


def put_life_ring_on(use_life_rings):
    if use_life_rings is False:
        return

    life_ring_pos = pyautogui.locateOnScreen('life_ring.png')
    empty_ring_slot_pos = pyautogui.locateOnScreen('empty_ring_slot.png')

    if life_ring_pos is not None:
        life_ring_pos_list = list(life_ring_pos)
        life_ring_x, life_ring_y = life_ring_pos_list[0], life_ring_pos_list[1]
        logging.info("Found a life ring at " + str(life_ring_x) + " " + str(life_ring_y))

        if empty_ring_slot_pos is not None:
            empty_ring_slot_pos_list = list(empty_ring_slot_pos)
            empty_ring_slot_x, empty_ring_slot_y = empty_ring_slot_pos_list[0], empty_ring_slot_pos_list[1]
            logging.info("Ring slot vacant at " + str(empty_ring_slot_x) + " " + str(
                empty_ring_slot_y) + ", moving life ring to ring slot")
            pyautogui.click(life_ring_x, life_ring_y, button='left', duration=random.uniform(0.2, 1.2))
            pyautogui.dragTo(empty_ring_slot_x, empty_ring_slot_y, duration=random.uniform(0.2, 1.2), button='left')
        else:
            logging.info("Ring slot already in use")
    else:
        logging.info("No life ring found")

def check_for_blank_rune():
    blank_rune_pos = pyautogui.locateOnScreen('blank_rune.png')
    if blank_rune_pos is None:
        exit_button_pos = pyautogui.locateOnScreen('exit_button.png') #Welp, I'm repeating myself a lot.
        exit_button_pos_list = list(exit_button_pos)
        exit_button_x, exit_button_y = exit_button_pos_list[0], exit_button_pos_list[1]
        pyautogui.click(exit_button_x, exit_button_y, button='left', duration=random.uniform(0.2, 1.2))

        exit_yes_pos = pyautogui.locateOnScreen('exit_yes.png')
        exit_yes_pos_list = list(exit_yes_pos)
        exit_yes_x, exit_yes_y = exit_yes_pos_list[0], exit_yes_pos_list[1]
        pyautogui.click(exit_yes_x, exit_yes_y, button='left', duration=random.uniform(0.2, 1.2))
        logging.critical("No blank rune found, exiting Tibia")
        logging.info("Exiting program. Good bye!")
        exit(1)

    else:
        logging.info("A blank rune was found")

runes = {
            'Great Fireball':  {'words': 'adori mas flam', 'manacost': 530, 'value': 57, 'number_made': 4},
            'Explosion': {'words': 'adevo mas hur', 'manacost': 570, 'value': 31, 'number_made': 6},
            'Sudden Death': {'words': 'adori gran mort', 'manacost': 0, 'value': 0, 'number_made': 0},
            'Thunderstorm': {'words': 'adori mas vis', 'manacost': 430, 'value': 47, 'number_made': 4},
            'Heavy Magic Missile': {'words': 'adori vis', 'manacost': 350, 'value': 12, 'number_made': 10},
        }

# # # # # # # # # # # # # #
rune = runes['Sudden Death'] #What if you make a typo?
use_life_rings = True
# # # # # # # # # # # # # #

logging.basicConfig(level=logging.INFO,format="%(asctime)s:%(levelname)s: %(message)s")
logging.basicConfig(level=logging.CRITICAL,format="%(asctime)s:%(levelname)s: %(message)s")
coloredlogs.install()

logging.info("Starting up, press CTRL-C to exit")

session_mana_spent = 0
session_value_generated = 0
manabar_full_counter = 0

try:
    while True:
        im = pyautogui.screenshot()
        manabar_color = im.getpixel((1163, 57)) #Find a way to calculate a manabar percentage maybe :o
        put_life_ring_on(use_life_rings)

        if manabar_color == (36, 37, 36):
            logging.info("Manabar empty")
            time.sleep(random.randint(1, 60))
            manabar_full_counter = 0

        elif manabar_color == (0, 70, 155):
            manabar_full_counter += 1
            if manabar_full_counter > 3:
                logging.critical("Manabar full more than 3 times in a row, mana is not being spent. Something is wrong. No blank rune?")
                exit(1)
            logging.info("Manabar full, counter: " + str(manabar_full_counter))
            create_rune(rune)
            session_value_generated = session_value_generated + (rune['value'] * rune['number_made'])
            session_mana_spent = session_mana_spent + rune['manacost']

            logging.info("session_mana_spent: " + str(session_mana_spent))
            logging.info("session_value_generated : " + str(session_value_generated))

            time.sleep(random.randint(1, 100))
        else:
            logging.info("Manabar did not match full or empty, Tibia not open or misaligned")
            time.sleep(60)
            manabar_full_counter = 0
except KeyboardInterrupt:
    logging.info("Caught KeyboardInterrupt, exiting")

#Problems to fix:
#If mana is full and you do the spell, and the mana continues to be full, you're out of blank runes.
#It only tries to eat if manabar == full

#maybe there should be a separate function called eat()?

#battleEye might detect rapid mouse movements, you need to make cursor movement look natural
#also, maybe randomize a bit the coordinates that you click on the chat bar and the food icons (that would be fun)

#Check for no food
#Check for no runes

#Nice to haves:
#Logging
#Statistics (mana used, runes made, food eaten, timeonline, rune types done etc)

#Create stats for: This session, today, all time.
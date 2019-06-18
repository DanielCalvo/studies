import pyautogui
import time
import random
import logging
import coloredlogs

def get_image_xy(image):
    image_pos = pyautogui.locateOnScreen(image)
    if image_pos is not None:
        image_pos_list = list(image_pos)
        image_x, image_y = image_pos_list[0], image_pos_list[1]
        logging.info("Image " + image + " found at " + str(image_x) + " " + str(image_y))
        return image_x, image_y
    else:
        logging.info("Image " + image + " not found on screen")
        return None, None

def create_rune(rune):
    logging.info("You called create rune with the following argument: " + str(rune))

    backpack_food_x, backpack_food_y = get_image_xy("images/ham.png")

    # First slot of backpack on top most position: X: 3140, Y: 428 -- 1770 455
    #backpack_food_x = 1770
    #backpack_food_y = 455

    pyautogui.click(backpack_food_x + random.randint(-5, 12), backpack_food_y + random.randint(-5, 12), button='right', duration=random.uniform(0.1, 1.2))
    time.sleep(random.uniform(0.05, 0.3))
    pyautogui.click(backpack_food_x + random.randint(-5, 12), backpack_food_y + random.randint(-5, 12), button='right', duration=random.uniform(0.1, 1.2))
    time.sleep(random.uniform(0.05, 0.3))
    pyautogui.click(backpack_food_x + random.randint(-5, 12), backpack_food_y + random.randint(-5, 12), button='right', duration=random.uniform(0.1, 1.2))
    time.sleep(random.uniform(0.05, 0.3))
    pyautogui.click(backpack_food_x + random.randint(-5, 12), backpack_food_y + random.randint(-5, 12), button='right', duration=random.uniform(0.1, 1.2))

    time.sleep(random.randint(1, 3))

    local_chat_x, local_chat_y = get_image_xy("images/local_chat.png")
    
    pyautogui.click(local_chat_x + (random.randint(-25, 25)), local_chat_y + (random.randint(-10, 50)),
                    button='left', duration=random.uniform(0.2, 1.2))

    pyautogui.typewrite(rune['words'], random.uniform(0.1, 0.2))
    pyautogui.typewrite('\n')
    logging.info("Rune created")


def put_life_ring_on(use_life_rings):
    if use_life_rings is False:
        return

    life_ring_pos = pyautogui.locateOnScreen('images/life_ring.png')
    empty_ring_slot_pos = pyautogui.locateOnScreen('images/empty_ring_slot.png')

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
    blank_rune_pos = pyautogui.locateOnScreen('images/blank_rune.png')
    if blank_rune_pos is None:
        exit_button_pos = pyautogui.locateOnScreen('images/exit_button.png') #Welp, I'm repeating myself a lot.
        exit_button_pos_list = list(exit_button_pos)
        exit_button_x, exit_button_y = exit_button_pos_list[0], exit_button_pos_list[1]
        pyautogui.click(exit_button_x, exit_button_y, button='left', duration=random.uniform(0.2, 1.2))

        exit_yes_pos = pyautogui.locateOnScreen('images/exit_yes.png')
        exit_yes_pos_list = list(exit_yes_pos)
        exit_yes_x, exit_yes_y = exit_yes_pos_list[0], exit_yes_pos_list[1]
        pyautogui.click(exit_yes_x, exit_yes_y, button='left', duration=random.uniform(0.2, 1.2))
        logging.critical("No blank rune found, exiting Tibia")
        logging.info("Exiting program. Good bye!")
        exit(1)

    else:
        logging.info("A blank rune was found")

runes = {
            'GFB':  {'words': 'adori mas flam', 'manacost': 530, 'value': 57, 'number_made': 4},
            'Explosion': {'words': 'adevo mas hur', 'manacost': 570, 'value': 31, 'number_made': 6},
            'SD': {'words': 'adori gran mort', 'manacost': 0, 'value': 0, 'number_made': 0},
            'Thunderstorm': {'words': 'adori mas vis', 'manacost': 430, 'value': 47, 'number_made': 4},
            'HMM': {'words': 'adori vis', 'manacost': 350, 'value': 12, 'number_made': 10},
        }

# # # # # # # # # # # # # #
rune = runes['GFB'] #What if you make a typo?
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
        #Removed all the manabar handling crap for the demo plz just work
        check_for_blank_rune()
        put_life_ring_on(True)
        create_rune(rune)
        session_value_generated = session_value_generated + (rune['value'] * rune['number_made'])
        session_mana_spent = session_mana_spent + rune['manacost']

        logging.info("session_mana_spent: " + str(session_mana_spent))
        logging.info("session_value_generated : " + str(session_value_generated))

        time.sleep(random.randint(1, 100))

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
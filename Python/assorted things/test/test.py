import pyautogui
import time
import random
import logging
import coloredlogs

def create_rune(rune):
    logging.info("You called create rune with the following argument: " + str(rune))

    # First slot of backpack on top most position: X: 3140, Y: 428

    pyautogui.click(3140 + random.randint(-9, 9), 428 + random.randint(-9, 9), button='right', duration=random.uniform(0.1, 1.2))
    time.sleep(random.uniform(0.05, 0.3))
    pyautogui.click(3140 + random.randint(-9, 9), 428 + random.randint(-9, 9), button='right', duration=random.uniform(0.1, 1.2))
    time.sleep(random.uniform(0.05, 0.3))
    pyautogui.click(3140 + random.randint(-9, 9), 428 + random.randint(-9, 9), button='right', duration=random.uniform(0.1, 1.2))
    time.sleep(random.uniform(0.05, 0.3))
    pyautogui.click(3140 + random.randint(-9, 9), 428 + random.randint(-9, 9), button='right', duration=random.uniform(0.1, 1.2))

    time.sleep(random.randint(1, 3))

    pyautogui.click(2100 + (random.randint(-500, 500)), 1063 + (random.randint(-15, 15)),
                    button='left', duration=random.uniform(0.2, 1.2))

    pyautogui.typewrite(rune['words'], random.uniform(0.1, 0.2))
    pyautogui.typewrite('\n')



    logging.info("Rune created")

runes = {
            'Great Fireball':  {'words': 'adori mas flam', 'manacost': 530, 'value': 57, 'number_made': 4},
            'Explosion': {'words': 'adevo mas hur', 'manacost': 570, 'value': 31, 'number_made': 6},
            'Sudden Death': {'words': 'adori gran mort', 'manacost': 570, 'value': 135, 'number_made': 3},
        }

# # # # # # # # # # # # # #
rune = runes['Great Fireball'] #What if you make a typo?
# # # # # # # # # # # # # #

logging.basicConfig(level=logging.INFO,format="%(asctime)s:%(levelname)s: %(message)s")
coloredlogs.install()

logging.info("Starting up, press CTRL-C to exit")

session_mana_spent = 0
session_value_generated = 0

try:
    while True:
        im = pyautogui.screenshot()
        manabar_color = im.getpixel((2650, 32))

        if manabar_color == (41, 41, 41):
            logging.info("Manabar empty")
            time.sleep(random.randint(1, 60))
        elif manabar_color == (0, 70, 155):
            logging.info("Manabar full")
            create_rune(rune)
            session_value_generated = session_value_generated + (rune['value'] * rune['number_made'])
            session_mana_spent = session_mana_spent + rune['manacost']

            logging.info("session_mana_spent: " + str(session_mana_spent))
            logging.info("session_value_generated : " + str(session_value_generated))

            time.sleep(random.randint(1, 100))
        else:
            logging.info("Manabar did not match full or empty, Tibia not open or misaligned")
            time.sleep(60)
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
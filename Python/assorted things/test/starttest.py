import pyautogui
import subprocess
import logging
import coloredlogs
import random
import time

#tibia_launcher_path = "/home/daniel/Downloads/Tibia/start-tibia-launcher.sh"
#subprocess.Popen(["/home/daniel/Downloads/Tibia/start-tibia-launcher.sh"])

logging.basicConfig(level=logging.INFO,format="%(asctime)s:%(levelname)s: %(message)s")
logging.basicConfig(level=logging.CRITICAL,format="%(asctime)s:%(levelname)s: %(message)s")
coloredlogs.install()

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


account_name_x, account_name_y = get_image_xy("images/account_name.png")
pyautogui.click(account_name_x + 200, account_name_y, button='left', duration=random.uniform(0.1, 1.2))
pyautogui.typewrite('123', random.uniform(0.1, 0.2))

password_x, password_y = get_image_xy("images/password.png")
pyautogui.click(password_x + 200, password_y, button='left', duration=random.uniform(0.1, 1.2))
pyautogui.typewrite('123456', random.uniform(0.1, 0.2))

login_x, login_y = get_image_xy("images/login.png")
pyautogui.click(login_x, login_y, button='left', duration=random.uniform(0.1, 1.2))

time.sleep(2)
character_x, character_y = get_image_xy("images/character.png")


pyautogui.click(character_x, character_y + 20, button='left', duration=random.uniform(0.1, 1.2))
time.sleep(0.05)
pyautogui.click(character_x, character_y + 20, button='left')

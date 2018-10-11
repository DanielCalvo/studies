from bs4 import BeautifulSoup
import requests

source = requests.get('https://rarbgmirror.xyz/torrents.php').text

soup = BeautifulSoup(source, 'lxml')

print(soup.prettify())
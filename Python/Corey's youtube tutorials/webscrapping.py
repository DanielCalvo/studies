
#Request library: Very popular!
#Uses beautifulsoup. lxml parser

#wget https://www.milanuncios.com/motos-de-segunda-mano-en-baleares/?desde=500

from fake_useragent import UserAgent
ua = UserAgent()
from bs4 import BeautifulSoup
import requests

headers = {ua.google}
print(headers)


#
#
#source = requests.get('https://www.milanuncios.com/motos-de-segunda-mano-en-baleares/?desde=500', ua.google).text
source = requests.get('https://www.milanuncios.com', ua.google).text
# soup = BeautifulSoup(source, 'lxml')
print(source)
print(source.sta)

#Using the file manually for now as milanuncios doesn't allow robots:

# with open('milanuncios.html') as html_file:
#     soup = BeautifulSoup(html_file, 'lxml')

#ad = soup.find('div', class_='aditem')
#print(ad.prettify())


# for ad in soup.find_all('div', class_='aditem'):
#
#     try:
#         price = ad.find('div', class_='aditem-price').text
#         print(price)
#     except AttributeError:
#         price = 'N/A'
#
#     try:
#         displacement = ad.find('div', class_='cc tag-mobile').text
#         print(displacement)
#     except AttributeError:
#         displacement = 'N/A'
#
#     try:
#         description = ad.find('div', class_='tx').text
#         print(description)
#     except AttributeError:
#         description = 'N/A'
#
#     try:
#         year = ad.find('div', class_='ano tag-mobile').text
#         print(year)
#     except AttributeError:
#         year = 'N/A'
#
#     try:
#         advert_id = ad.find('div', class_='x5').text.strip()
#         print(advert_id)
#     except AttributeError:
#         advert_id = 'N/A'
#
#     try:
#         location = ad.find('div', class_='x4').text
#         location = location.split(' ')
#         location = location[-1]
#         print(location)
#     except AttributeError:
#         location = 'N/A'
#
#     try:
#         time_since_last_renewed = ad.find('div', class_='x6').text
#         print(time_since_last_renewed)
#     except AttributeError:
#         time_since_last_renewed = 'N/A'


    #TODO: Try & catch so that the program doesn't crash when a given field is missing!
    #TODO: https://www.youtube.com/watch?v=ng2o98k983k&index=43&t=2305s&list=PL-osiE80TeTt2d9bfVyTiXJA-UTHn6WwU



#TODO: Kilometers. Can't do it on the current example as this guy doesn't have the mileage available on the Burgman.
#TODO: How do you handle exceptions when a given field isn't avaialble?


#I want to grab the link, here. <a class="aditem-detail-title" href="https://www.milanuncios.com/motos-de-carretera/suzuki-burgman-200-249087387.htm" target="_blank">


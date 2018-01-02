
#Request library: Very popular!
#Uses beautifulsoup. lxml parser

#wget https://www.milanuncios.com/motos-de-segunda-mano-en-baleares/?desde=500

from bs4 import BeautifulSoup
import requests

#source = requests.get('https://www.milanuncios.com/motos-de-segunda-mano-en-baleares/?desde=500').text
#soup = BeautifulSoup(source, 'lxml')

#Using the file manually for now as milanuncios doesn't allow robots:

with open('milanuncios.html') as html_file:
    soup = BeautifulSoup(html_file, 'lxml')

#ad = soup.find('div', class_='aditem')
#print(ad.prettify())


for ad in soup.find_all('div', class_='aditem'):

     price = ad.find('div', class_='aditem-price').text
     print(price)

     displacement = ad.find('div', class_='cc tag-mobile').text
     print(displacement)

     description = ad.find('div', class_='tx').text
     print(description)

    #TODO: Try & catch so that the program doesn't crash when a given field is missing!
    #TODO: https://www.youtube.com/watch?v=ng2o98k983k&index=43&t=2305s&list=PL-osiE80TeTt2d9bfVyTiXJA-UTHn6WwU

     year = ad.find('div', class_='ano tag-mobile').text
     print(year)


#
#     location = ad.find('div', class_='x4').text
#     location = location.split(' ')
#     location = location[-1]
#     print(location)
#
#     advert_id = ad.find('div', class_='x5').text.strip()
#     print(advert_id)
#
#     time_since_last_renewed = ad.find('div', class_='x6').text
#     print(time_since_last_renewed)
#
#

#TODO: Kilometers. Can't do it on the current example as this guy doesn't have the mileage available on the Burgman.
#TODO: How do you handle exceptions when a given field isn't avaialble?


#I want to grab the link, here. <a class="aditem-detail-title" href="https://www.milanuncios.com/motos-de-carretera/suzuki-burgman-200-249087387.htm" target="_blank">


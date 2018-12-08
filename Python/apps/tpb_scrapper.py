from bs4 import BeautifulSoup
import requests

tpb_main_url = 'https://thepiratebay.pet'
tbp_crawl_url = tpb_main_url + '/top/all'

source = requests.get(tbp_crawl_url).text
soup = BeautifulSoup(source, 'lxml')
torrent_url_list = []

print("Filtering " + tbp_crawl_url)

for article in soup.find_all('a'):
    if "/torrent/" in str(article) and str(tpb_main_url + article.get('href')) not in torrent_url_list:
        torrent_url_list.append(tpb_main_url + article.get('href'))


for torrent_url in torrent_url_list:
    torrent_url_source = requests.get(torrent_url).text
    torrent_url_soup = BeautifulSoup(torrent_url_source, 'lxml')

    for article in torrent_url_soup.find_all('a'):
        if "magnet:?" in str(article):
            print("transmission-remote -n \'transmission:transmission\' -a " + "\'" + article.get('href') + "\'")
            break
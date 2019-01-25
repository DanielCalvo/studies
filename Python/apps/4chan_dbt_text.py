
import json
from urllib.request import urlopen
from html.parser import HTMLParser
from termcolor import colored, cprint

class MyHTMLParser(HTMLParser):
    def handle_data(self, data):
       cprint(data,'green',end=' ')

def post_shortendate(post_now):
    split_date = post_now.split('(')
    result = '(' + split_date[1].rstrip()
    return result

def get_dbt_thread():
    response = urlopen('http://a.4cdn.org/o/catalog.json')
    catalog_json = json.loads(response.read())
    for pages in catalog_json:
        for threads in pages['threads']:
            try:
                if 'dbt' in threads['sub'] or 'dbt' in threads['com']: #TODO: [1]
                    dbt_thread_urlopen = urlopen('http://a.4cdn.org/o/thread/' + str(threads['no']) + '.json')
                    dbt_thread_data = json.loads(dbt_thread_urlopen.read())
                    return dbt_thread_data
            except KeyError:
                pass
def print_db_posts(dbt_thread):
    for posts in dbt_thread['posts']:
        try:
            cprint(post_shortendate(posts['now']), 'cyan', end=' ')
            parser.feed(posts['com'])
            print()
        except KeyError:
            pass

parser = MyHTMLParser()
dbt_thread = get_dbt_thread()
print_db_posts(dbt_thread)


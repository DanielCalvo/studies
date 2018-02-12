
import json
from urllib.request import urlopen
from html.parser import HTMLParser
import time

class MyHTMLParser(HTMLParser):
    def handle_data(self, data):
        print(data, end=' ') #TODO: Return data instead of printing it, print it later on the main program

def post_shortendate(post_now):
    split_date = post_now.split('(')
    result = '(' + split_date[1].rstrip()
    return result

parser = MyHTMLParser()
dbt_found = False
#TODO: If no initial json, save the catalog json
#TODO: Download the catalog every time you run the program. Display display which threads had replies, and how many (compared to the last stored json)
#TODO: Display thread and/or new replies on thread?
#TODO: Functions and stuff!
#TODO [1]: Improve matching as it could match any thread containing "dbt". Favour matching occurences of /dbt/ over /dbt/, as well as a match in the subject and comment line scoring higher than just one in either of them
#TODO: Do a function that gets the catalog
#TODO: Expand replies on the post that is replying to them (after the contents of post that is replying)
#TODO: Use pygments to colorize output!
#TODO: Handle bad responses from the API (appears to happen!)
#TODO: Adjust that try catch. Apparently, as soon as it catches the exception, it goes to the except block. Put the try/except only around the stuff that actually fails!

# Ran once to save the 4chan catalog to a file so we don't make API calls every time we run the program...

response = urlopen('http://a.4cdn.org/o/catalog.json')
catalog_json = json.loads(response.read())

for pages in catalog_json:
    for threads in pages['threads']:
        try:
            if 'dbt' in threads['sub'] or 'dbt' in threads['com']: #TODO: [1]
                dbt_found = True
                dbt_thread_urlopen = urlopen('http://a.4cdn.org/o/thread/' + str(threads['no']) + '.json')
                dbt_thread_data = json.loads(dbt_thread_urlopen.read())
                for posts in dbt_thread_data['posts']:
                    print(post_shortendate(posts['now']), end=' ')
                    parser.feed(posts['com'])
                    print()
        except KeyError:
            pass
        if dbt_found:
            break
    if dbt_found:
        break


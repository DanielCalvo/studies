
import json
from urllib.request import urlopen
from html.parser import HTMLParser

class MyHTMLParser(HTMLParser):
    def handle_data(self, data):
        print(data, end=' ') #TODO: Return data instead of printing it, print it later on the main program

def post_shortendate(post_now):
    split_date = post_now.split('(')
    result = '(' + split_date[1].rstrip()
    return result

parser = MyHTMLParser()

#TODO: If no initial json, save the catalog json
#TODO: Download the catalog every time you run the program. Display display which threads had replies, and how many (compared to the last stored json)
#TODO: Display thread and/or new replies on thread?
#TODO: Functions and stuff!
#TODO [1]: Improve matching as it could match any thread containing "dbt". Favour matching occurences of /dbt/ over /dbt/, as well as a match in the subject and comment line scoring higher than just one in either of them
#TODO: Do a function that gets the catalog
#TODO: Expand replies on the post that is replying to them (after the contents of post that is replying)
#TODO: Use pygments to colorize output!

# Ran once to save the 4chan catalog to a file so we don't make API calls every time we run the program...

#response = urlopen('http://a.4cdn.org/o/catalog.json')
#data = json.loads(response.read())
#fh = open('4chan_catalog.json', 'w')
#json.dump(data, fh, indent=2)

catalog_file = open('4chan_catalog.json', 'r')
catalog_json = json.load(catalog_file)
#print(json.dumps(catalog_json, indent=2))

for pages in catalog_json:
    for threads in pages['threads']:
        try:
            if 'dbt' in threads['sub'] or 'dbt' in threads['com']: #TODO: [1]
                dbt_thread_urlopen = urlopen('http://a.4cdn.org/o/thread/' + str(threads['no']) + '.json')
                dbt_thread_data = json.loads(dbt_thread_urlopen.read())
                dbt_thread_fh = open(str(threads['no']) + '.json', 'w')
                json.dump(dbt_thread_data, dbt_thread_fh, indent=2)
        except KeyError:
            pass

for posts in dbt_thread_data['posts']:
    # print(posts['no'], ' - ', posts['now'], parser.feed(html.unescape(posts['com'])))
    print(posts['no'], post_shortendate(posts['now']), end=' ')
    try:
        parser.feed(posts['com'])
        print()
    except KeyError:
        pass

    #print()



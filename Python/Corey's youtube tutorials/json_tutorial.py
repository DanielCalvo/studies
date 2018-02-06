
import json
from urllib.request import urlopen

#TODO: If no initial json, save the catalog json
#TODO: Download the catalog every time you run the program. Display display which threads had replies, and how many (compared to the last stored json)


# Ran once to save the 4chan catalog to a file so we don't make API calls every time we run the program...
response = urlopen('http://a.4cdn.org/o/catalog.json')
data = json.loads(response.read())

fh = open('4chan_catalog.json', 'w')
json.dump(data, fh, indent=2)

# with open('4chan_catalog.json') as f:
#      data = json.loads(f)
#
# print(json.dumps(data, ident=4))
#
#
# for page in data:
#     for thread in page['threads']:
#         print(thread['no'], thread['replies'])
#     break
#
#
#
# data = json.loads(response.read())
# #print(json.dumps(data, indent=2))
#
# print(data[0]['threads'])
#
# for threadnum in data[0]['threads']:
#     print(threadnum)
#
#
# #for page in data:
# #    print(page)
#
#
#
# #https://www.youtube.com/watch?v=9N6a-VLBa2I
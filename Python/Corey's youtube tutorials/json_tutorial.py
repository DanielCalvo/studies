
import json
from urllib.request import urlopen

#
# #
# with open('states.json') as f:
#     data = json.load(f)
# #
# for state in data['states']:
#     del state['area_codes']
#
# with open('new_states.json', 'w') as f1:
#     json.dump(data, f1, indent=2)

response = urlopen('http://a.4cdn.org/n/catalog.json')

data = json.loads(response.read())
print(json.dumps(data, indent=2))

for thread in data['threads']:
    print(thread)

#https://www.youtube.com/watch?v=9N6a-VLBa2I
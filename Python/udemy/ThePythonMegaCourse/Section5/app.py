import json
from difflib import get_close_matches

data = json.load(open("data.json", "r"))

def translate(word):

    # if word[0].isupper(): #If word starts with uppercase
    #     word = word[0] + word[1:].lower()
    #     print("Word is:", word)
    # else:
    #     word = word.lower()
    #     print("Word is:", word)

    word = word.lower()
    if word in data:
        return data[word]
    elif word.title() in data:
        return data[word.title()]
    elif word.upper() in data:
        return data[word.upper()]
    else:
        word_aprox = get_close_matches(word, data.keys(), cutoff=0.8)
        if len(word_aprox) > 0:
            print(word,"is not present in dictionary. Assuming to be a typo of",  word_aprox[0])
            return data[word_aprox[0]]
        else:
            return "The word %s does not exist on the dictionary or cannot be assumed to be a typo" % word

word = input("Enter word: ")

output = translate(word)

if type(output) == list:
    for item in output:
        print(item)
else:
    print(output)
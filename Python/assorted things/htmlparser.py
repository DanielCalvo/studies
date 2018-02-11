from html.parser import HTMLParser

class MyHTMLParser(HTMLParser):


    def handle_data(self, data):
        #print("Data     :", data)
        #print("Data:", data)
        print(data)
        if data.endswith('\n'):
            print("yeah")



parser = MyHTMLParser()


mystring = '02/09/18(Fri)16:42:14'



def post_shortendate(post_now):
    split_date = post_now.split('(')
    result = '(' + split_date[1]
    return result

print(post_shortendate(mystring))
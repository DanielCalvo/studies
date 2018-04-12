
while True:
    reply = input("Enter text: ")
    if reply == 'stop': break
    try:
        print (float(reply) ** 2)
    except:
        print ("Bad!")


import http.server
import socketserver
from http import HTTPStatus
import socket
import os
import datetime
import yaml

parameters_yaml = open("parameters.yaml", "r")
parameters = yaml.load(parameters_yaml)

hostname = socket.gethostname()
git_head = os.popen("cat .git/HEAD").read()

mystring = "Running on: " + hostname + "\nAt branch: " + git_head + "Deployed at time: " + str(datetime.datetime.now())
print(mystring)
mybytes = bytes(mystring, 'utf-8')


class Handler(http.server.SimpleHTTPRequestHandler):
    def do_GET(self):
        self.send_response(HTTPStatus.OK)
        self.end_headers()
        self.wfile.write(mybytes)


httpd = socketserver.TCPServer(('', 80), Handler)
httpd.serve_forever()

#Needs to connect to the database
#Needs to have a parameter file to connect to a database and put logs in place
#Needs to write out logs!
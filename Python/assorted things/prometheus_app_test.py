import http.server
import prometheus_client
from io import BytesIO
import time
import random

HELLO_REQUESTS = prometheus_client.Counter("hello_world_total","Number of hello worlds requested")
SALES = prometheus_client.Counter("hello_world_sales_euro_total","Number of hello worlds requested")
LATENCY = prometheus_client.Summary("hello_world_latency_seconds","Time it takes for a hello world to happen")

class MyHandler(http.server.BaseHTTPRequestHandler):

    def do_GET(self):
        start = time.time()
        euros = random.random()
        SALES.inc(euros)
        HELLO_REQUESTS.inc()
        self.send_response(200)
        self.end_headers()
        self.wfile.write(b'Hello, world!')
        LATENCY.observe(time.time() - start)

    def do_POST(self):
        content_length = int(self.headers['Content-Length'])
        body = self.rfile.read(content_length)
        self.send_response(200)
        self.end_headers()
        response = BytesIO()
        response.write(b'This is POST request. ')
        response.write(b'Received: ')
        response.write(body)
        self.wfile.write(response.getvalue())


httpd = http.server.HTTPServer(('localhost', 8000), MyHandler)
prometheus_client.start_http_server(8001)
httpd.serve_forever()


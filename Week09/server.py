from http.server import HTTPServer, BaseHTTPRequestHandler

class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        # Set response status code to 200 (OK)
        self.send_response(200)
        # Set the content-type header
        self.send_header("Content-type", "text/plain")
        self.end_headers()
        # Write the response body
        self.wfile.write(b"Hello, world!")

def run(server_class=HTTPServer, handler_class=SimpleHTTPRequestHandler):
    # Define server address and port
    server_address = ('0.0.0.0', 5002)
    httpd = server_class(server_address, handler_class)
    print("Serving on http://localhost:5002")
    httpd.serve_forever()

if __name__ == "__main__":
    run()

from http.server import SimpleHTTPRequestHandler, ThreadingHTTPServer
import json
import os
import sys
import urllib.error
import urllib.request


ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), ".."))
UPSTREAM_URL = os.environ.get("CKEY_CHAT_COMPLETIONS_URL", "https://ckey.vn/v1/chat/completions")


class Handler(SimpleHTTPRequestHandler):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, directory=ROOT, **kwargs)

    def end_headers(self):
        self.send_header("Access-Control-Allow-Origin", "*")
        self.send_header("Access-Control-Allow-Headers", "Content-Type, Authorization")
        self.send_header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        super().end_headers()

    def do_OPTIONS(self):
        self.send_response(204)
        self.end_headers()

    def do_POST(self):
        if self.path != "/api/chat/completions":
            self.send_json(404, {"error": {"message": "Not found"}})
            return

        api_key = os.environ.get("CKEY_API_KEY")
        if not api_key:
            self.send_json(500, {"error": {"message": "Missing CKEY_API_KEY environment variable"}})
            return

        length = int(self.headers.get("Content-Length", "0"))
        body = self.rfile.read(length)

        request = urllib.request.Request(
            UPSTREAM_URL,
            data=body,
            method="POST",
            headers={
                "Authorization": f"Bearer {api_key}",
                "Content-Type": "application/json",
                "Accept": "application/json",
                "User-Agent": "HuyenLichTarotPrototype/1.0",
            },
        )

        try:
            with urllib.request.urlopen(request, timeout=120) as response:
                payload = response.read()
                self.send_response(response.status)
                self.send_header("Content-Type", response.headers.get("Content-Type", "application/json"))
                self.send_header("Content-Length", str(len(payload)))
                self.end_headers()
                self.wfile.write(payload)
        except urllib.error.HTTPError as error:
            payload = error.read() or json.dumps({"error": {"message": str(error)}}).encode("utf-8")
            self.send_response(error.code)
            self.send_header("Content-Type", "application/json")
            self.send_header("Content-Length", str(len(payload)))
            self.end_headers()
            self.wfile.write(payload)
        except Exception as error:
            self.send_json(502, {"error": {"message": str(error)}})

    def send_json(self, status, payload):
        data = json.dumps(payload).encode("utf-8")
        self.send_response(status)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(data)))
        self.end_headers()
        self.wfile.write(data)


def main():
    port = int(sys.argv[1]) if len(sys.argv) > 1 else 5175
    server = ThreadingHTTPServer(("127.0.0.1", port), Handler)
    print(f"Tarot prototype server: http://127.0.0.1:{port}/prototype/tarot-flow.html", flush=True)
    server.serve_forever()


if __name__ == "__main__":
    main()

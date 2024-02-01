curl -X POST "https://bvg-apps-ext.hafas.de/bin/mgate.exe" \
  -H "Content-Type: application/json" \
  -H "Accept-Encoding: gzip, br, deflate" \
  -H "Accept: application/json" \
  -H "user-agent: YourUserAgentString" \
  -H "connection: close" \
  --data '{
    "lang": "en",
    "svcReqL": [
      {
        "meth": "ServerInfo",
        "req": {
          "getVersionInfo": "true"
        }
      }
    ],
    "client": {
      "type": "IPA",
      "id": "BVG",
      "v": "6020000",
      "name": "FahrInfo"
    },
    "ext": "BVG.1",
    "ver": "1.44",
    "auth": {
      "type": "AID",
      "aid": "YoJ05NartnanEGCj"
    }
  }' --compressed | jq .
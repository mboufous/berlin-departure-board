# find a way to get less verbosity

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
        "cfg": { "polyEnc": "GPA" },
        "meth": "LocMatch",
        "req": {
          "input": {
            "loc": {
              "type": "ALL",
              "name": "reinickendorfer?"
            },
            "maxLoc": 5,
            "field": "S"
            }
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

  {
  "ver": "1.44",
  "ext": "BVG.1",
  "lang": "eng",
  "id": "edi2zzaakga5py8g",
  "err": "OK",
  "graph": {
    "id": "standard",
    "index": 0
  },
  "subGraph": {
    "id": "global",
    "index": 0
  },
  "view": {
    "id": "standard",
    "index": 0,
    "type": "WGS84"
  },
  "svcResL": [
    {
      "meth": "LocMatch",
      "err": "OK",
      "res": {
        "common": {
          "prodL": [
            {
              "name": "U6",
              "nameS": "U6",
              "icoX": 0,
              "cls": 2,
              "prodCtx": {
                "lineId": "U6"
              }
            },
            {
              "name": "120",
              "nameS": "120",
              "icoX": 1,
              "cls": 8,
              "prodCtx": {
                "lineId": "120"
              }
            },
            {
              "name": "147",
              "nameS": "147",
              "icoX": 1,
              "cls": 8,
              "prodCtx": {
                "lineId": "147"
              }
            },
            {
              "name": "M27",
              "nameS": "M27",
              "icoX": 1,
              "cls": 8,
              "prodCtx": {
                "lineId": "M27"
              }
            },
            {
              "name": "N6",
              "nameS": "N6",
              "icoX": 1,
              "cls": 8,
              "prodCtx": {
                "lineId": "N6"
              }
            },
            {
              "name": "N20",
              "nameS": "N20",
              "icoX": 1,
              "cls": 8,
              "prodCtx": {
                "lineId": "N20"
              }
            },
            {
              "name": "808",
              "nameS": "808",
              "icoX": 1,
              "cls": 8,
              "prodCtx": {
                "lineId": "808"
              }
            },
            {
              "name": "809",
              "nameS": "809",
              "icoX": 1,
              "cls": 8,
              "prodCtx": {
                "lineId": "809"
              }
            }
          ],
          "icoL": [
            {
              "res": "prod_sub_t",
              "fg": {
                "r": 255,
                "g": 255,
                "b": 255
              },
              "bg": {
                "r": 0,
                "g": 51,
                "b": 153
              }
            },
            {
              "res": "prod_bus_t",
              "fg": {
                "r": 255,
                "g": 255,
                "b": 255
              },
              "bg": {
                "r": 153,
                "g": 51,
                "b": 153
              }
            },
            {
              "res": "loc_addr"
            }
          ]
        },
        "match": {
          "field": "S",
          "state": "L",
          "locL": [
            {
              "lid": "A=1@O=U Reinickendorfer Str. (Berlin)@X=13370385@Y=52539891@U=86@L=900008102@B=1@p=1697107651@",
              "type": "S",
              "name": "U Reinickendorfer Str. (Berlin)",
              "icoX": 0,
              "extId": "900008102",
              "state": "F",
              "crd": {
                "x": 13370385,
                "y": 52539891,
                "floor": 0
              },
              "pCls": 10,
              "pRefL": [
                0,
                1,
                2,
                3,
                4,
                5
              ],
              "wt": 12490,
              "gidL": [
                "A×de:11000:900008102"
              ],
              "TZOffset": 120,
              "chgTime": "000300"
            },
            {
              "lid": "A=1@O=U Reinickendorfer Str./Sellerstr. (Berlin)@X=13370385@Y=52539891@U=86@L=900008107@B=1@p=1697107651@",
              "type": "S",
              "name": "U Reinickendorfer Str./Sellerstr. (Berlin)",
              "icoX": 1,
              "extId": "900008107",
              "state": "F",
              "crd": {
                "x": 13370385,
                "y": 52539891,
                "floor": 0
              },
              "pCls": 8,
              "pRefL": [
                1,
                4,
                5
              ],
              "wt": 923,
              "TZOffset": 120,
              "chgTime": "000300"
            },
            {
              "lid": "A=1@O=Hennigsdorf, Reinickendorfer Str.@X=13197180@Y=52655097@U=86@L=900203382@B=1@p=1697107651@",
              "type": "S",
              "name": "Hennigsdorf, Reinickendorfer Str.",
              "icoX": 1,
              "extId": "900203382",
              "state": "F",
              "crd": {
                "x": 13196929,
                "y": 52655088,
                "floor": 0
              },
              "pCls": 8,
              "pRefL": [
                6,
                7
              ],
              "wt": 159,
              "gidL": [
                "A×de:12065:900203382"
              ],
              "TZOffset": 120,
              "chgTime": "000300"
            },
            {
              "lid": "A=1@O=U Reinickendorfer Str./Fennstr. (Berlin)@X=13368713@Y=52541195@U=86@L=900008104@B=1@p=1697107651@",
              "type": "S",
              "name": "U Reinickendorfer Str./Fennstr. (Berlin)",
              "icoX": 0,
              "extId": "900008104",
              "state": "F",
              "crd": {
                "x": 13368713,
                "y": 52541195,
                "floor": 0
              },
              "pCls": 10,
              "pRefL": [
                0,
                1,
                2,
                3,
                4,
                5
              ],
              "gidL": [
                "A×de:11000:900008104"
              ],
              "TZOffset": 120,
              "chgTime": "000300"
            },
            {
              "lid": "A=2@O=13347 Berlin-Gesundbrunnen, Reinickendorfer Str.@X=13370106@Y=52540979@U=103@b=770005099@B=1@p=1592347662@",
              "type": "A",
              "name": "13347 Berlin-Gesundbrunnen, Reinickendorfer Str.",
              "icoX": 2,
              "state": "M",
              "crd": {
                "x": 13370106,
                "y": 52540979
              }
            }
          ]
        }
      }
    }
  ]
}
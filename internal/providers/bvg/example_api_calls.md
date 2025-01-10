## Server status

Request:
```bash
curl -X POST "https://bvg-apps-ext.hafas.de/bin/mgate.exe" \
  -H "Content-Type: application/json" \
  -H "Accept-Encoding: gzip, br, deflate" \
  -H "Accept: application/json" \
  -H "User-Agent: YourUserAgentString" \
  -H "Connection: close" \
  --data-raw '{
      "lang": "en",
      "svcReqL": [
            {
                "meth": "ServerInfo",
                "req": {
                }
            }
        ],
      "client": {
            "type": "WEB",
            "l":"vs_webapp",
            "id": "VBB",
            "v": 10003,
            "name": "webapp"
        },
      "ext": "BVG.1",
      "ver": "1.72",
      "auth": {
          "type": "AID",
          "aid": "dVg4TZbW8anjx9ztPwe2uk4LVRi9wO"
      }
  }' \
  --compressed | jq .
```
<details>
  <summary>Response</summary>

```json
{
    "ver": "1.72",
    "ext": "BVG.1",
    "lang": "eng",
    "id": "63icujaxm8kpym4k",
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
        "meth": "ServerInfo",
        "err": "OK",
        "res": {
            "fpB": "20241215",
            "fpE": "20251213",
            "sD": "20241222",
            "sT": "170138",
            "enc": "UNKNOWN",
            "planrtTS": "1734883267"
      }
    }
  ]
}
```
</details>

## Get Station details

Request:
```bash
curl -X POST "https://bvg-apps-ext.hafas.de/bin/mgate.exe" \
  -H "Content-Type: application/json" \
  -H "Accept-Encoding: gzip, br, deflate" \
  -H "Accept: application/json" \
  -H "User-Agent: YourUserAgentString" \
  -H "Connection: close" \
  --data-raw '{
      "lang": "en",
      "svcReqL": [
            {
                "meth": "LocDetails",
                "req": {
                  "getHIM":false,
                  "getProducts":true,
                  "getIcons":false,
                  "getAttributes":true,
                  "getTariff":false,
                  "locL": [
                    {
                      "type": "S",
                      "lid": "A=1@L=900008102@"
                    }
                  ]
                }
            }
        ],
      "client": {
            "type": "WEB",
            "l":"vs_webapp",
            "id": "VBB",
            "v": 10003,
            "name": "webapp"
        },
      "ext": "BVG.1",
      "ver": "1.72",
      "auth": {
          "type": "AID",
          "aid": "dVg4TZbW8anjx9ztPwe2uk4LVRi9wO"
      }
  }' \
  --compressed | jq .
```

<details>
    <summary>Response</summary>
    ```json
    {
      "ver": "1.72",
      "ext": "BVG.1",
      "lang": "eng",
      "id": "s3gqkhxw4wud9m8x",
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
          "meth": "LocDetails",
          "err": "OK",
          "res": {
            "common": {
              "locL": [
                {
                  "lid": "A=1@O=U Reinickendorfer Str./Fennstr. (Berlin)@X=13368713@Y=52541195@U=86@L=900008104@",
                  "type": "S",
                  "name": "U Reinickendorfer Str./Fennstr. (Berlin)",
                  "icoX": 2,
                  "extId": "900008104",
                  "state": "F",
                  "crd": {
                    "x": 13368713,
                    "y": 52541195,
                    "floor": 0
                  },
                  "pCls": 0,
                  "globalIdL": [
                    {
                      "id": "de:11000:900008104",
                      "type": "A"
                    }
                  ],
                  "TZOffset": 60,
                  "chgTime": "000300",
                  "isFavrbl": true,
                  "isHstrbl": true,
                  "countryCodeL": [
                    "de"
                  ]
                }
              ],
              "prodL": [
                {
                  "pid": "L::1::U::B3041857784::de:VBB:11000000|U-Bahn|U6:::*",
                  "name": "U6",
                  "nameS": "U6",
                  "icoX": 0,
                  "cls": 2,
                  "prodCtx": {
                    "name": "U6",
                    "line": "U6",
                    "lineId": "de:VBB:11000000|U-Bahn|U6:",
                    "catOut": "U       ",
                    "catOutS": "U",
                    "catOutL": "U       "
                  }
                },
                {
                  "pid": "L::3::Bus::B3041857784::de:VBB:11000000|Bus|120:::*",
                  "name": "120",
                  "nameS": "120",
                  "icoX": 1,
                  "cls": 8,
                  "prodCtx": {
                    "name": "120",
                    "line": "120",
                    "lineId": "de:VBB:11000000|Bus|120:",
                    "catOut": "Bus     ",
                    "catOutS": "B",
                    "catOutL": "Bus     "
                  }
                },
                {
                  "pid": "L::3::Bus::B3041857784::de:VBB:11000000|Bus|147:::*",
                  "name": "147",
                  "nameS": "147",
                  "icoX": 1,
                  "cls": 8,
                  "prodCtx": {
                    "name": "147",
                    "line": "147",
                    "lineId": "de:VBB:11000000|Bus|147:",
                    "catOut": "Bus     ",
                    "catOutS": "B",
                    "catOutL": "Bus     "
                  },
                  "himIdL": [
                    "HIM_FREETEXT_222451",
                    "HIM_FREETEXT_248518"
                  ]
                },
                {
                  "pid": "L::3::Bus::B3041857784::de:VBB:11000000|Bus|M27:::*",
                  "name": "M27",
                  "nameS": "M27",
                  "icoX": 1,
                  "cls": 8,
                  "prodCtx": {
                    "name": "M27",
                    "line": "M27",
                    "lineId": "de:VBB:11000000|Bus|M27:",
                    "catOut": "Bus     ",
                    "catOutS": "B",
                    "catOutL": "Bus     "
                  },
                  "himIdL": [
                    "HIM_FREETEXT_248771",
                    "HIM_FREETEXT_248773"
                  ]
                },
                {
                  "pid": "L::3::Bus::B3041857784::de:VBB:11000000|Bus|N6:::*",
                  "name": "N6",
                  "nameS": "N6",
                  "icoX": 1,
                  "cls": 8,
                  "prodCtx": {
                    "name": "N6",
                    "line": "N6",
                    "lineId": "de:VBB:11000000|Bus|N6:",
                    "catOut": "Bus     ",
                    "catOutS": "B",
                    "catOutL": "Bus     "
                  },
                  "himIdL": [
                    "HIM_FREETEXT_222408",
                    "HIM_FREETEXT_238333"
                  ]
                },
                {
                  "pid": "L::3::Bus::B3041857784::de:VBB:11000000|Bus|N20:::*",
                  "name": "N20",
                  "nameS": "N20",
                  "icoX": 1,
                  "cls": 8,
                  "prodCtx": {
                    "name": "N20",
                    "line": "N20",
                    "lineId": "de:VBB:11000000|Bus|N20:",
                    "catOut": "Bus     ",
                    "catOutS": "B",
                    "catOutL": "Bus     "
                  },
                  "himIdL": [
                    "HIM_FREETEXT_253181"
                  ]
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
                  "res": "loc_stop"
                }
              ]
            },
            "locL": [
              {
                "lid": "A=1@O=U Reinickendorfer Str. (Berlin)@X=13370385@Y=52539891@U=86@L=900008102@",
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
                "wt": 14829,
                "stopLocL": [
                  0
                ],
                "globalIdL": [
                  {
                    "id": "de:11000:900008102",
                    "type": "A"
                  }
                ],
                "TZOffset": 60,
                "chgTime": "000300",
                "isFavrbl": true,
                "isHstrbl": true,
                "countryCodeL": [
                  "de"
                ]
              }
            ]
          }
        }
      ]
    }
    ```
</details>

## Get Departure Board

Request:
```bash
curl -X POST "https://bvg-apps-ext.hafas.de/bin/mgate.exe" \
  -H "Content-Type: application/json" \
  -H "Accept-Encoding: gzip, br, deflate" \
  -H "Accept: application/json" \
  -H "User-Agent: YourUserAgentString" \
  -H "Connection: close" \
  --data-raw '{
      "lang": "en",
      "svcReqL": [
            {
                "meth": "StationBoard",
                "req": {
                  "jnyFltrL":[
                      {
                        "type":"PROD",
                        "mode":"INC",
                        "value":127
                      }
                  ],
                  "type": "DEP",
                  "sort": "PT",
                  "maxJny": 4,
                  "date": "20250109",
                  "time": "214038",
                  "stbLoc": {
                    "type": "S",
                    "lid": "A=1@L=900008102@"
                  }
                }
            }
        ],
      "client": {
            "type": "WEB",
            "l":"vs_webapp",
            "id": "VBB",
            "v": 10003,
            "name": "webapp"
        },
      "ext": "BVG.1",
      "ver": "1.72",
      "auth": {
          "type": "AID",
          "aid": "dVg4TZbW8anjx9ztPwe2uk4LVRi9wO"
      }
  }' \
  --compressed | jq .
```

<details>
    <summary>Response</summary>
    ```json
  {
    "ver": "1.72",
    "ext": "BVG.1",
    "lang": "eng",
    "id": "fh2xqh2gk8wiykcs",
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
        "meth": "StationBoard",
        "err": "OK",
        "res": {
          "common": {
            "locL": [
              {
                "lid": "A=1@O=U Reinickendorfer Str. (Berlin)@X=13370385@Y=52539891@U=86@L=900008102@",
                "type": "S",
                "name": "U Reinickendorfer Str. (Berlin)",
                "icoX": 2,
                "extId": "900008102",
                "state": "F",
                "crd": {
                  "x": 13370385,
                  "y": 52539891,
                  "floor": 0
                },
                "pCls": 10,
                "pRefL": [
                  1,
                  2,
                  3,
                  4,
                  5,
                  6
                ],
                "globalIdL": [
                  {
                    "id": "de:11000:900008102",
                    "type": "A"
                  }
                ],
                "chgTime": "000300",
                "countryCodeL": [
                  "de"
                ]
              },
              {
                "lid": "A=1@O=U Leopoldplatz (Berlin)@X=13359391@Y=52546489@U=86@L=900009102@",
                "type": "S",
                "name": "U Leopoldplatz (Berlin)",
                "icoX": 2,
                "extId": "900009102",
                "state": "F",
                "crd": {
                  "x": 13359391,
                  "y": 52546489,
                  "floor": 0
                },
                "pCls": 10,
                "globalIdL": [
                  {
                    "id": "de:11000:900009102",
                    "type": "A"
                  }
                ],
                "chgTime": "000300",
                "countryCodeL": [
                  "de"
                ]
              },
              {
                "lid": "A=1@O=U Kurt-Schumacher-Platz (Berlin)@X=13327326@Y=52563479@U=86@L=900086102@",
                "type": "S",
                "name": "U Kurt-Schumacher-Platz (Berlin)",
                "icoX": 2,
                "extId": "900086102",
                "state": "F",
                "crd": {
                  "x": 13327326,
                  "y": 52563479,
                  "floor": 0
                },
                "pCls": 10,
                "globalIdL": [
                  {
                    "id": "de:11000:900086102",
                    "type": "A"
                  }
                ],
                "chgTime": "000300",
                "countryCodeL": [
                  "de"
                ]
              },
              {
                "lid": "A=1@O=Hainbuchenstr. (Berlin)@X=13274820@Y=52634251@U=86@L=900092202@",
                "type": "S",
                "name": "Hainbuchenstr. (Berlin)",
                "icoX": 0,
                "extId": "900092202",
                "state": "F",
                "crd": {
                  "x": 13274820,
                  "y": 52634251,
                  "floor": 0
                },
                "pCls": 8,
                "globalIdL": [
                  {
                    "id": "de:11000:900092202",
                    "type": "A"
                  }
                ],
                "chgTime": "000300",
                "countryCodeL": [
                  "de"
                ]
              },
              {
                "lid": "A=1@O=S+U Berlin Hauptbahnhof@X=13369072@Y=52525607@U=86@L=900003201@",
                "type": "S",
                "name": "S+U Berlin Hauptbahnhof",
                "icoX": 8,
                "extId": "900003201",
                "state": "F",
                "crd": {
                  "x": 13369072,
                  "y": 52525607,
                  "floor": 0
                },
                "pCls": 111,
                "globalIdL": [
                  {
                    "id": "de:11000:900003201",
                    "type": "A"
                  }
                ],
                "chgTime": "001000",
                "countryCodeL": [
                  "de"
                ]
              },
              {
                "lid": "A=1@O=U Osloer Str. (Berlin)@X=13373279@Y=52557105@U=86@L=900009202@",
                "type": "S",
                "name": "U Osloer Str. (Berlin)",
                "icoX": 2,
                "extId": "900009202",
                "state": "F",
                "crd": {
                  "x": 13373279,
                  "y": 52557105,
                  "floor": 0
                },
                "pCls": 14,
                "globalIdL": [
                  {
                    "id": "de:11000:900009202",
                    "type": "A"
                  }
                ],
                "chgTime": "000300",
                "countryCodeL": [
                  "de"
                ]
              }
            ],
            "prodL": [
              {
                "pid": "L::3::Bus::B3041857784::de:VBB:11000000|Bus|147:::*",
                "name": "147",
                "nameS": "147",
                "number": "147",
                "icoX": 0,
                "cls": 8,
                "oprX": 0,
                "prodCtx": {
                  "name": "     147",
                  "num": "100170",
                  "line": "147",
                  "lineId": "de:VBB:11000000|Bus|147:",
                  "matchId": "147",
                  "catOut": "Bus     ",
                  "catOutS": "B",
                  "catOutL": "Bus     ",
                  "catIn": "B",
                  "catCode": "3",
                  "admin": "BVB---"
                },
                "himIdL": [
                  "HIM_FREETEXT_222451",
                  "HIM_FREETEXT_248518"
                ]
              },
              {
                "name": "U6",
                "nameS": "U6",
                "icoX": 2,
                "cls": 2,
                "prodCtx": {
                  "name": "U6",
                  "line": "U6",
                  "lineId": "de:VBB:11000000|U-Bahn|U6:",
                  "catOut": "U       ",
                  "catOutS": "U",
                  "catOutL": "U       "
                }
              },
              {
                "name": "120",
                "nameS": "120",
                "icoX": 0,
                "cls": 8,
                "prodCtx": {
                  "name": "120",
                  "line": "120",
                  "lineId": "de:VBB:11000000|Bus|120:",
                  "catOut": "Bus     ",
                  "catOutS": "B",
                  "catOutL": "Bus     "
                }
              },
              {
                "name": "147",
                "nameS": "147",
                "icoX": 0,
                "cls": 8,
                "prodCtx": {
                  "name": "147",
                  "line": "147",
                  "lineId": "de:VBB:11000000|Bus|147:",
                  "catOut": "Bus     ",
                  "catOutS": "B",
                  "catOutL": "Bus     "
                }
              },
              {
                "name": "M27",
                "nameS": "M27",
                "icoX": 0,
                "cls": 8,
                "prodCtx": {
                  "name": "M27",
                  "line": "M27",
                  "lineId": "de:VBB:11000000|Bus|M27:",
                  "catOut": "Bus     ",
                  "catOutS": "B",
                  "catOutL": "Bus     "
                }
              },
              {
                "name": "N6",
                "nameS": "N6",
                "icoX": 0,
                "cls": 8,
                "prodCtx": {
                  "name": "N6",
                  "line": "N6",
                  "lineId": "de:VBB:11000000|Bus|N6:",
                  "catOut": "Bus     ",
                  "catOutS": "B",
                  "catOutL": "Bus     "
                }
              },
              {
                "name": "N20",
                "nameS": "N20",
                "icoX": 0,
                "cls": 8,
                "prodCtx": {
                  "name": "N20",
                  "line": "N20",
                  "lineId": "de:VBB:11000000|Bus|N20:",
                  "catOut": "Bus     ",
                  "catOutS": "B",
                  "catOutL": "Bus     "
                }
              },
              {
                "pid": "L::1::U::B3041857784::de:VBB:11000000|U-Bahn|U6:::*",
                "name": "U6",
                "nameS": "U6",
                "number": "U6",
                "icoX": 2,
                "cls": 2,
                "oprX": 0,
                "prodCtx": {
                  "name": "      U6",
                  "num": "16482",
                  "line": "U6",
                  "lineId": "de:VBB:11000000|U-Bahn|U6:",
                  "matchId": "U6",
                  "catOut": "U       ",
                  "catOutS": "U",
                  "catOutL": "U       ",
                  "catIn": "U",
                  "catCode": "1",
                  "admin": "BVU---"
                },
                "himIdL": [
                  "HIM_FREETEXT_258304"
                ]
              },
              {
                "pid": "L::3::Bus::B3041857784::de:VBB:11000000|Bus|120:::*",
                "name": "120",
                "nameS": "120",
                "number": "120",
                "icoX": 0,
                "cls": 8,
                "oprX": 0,
                "prodCtx": {
                  "name": "     120",
                  "num": "3470",
                  "line": "120",
                  "lineId": "de:VBB:11000000|Bus|120:",
                  "matchId": "120",
                  "catOut": "Bus     ",
                  "catOutS": "B",
                  "catOutL": "Bus     ",
                  "catIn": "B",
                  "catCode": "3",
                  "admin": "BVB---"
                }
              },
              {
                "pid": "L::3::Bus::B3041857784::de:VBB:11000000|Bus|120:::*",
                "name": "120",
                "nameS": "120",
                "number": "120",
                "icoX": 0,
                "cls": 8,
                "oprX": 0,
                "prodCtx": {
                  "name": "     120",
                  "num": "138166",
                  "line": "120",
                  "lineId": "de:VBB:11000000|Bus|120:",
                  "matchId": "120",
                  "catOut": "Bus     ",
                  "catOutS": "B",
                  "catOutL": "Bus     ",
                  "catIn": "B",
                  "catCode": "3",
                  "admin": "BVB---"
                }
              },
              {
                "pid": "L::3::Bus::B3041857784::de:VBB:11000000|Bus|M27:::*",
                "name": "M27",
                "nameS": "M27",
                "number": "M27",
                "icoX": 0,
                "cls": 8,
                "oprX": 0,
                "prodCtx": {
                  "name": "     M27",
                  "num": "13893",
                  "line": "M27",
                  "lineId": "de:VBB:11000000|Bus|M27:",
                  "matchId": "M27",
                  "catOut": "Bus     ",
                  "catOutS": "B",
                  "catOutL": "Bus     ",
                  "catIn": "B",
                  "catCode": "3",
                  "admin": "BVB---"
                },
                "himIdL": [
                  "HIM_FREETEXT_248771",
                  "HIM_FREETEXT_248773"
                ]
              }
            ],
            "opL": [
              {
                "name": "Berliner Verkehrsbetriebe",
                "url": "https://www.bvg.de/",
                "icoX": 1,
                "id": "796",
                "matchId": "Berliner Verkehrsbetriebe"
              }
            ],
            "remL": [
              {
                "type": "A",
                "code": "text.occup.loc.max.11",
                "icoX": 3,
                "txtN": "low occupancy expected"
              },
              {
                "type": "A",
                "code": "text.occup.loc.max.12",
                "icoX": 6,
                "txtN": "medium occupancy expected"
              },
              {
                "type": "A",
                "code": "FK",
                "prio": 260,
                "icoX": 7,
                "txtN": "Bicycle conveyance"
              }
            ],
            "icoL": [
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
                "res": "BVG",
                "txt": "Berliner Verkehrsbetriebe"
              },
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
                "res": "occup_fig_low"
              },
              {
                "res": "rt_ont",
                "txtA": "on time"
              },
              {
                "res": "rt_cnf"
              },
              {
                "res": "occup_fig_mid"
              },
              {
                "res": "attr_info"
              },
              {
                "res": "prod_comm_t",
                "fg": {
                  "r": 255,
                  "g": 255,
                  "b": 255
                },
                "bg": {
                  "r": 55,
                  "g": 135,
                  "b": 74
                }
              }
            ],
            "tcocL": [
              {
                "c": "SECOND",
                "r": 11,
                "s": "L"
              },
              {
                "c": "SECOND",
                "r": 12,
                "s": "M"
              }
            ],
            "lDrawStyleL": [
              {
                "sIcoX": 0,
                "type": "SOLID",
                "bg": {
                  "r": 153,
                  "g": 51,
                  "b": 153
                }
              },
              {
                "type": "SOLID",
                "bg": {
                  "r": 153,
                  "g": 51,
                  "b": 153
                }
              },
              {
                "sIcoX": 2,
                "type": "SOLID",
                "bg": {
                  "r": 0,
                  "g": 51,
                  "b": 153
                }
              },
              {
                "type": "SOLID",
                "bg": {
                  "r": 0,
                  "g": 51,
                  "b": 153
                }
              }
            ],
            "timeStyleL": [
              {
                "mode": "ABS"
              },
              {
                "mode": "CNT"
              },
              {
                "mode": "HIDE"
              },
              {
                "mode": "ABS",
                "icoX": 4
              },
              {
                "mode": "CNT",
                "icoX": 5
              },
              {
                "mode": "DLT",
                "fg": {
                  "r": 212,
                  "g": 12,
                  "b": 66
                }
              },
              {
                "mode": "ABS",
                "fg": {
                  "r": 212,
                  "g": 12,
                  "b": 66
                }
              }
            ]
          },
          "type": "DEP",
          "jnyL": [
            {
              "jid": "1|61961|4|86|9012025",
              "date": "20250109",
              "prodX": 7,
              "dirTxt": "Kurt-Schumacher-Platz",
              "dirFlg": "2",
              "status": "P",
              "isRchbl": true,
              "stbStop": {
                "locX": 0,
                "idx": 17,
                "dProdX": 7,
                "dPltfS": {
                  "type": "PL",
                  "txt": "2"
                },
                "dTimeS": "214000",
                "dTimeR": "214000",
                "dTimeFS": {
                  "styleX": 2
                },
                "dTimeFR": {
                  "styleX": 3,
                  "txtA": "21:40 on time"
                },
                "dTimeFC": {
                  "styleX": 4
                },
                "dProgType": "PROGNOSED",
                "dTrnCmpSX": {
                  "tcocX": [
                    1
                  ]
                },
                "msgL": [
                  {
                    "type": "REM",
                    "remX": 1,
                    "sty": "I",
                    "tagL": [
                      "RES_LOC_H3"
                    ],
                    "sort": 415760384
                  }
                ],
                "type": "N"
              },
              "pos": {
                "x": 13379598,
                "y": 52533392
              },
              "msgL": [
                {
                  "type": "REM",
                  "remX": 2,
                  "sty": "I",
                  "fLocX": 0,
                  "tLocX": 2,
                  "tagL": [
                    "RES_JNY_DTL"
                  ],
                  "sort": 839385088
                }
              ],
              "subscr": "F",
              "prodL": [
                {
                  "prodX": 7,
                  "fLocX": 0,
                  "tLocX": 2,
                  "fIdx": 17,
                  "tIdx": 23
                }
              ],
              "sumLDrawStyleX": 2,
              "resLDrawStyleX": 3,
              "trainStartDate": "20250109"
            },
            {
              "jid": "1|89906|0|86|9012025",
              "date": "20250109",
              "prodX": 8,
              "dirTxt": "S+U Wittenau -> Bus 220",
              "dirFlg": "2",
              "status": "P",
              "isRchbl": true,
              "stbStop": {
                "locX": 0,
                "idx": 6,
                "dProdX": 8,
                "dTimeS": "214200",
                "dTimeR": "214300",
                "dTimeFS": {
                  "styleX": 0
                },
                "dTimeFR": {
                  "styleX": 5,
                  "txtA": "1 minutes late"
                },
                "dTimeFC": {
                  "styleX": 4
                },
                "dProgType": "PROGNOSED",
                "dTrnCmpSX": {
                  "tcocX": [
                    1
                  ]
                },
                "msgL": [
                  {
                    "type": "REM",
                    "remX": 1,
                    "sty": "I",
                    "tagL": [
                      "RES_LOC_H3"
                    ],
                    "sort": 415760384
                  }
                ],
                "type": "N"
              },
              "pos": {
                "x": 13374951,
                "y": 52529581
              },
              "subscr": "F",
              "prodL": [
                {
                  "prodX": 8,
                  "fLocX": 0,
                  "tLocX": 3,
                  "fIdx": 6,
                  "tIdx": 47
                }
              ],
              "sumLDrawStyleX": 0,
              "resLDrawStyleX": 1,
              "trainStartDate": "20250109"
            },
            {
              "jid": "1|24674|0|86|9012025",
              "date": "20250109",
              "prodX": 9,
              "dirTxt": "S+U Hauptbahnhof",
              "dirFlg": "1",
              "status": "P",
              "isRchbl": true,
              "stbStop": {
                "locX": 0,
                "idx": 33,
                "dProdX": 9,
                "dTimeS": "214300",
                "dTimeR": "214500",
                "dTimeFS": {
                  "styleX": 0
                },
                "dTimeFR": {
                  "styleX": 5,
                  "txtA": "2 minutes late"
                },
                "dTimeFC": {
                  "styleX": 4
                },
                "dProgType": "PROGNOSED",
                "dTrnCmpSX": {
                  "tcocX": [
                    0
                  ]
                },
                "msgL": [
                  {
                    "type": "REM",
                    "remX": 0,
                    "sty": "I",
                    "tagL": [
                      "RES_LOC_H3"
                    ],
                    "sort": 415760384
                  }
                ],
                "type": "N"
              },
              "pos": {
                "x": 13352999,
                "y": 52549806
              },
              "subscr": "F",
              "prodL": [
                {
                  "prodX": 9,
                  "fLocX": 0,
                  "tLocX": 4,
                  "fIdx": 33,
                  "tIdx": 38
                }
              ],
              "sumLDrawStyleX": 0,
              "resLDrawStyleX": 1,
              "trainStartDate": "20250109"
            },
            {
              "jid": "1|91430|7|86|9012025",
              "date": "20250109",
              "prodX": 10,
              "dirTxt": "S Wollankstr. -> 255 Ri. U Osloer Str.",
              "dirFlg": "2",
              "status": "P",
              "isRchbl": true,
              "stbStop": {
                "locX": 0,
                "idx": 18,
                "dProdX": 10,
                "dTimeS": "214300",
                "dTimeR": "214200",
                "dTimeFS": {
                  "styleX": 0,
                  "txtA": "departure 21:43 is early"
                },
                "dTimeFR": {
                  "styleX": 6,
                  "txtA": ", new departure 21:42 confirmed"
                },
                "dTimeFC": {
                  "styleX": 4
                },
                "dProgType": "PROGNOSED",
                "dTrnCmpSX": {
                  "tcocX": [
                    1
                  ]
                },
                "msgL": [
                  {
                    "type": "REM",
                    "remX": 1,
                    "sty": "I",
                    "tagL": [
                      "RES_LOC_H3"
                    ],
                    "sort": 415760384
                  }
                ],
                "type": "N"
              },
              "pos": {
                "x": 13358950,
                "y": 52534291
              },
              "subscr": "F",
              "prodL": [
                {
                  "prodX": 10,
                  "fLocX": 0,
                  "tLocX": 5,
                  "fIdx": 18,
                  "tIdx": 32
                }
              ],
              "sumLDrawStyleX": 0,
              "resLDrawStyleX": 1,
              "trainStartDate": "20250109"
            },
            {
              "jid": "1|29116|0|86|9012025",
              "date": "20250109",
              "prodX": 0,
              "dirTxt": "U Leopoldplatz via S+U Wedding",
              "dirFlg": "1",
              "status": "P",
              "isRchbl": true,
              "stbStop": {
                "locX": 0,
                "idx": 20,
                "dProdX": 0,
                "dTimeS": "214500",
                "dTimeFS": {
                  "styleX": 0
                },
                "dTimeFC": {
                  "styleX": 1
                },
                "dTrnCmpSX": {
                  "tcocX": [
                    0
                  ]
                },
                "msgL": [
                  {
                    "type": "REM",
                    "remX": 0,
                    "sty": "I",
                    "tagL": [
                      "RES_LOC_H3"
                    ],
                    "sort": 415760384
                  }
                ],
                "type": "N"
              },
              "pos": {
                "x": 13372560,
                "y": 52527504
              },
              "subscr": "F",
              "prodL": [
                {
                  "prodX": 0,
                  "fLocX": 0,
                  "tLocX": 1,
                  "fIdx": 20,
                  "tIdx": 23
                }
              ],
              "sumLDrawStyleX": 0,
              "resLDrawStyleX": 1,
              "trainStartDate": "20250109"
            }
          ],
          "fpB": "20250105",
          "fpE": "20251213",
          "planrtTS": "1736455064",
          "sD": "20250109",
          "sT": "213824",
          "locRefL": [
            0
          ]
        }
      }
    ]
  }
    ```
</details>  
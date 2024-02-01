curl -X POST "https://bvg-apps-ext.hafas.de/bin/mgate.exe" \
  -H "Content-Type: application/json" \
  -H "user-agent: YourUserAgentString" \
  --data '{
    "lang": "en",
    "svcReqL": [
      {
        "meth": "LocDetails",
        "req": {
          "locL": [{"type":"S","lid":"A=1@L=900008102@"}]
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

# response
 {
   "ver": "1.44",
   "ext": "BVG.1",
   "lang": "eng",
   "id": "ztw6whusmqxvpw4k",
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
               "gidL": [
                 "A×de:11000:900008104"
               ],
               "TZOffset": 60,
               "chgTime": "000300"
             }
           ],
           "prodL": [
             {
               "pid": "L::1::U::B3041857784::U_3041857784_U6::*",
               "name": "U6",
               "nameS": "U6",
               "icoX": 0,
               "cls": 2,
               "prodCtx": {
                 "name": "U6",
                 "line": "U6",
                 "lineId": "U6",
                 "catOut": "U       ",
                 "catOutS": "U",
                 "catOutL": "U       "
               }
             },
             {
               "pid": "L::3::Bus::B3041857784::Bus_3041857784_120::*",
               "name": "120",
               "nameS": "120",
               "icoX": 1,
               "cls": 8,
               "prodCtx": {
                 "name": "120",
                 "line": "120",
                 "lineId": "120",
                 "catOut": "Bus     ",
                 "catOutS": "B",
                 "catOutL": "Bus     "
               }
             },
             {
               "pid": "L::3::Bus::B3041857784::Bus_3041857784_147::*",
               "name": "147",
               "nameS": "147",
               "icoX": 1,
               "cls": 8,
               "prodCtx": {
                 "name": "147",
                 "line": "147",
                 "lineId": "147",
                 "catOut": "Bus     ",
                 "catOutS": "B",
                 "catOutL": "Bus     "
               }
             },
             {
               "pid": "L::3::Bus::B3041857784::Bus_3041857784_M27::*",
               "name": "M27",
               "nameS": "M27",
               "icoX": 1,
               "cls": 8,
               "prodCtx": {
                 "name": "M27",
                 "line": "M27",
                 "lineId": "M27",
                 "catOut": "Bus     ",
                 "catOutS": "B",
                 "catOutL": "Bus     "
               }
             },
             {
               "pid": "L::3::Bus::B3041857784::Bus_3041857784_N6::*",
               "name": "N6",
               "nameS": "N6",
               "icoX": 1,
               "cls": 8,
               "prodCtx": {
                 "name": "N6",
                 "line": "N6",
                 "lineId": "N6",
                 "catOut": "Bus     ",
                 "catOutS": "B",
                 "catOutL": "Bus     "
               }
             },
             {
               "pid": "L::3::Bus::B3041857784::Bus_3041857784_N20::*",
               "name": "N20",
               "nameS": "N20",
               "icoX": 1,
               "cls": 8,
               "prodCtx": {
                 "name": "N20",
                 "line": "N20",
                 "lineId": "N20",
                 "catOut": "Bus     ",
                 "catOutS": "B",
                 "catOutL": "Bus     "
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
             "wt": 12407,
             "stopLocL": [
               0
             ],
             "gidL": [
               "A×de:11000:900008102"
             ],
             "TZOffset": 60,
             "chgTime": "000300"
           }
         ]
       }
     }
   ]
 }

# ERROR unvalide station id
curl -X POST "https://bvg-apps-ext.hafas.de/bin/mgate.exe" \
  -H "Content-Type: application/json" \
  -H "user-agent: YourUserAgentString" \
  --data '{
    "lang": "en",
    "svcReqL": [
      {
        "meth": "LocDetails",
        "req": {
          "locL": [{"type":"S","lid":"A=1@L=000009999@"}]
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

# response
{
  "ver": "1.44",
  "ext": "BVG.1",
  "lang": "eng",
  "id": "3s2uwj8cki269gwg",
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
      "err": "LOCATION",
      "errTxt": "HCI Service: location missing or invalid",
      "errTxtOut": "An internal error occurred during the search"
    }
  ]
}

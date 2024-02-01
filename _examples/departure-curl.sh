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
      "meth": "StationBoard",
      "req": {
        "type": "DEP",
        "date": "20240124",
        "time": "210100",
        "stbLoc": {
          "type": "S",
          "lid": "A=1@L=900008102@"
        },
        "jnyFltrL": [
          {
            "type": "PROD",
            "mode": "INC",
            "value": 127
          }
        ],
        "dur": 10
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

# Result
{
  "ver": "1.44",
  "ext": "BVG.1",
  "lang": "eng",
  "id": "w9i4xz8446x3gg4s",
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
              "gidL": [
                "A×de:11000:900008102"
              ],
              "chgTime": "000300"
            },
            {
              "lid": "A=1@O=Wilhelmsruher Damm (Berlin)@X=13366726@Y=52597278@U=86@L=900096307@",
              "type": "S",
              "name": "Wilhelmsruher Damm (Berlin)",
              "icoX": 0,
              "extId": "900096307",
              "state": "F",
              "crd": {
                "x": 13366726,
                "y": 52597278,
                "floor": 0
              },
              "pCls": 8,
              "gidL": [
                "A×de:11000:900096307"
              ],
              "chgTime": "000300"
            },
            {
              "lid": "A=1@O=S+U Jungfernheide Bhf (Berlin)@X=13300125@Y=52530452@U=86@L=900020201@",
              "type": "S",
              "name": "S+U Jungfernheide Bhf (Berlin)",
              "icoX": 4,
              "extId": "900020201",
              "state": "F",
              "crd": {
                "x": 13300125,
                "y": 52530452,
                "floor": 0
              },
              "pCls": 75,
              "gidL": [
                "A×de:11000:900020201"
              ],
              "chgTime": "000300"
            },
            {
              "lid": "A=1@O=Hadlichstr. (Berlin)@X=13414369@Y=52568854@U=86@L=900130005@",
              "type": "S",
              "name": "Hadlichstr. (Berlin)",
              "icoX": 0,
              "extId": "900130005",
              "state": "F",
              "crd": {
                "x": 13414369,
                "y": 52568854,
                "floor": 0
              },
              "pCls": 8,
              "gidL": [
                "A×de:11000:900130005"
              ],
              "chgTime": "000300"
            },
            {
              "lid": "A=1@O=S+U Tempelhof (Berlin)@X=13385900@Y=52470575@U=86@L=900068201@",
              "type": "S",
              "name": "S+U Tempelhof (Berlin)",
              "icoX": 4,
              "extId": "900068201",
              "state": "F",
              "crd": {
                "x": 13385900,
                "y": 52470575,
                "floor": 0
              },
              "pCls": 11,
              "gidL": [
                "A×de:11000:900068201"
              ],
              "chgTime": "000300"
            },
            {
              "lid": "A=1@O=S+U Berlin Hauptbahnhof@X=13369072@Y=52525607@U=86@L=900003201@",
              "type": "S",
              "name": "S+U Berlin Hauptbahnhof",
              "icoX": 4,
              "extId": "900003201",
              "state": "F",
              "crd": {
                "x": 13369072,
                "y": 52525607,
                "floor": 0
              },
              "pCls": 111,
              "gidL": [
                "A×de:11000:900003201"
              ],
              "chgTime": "001000"
            },
            {
              "lid": "A=1@O=U Alt-Mariendorf (Berlin)@X=13387922@Y=52438610@U=86@L=900070301@",
              "type": "S",
              "name": "U Alt-Mariendorf (Berlin)",
              "icoX": 2,
              "extId": "900070301",
              "state": "F",
              "crd": {
                "x": 13387922,
                "y": 52438610,
                "floor": 0
              },
              "pCls": 2,
              "gidL": [
                "A×de:11000:900070301"
              ],
              "chgTime": "000300"
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
              "gidL": [
                "A×de:11000:900086102"
              ],
              "chgTime": "000300"
            }
          ],
          "prodL": [
            {
              "pid": "L::3::Bus::B3041857784::Bus_3041857784_120::*",
              "name": "120",
              "nameS": "120",
              "number": "120",
              "icoX": 0,
              "cls": 8,
              "oprX": 0,
              "prodCtx": {
                "name": "     120",
                "num": "33153",
                "line": "120",
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
              "name": "U6",
              "nameS": "U6",
              "icoX": 2,
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
              "name": "120",
              "nameS": "120",
              "icoX": 0,
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
              "name": "147",
              "nameS": "147",
              "icoX": 0,
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
              "name": "M27",
              "nameS": "M27",
              "icoX": 0,
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
              "name": "N6",
              "nameS": "N6",
              "icoX": 0,
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
              "name": "N20",
              "nameS": "N20",
              "icoX": 0,
              "cls": 8,
              "prodCtx": {
                "name": "N20",
                "line": "N20",
                "lineId": "N20",
                "catOut": "Bus     ",
                "catOutS": "B",
                "catOutL": "Bus     "
              }
            },
            {
              "pid": "L::3::Bus::B3041857784::Bus_3041857784_M27::*",
              "name": "M27",
              "nameS": "M27",
              "number": "M27",
              "icoX": 0,
              "cls": 8,
              "oprX": 0,
              "prodCtx": {
                "name": "     M27",
                "num": "28532",
                "line": "M27",
                "matchId": "M27",
                "catOut": "Bus     ",
                "catOutS": "B",
                "catOutL": "Bus     ",
                "catIn": "B",
                "catCode": "3",
                "admin": "BVB---"
              }
            },
            {
              "pid": "L::3::Bus::B3041857784::Bus_3041857784_M27::*",
              "name": "M27",
              "nameS": "M27",
              "number": "M27",
              "icoX": 0,
              "cls": 8,
              "oprX": 0,
              "prodCtx": {
                "name": "     M27",
                "num": "27352",
                "line": "M27",
                "matchId": "M27",
                "catOut": "Bus     ",
                "catOutS": "B",
                "catOutL": "Bus     ",
                "catIn": "B",
                "catCode": "3",
                "admin": "BVB---"
              }
            },
            {
              "pid": "L::1::U::B3041857784::U_3041857784_U6::*",
              "name": "U6",
              "nameS": "U6",
              "number": "U6",
              "icoX": 2,
              "cls": 2,
              "oprX": 0,
              "prodCtx": {
                "name": "      U6",
                "num": "16687",
                "line": "U6",
                "matchId": "U6",
                "catOut": "U       ",
                "catOutS": "U",
                "catOutL": "U       ",
                "catIn": "U",
                "catCode": "1",
                "admin": "BVU---"
              }
            },
            {
              "pid": "L::3::Bus::B3041857784::Bus_3041857784_120::*",
              "name": "120",
              "nameS": "120",
              "number": "120",
              "icoX": 0,
              "cls": 8,
              "oprX": 0,
              "prodCtx": {
                "name": "     120",
                "num": "33061",
                "line": "120",
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
              "pid": "L::1::U::B3041857784::U_3041857784_U6::*",
              "name": "U6",
              "nameS": "U6",
              "number": "U6",
              "icoX": 2,
              "cls": 2,
              "oprX": 0,
              "prodCtx": {
                "name": "      U6",
                "num": "16750",
                "line": "U6",
                "matchId": "U6",
                "catOut": "U       ",
                "catOutS": "U",
                "catOutL": "U       ",
                "catIn": "U",
                "catCode": "1",
                "admin": "BVU---"
              }
            },
            {
              "pid": "L::1::U::B3041857784::U_3041857784_U6::*",
              "name": "U6",
              "nameS": "U6",
              "number": "U6",
              "icoX": 2,
              "cls": 2,
              "oprX": 0,
              "prodCtx": {
                "name": "      U6",
                "num": "16369",
                "line": "U6",
                "matchId": "U6",
                "catOut": "U       ",
                "catOutS": "U",
                "catOutL": "U       ",
                "catIn": "U",
                "catCode": "1",
                "admin": "BVU---"
              }
            }
          ],
          "opL": [
            {
              "name": "Berliner Verkehrsbetriebe",
              "url": "https://www.bvg.de/",
              "icoX": 1,
              "id": "796"
            }
          ],
          "remL": [
            {
              "type": "A",
              "code": "FK",
              "prio": 260,
              "icoX": 5,
              "txtN": "Bicycle conveyance"
            }
          ],
          "himL": [
            {
              "hid": "213723",
              "act": true,
              "head": "Rail strike until Monday, 29 January, 6 p.m.",
              "text": "Major restrictions on Deutsche Bahn railway services (S-Bahn, RE, RB, long-distance services). Underground trains, trams, buses and ferries are running.<br>",
              "icoX": 3,
              "prio": 100,
              "prod": 16383,
              "src": 50,
              "sDate": "20240122",
              "sTime": "020000",
              "eDate": "20240129",
              "eTime": "210000",
              "sDaily": "000000",
              "eDaily": "235900",
              "comp": "VBB",
              "catRefL": [
                0
              ],
              "pubChL": [
                {
                  "name": "TIMETABLE",
                  "fDate": "20240124",
                  "fTime": "075000",
                  "tDate": "20240129",
                  "tTime": "210000"
                }
              ]
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
              "res": "HIM2"
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
            },
            {
              "res": "attr_info"
            }
          ],
          "himMsgCatL": [
            {
              "id": 2
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
            }
          ]
        },
        "type": "DEP",
        "jnyL": [
          {
            "jid": "1|18087|3|86|24012024",
            "date": "20240124",
            "prodX": 0,
            "dirTxt": "Märkisches Viertel, Wilhelmsruher Damm",
            "dirFlg": "2",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 6,
              "dProdX": 0,
              "dTimeS": "210100",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFC": {
                "styleX": 1
              },
              "type": "N"
            },
            "msgL": [
              {
                "type": "HIM",
                "himX": 0,
                "sty": "M",
                "fLocX": 0,
                "tagL": [
                  "RES_GLB_HDR_H3",
                  "SUM_GLB_HDR_H3"
                ],
                "sort": 374341631
              }
            ],
            "subscr": "F",
            "prodL": [
              {
                "prodX": 0,
                "fLocX": 0,
                "tLocX": 1,
                "fIdx": 6,
                "tIdx": 39
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240124"
          },
          {
            "jid": "1|31791|5|86|24012024",
            "date": "20240124",
            "prodX": 7,
            "dirTxt": "S+U Jungfernheide",
            "dirFlg": "1",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 13,
              "dProdX": 7,
              "dTimeS": "210300",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFC": {
                "styleX": 1
              },
              "type": "N"
            },
            "msgL": [
              {
                "type": "HIM",
                "himX": 0,
                "sty": "M",
                "fLocX": 0,
                "tagL": [
                  "RES_GLB_HDR_H3",
                  "SUM_GLB_HDR_H3"
                ],
                "sort": 374341631
              }
            ],
            "subscr": "F",
            "prodL": [
              {
                "prodX": 7,
                "fLocX": 0,
                "tLocX": 2,
                "fIdx": 13,
                "tIdx": 31
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240124"
          },
          {
            "jid": "1|31997|4|86|24012024",
            "date": "20240124",
            "prodX": 8,
            "dirTxt": "S+U Pankow",
            "dirFlg": "2",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 18,
              "dProdX": 8,
              "dTimeS": "210300",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFC": {
                "styleX": 1
              },
              "type": "N"
            },
            "msgL": [
              {
                "type": "HIM",
                "himX": 0,
                "sty": "M",
                "fLocX": 0,
                "tagL": [
                  "RES_GLB_HDR_H3",
                  "SUM_GLB_HDR_H3"
                ],
                "sort": 374341631
              }
            ],
            "subscr": "F",
            "prodL": [
              {
                "prodX": 8,
                "fLocX": 0,
                "tLocX": 3,
                "fIdx": 18,
                "tIdx": 31
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240124"
          },
          {
            "jid": "1|46152|1|86|24012024",
            "date": "20240124",
            "prodX": 9,
            "dirTxt": "Tempelhof",
            "dirFlg": "1",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 6,
              "dProdX": 9,
              "dPltfS": {
                "type": "PL",
                "txt": "1"
              },
              "dTimeS": "210300",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFC": {
                "styleX": 1
              },
              "type": "N"
            },
            "msgL": [
              {
                "type": "REM",
                "remX": 0,
                "sty": "I",
                "fLocX": 0,
                "tLocX": 4,
                "tagL": [
                  "RES_JNY_DTL"
                ],
                "sort": 839385088
              },
              {
                "type": "HIM",
                "himX": 0,
                "sty": "M",
                "fLocX": 0,
                "tagL": [
                  "RES_GLB_HDR_H3",
                  "SUM_GLB_HDR_H3"
                ],
                "sort": 374341631
              }
            ],
            "subscr": "F",
            "prodL": [
              {
                "prodX": 9,
                "fLocX": 0,
                "tLocX": 4,
                "fIdx": 6,
                "tIdx": 18
              }
            ],
            "sumLDrawStyleX": 2,
            "resLDrawStyleX": 3,
            "trainStartDate": "20240124"
          },
          {
            "jid": "1|18000|1|86|24012024",
            "date": "20240124",
            "prodX": 10,
            "dirTxt": "S+U Hauptbahnhof",
            "dirFlg": "1",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 33,
              "dProdX": 10,
              "dTimeS": "210400",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFC": {
                "styleX": 1
              },
              "type": "N"
            },
            "msgL": [
              {
                "type": "HIM",
                "himX": 0,
                "sty": "M",
                "fLocX": 0,
                "tagL": [
                  "RES_GLB_HDR_H3",
                  "SUM_GLB_HDR_H3"
                ],
                "sort": 374341631
              }
            ],
            "subscr": "F",
            "prodL": [
              {
                "prodX": 10,
                "fLocX": 0,
                "tLocX": 5,
                "fIdx": 33,
                "tIdx": 38
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240124"
          },
          {
            "jid": "1|46140|2|86|24012024",
            "date": "20240124",
            "prodX": 11,
            "dirTxt": "Alt-Mariendorf",
            "dirFlg": "1",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 6,
              "dProdX": 11,
              "dPltfS": {
                "type": "PL",
                "txt": "1"
              },
              "dTimeS": "210800",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFC": {
                "styleX": 1
              },
              "type": "N"
            },
            "msgL": [
              {
                "type": "REM",
                "remX": 0,
                "sty": "I",
                "fLocX": 0,
                "tLocX": 6,
                "tagL": [
                  "RES_JNY_DTL"
                ],
                "sort": 839385088
              },
              {
                "type": "HIM",
                "himX": 0,
                "sty": "M",
                "fLocX": 0,
                "tagL": [
                  "RES_GLB_HDR_H3",
                  "SUM_GLB_HDR_H3"
                ],
                "sort": 374341631
              }
            ],
            "subscr": "F",
            "prodL": [
              {
                "prodX": 11,
                "fLocX": 0,
                "tLocX": 6,
                "fIdx": 6,
                "tIdx": 23
              }
            ],
            "sumLDrawStyleX": 2,
            "resLDrawStyleX": 3,
            "trainStartDate": "20240124"
          },
          {
            "jid": "1|46200|1|86|24012024",
            "date": "20240124",
            "prodX": 12,
            "dirTxt": "Kurt-Schumacher-Platz",
            "dirFlg": "2",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 17,
              "dProdX": 12,
              "dPltfS": {
                "type": "PL",
                "txt": "2"
              },
              "dTimeS": "211000",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFC": {
                "styleX": 1
              },
              "type": "N"
            },
            "msgL": [
              {
                "type": "REM",
                "remX": 0,
                "sty": "I",
                "fLocX": 0,
                "tLocX": 7,
                "tagL": [
                  "RES_JNY_DTL"
                ],
                "sort": 839385088
              },
              {
                "type": "HIM",
                "himX": 0,
                "sty": "M",
                "fLocX": 0,
                "tagL": [
                  "RES_GLB_HDR_H3",
                  "SUM_GLB_HDR_H3"
                ],
                "sort": 374341631
              }
            ],
            "subscr": "F",
            "prodL": [
              {
                "prodX": 12,
                "fLocX": 0,
                "tLocX": 7,
                "fIdx": 17,
                "tIdx": 23
              }
            ],
            "sumLDrawStyleX": 2,
            "resLDrawStyleX": 3,
            "trainStartDate": "20240124"
          }
        ],
        "fpB": "20240121",
        "fpE": "20241214",
        "planrtTS": "1706363408",
        "sD": "20240127",
        "sT": "145026",
        "locRefL": [
          0
        ]
      }
    }
  ]
}
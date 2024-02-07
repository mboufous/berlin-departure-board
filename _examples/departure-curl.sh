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
  "id": "vdwaynmq4su2hick",
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
              "lid": "A=1@O=S+U Jungfernheide Bhf (Berlin)@X=13300125@Y=52530452@U=86@L=900020201@",
              "type": "S",
              "name": "S+U Jungfernheide Bhf (Berlin)",
              "icoX": 5,
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
              "icoX": 5,
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
              "gidL": [
                "A×de:11000:900009102"
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
              "lid": "A=1@O=U Märkisches Museum/Inselstr. (Berlin)@X=13411438@Y=52512753@U=86@L=900100516@",
              "type": "S",
              "name": "U Märkisches Museum/Inselstr. (Berlin)",
              "icoX": 0,
              "extId": "900100516",
              "state": "F",
              "crd": {
                "x": 13411438,
                "y": 52512753,
                "floor": 0
              },
              "pCls": 8,
              "gidL": [
                "A×de:11000:900100516"
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
            }
          ],
          "prodL": [
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
                "num": "27960",
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
                "num": "27565",
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
                "num": "16762",
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
              "pid": "L::3::Bus::B3041857784::Bus_3041857784_147::*",
              "name": "147",
              "nameS": "147",
              "number": "147",
              "icoX": 0,
              "cls": 8,
              "oprX": 0,
              "prodCtx": {
                "name": "     147",
                "num": "35582",
                "line": "147",
                "matchId": "147",
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
                "num": "16378",
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
                "num": "16825",
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
              "pid": "L::3::Bus::B3041857784::Bus_3041857784_147::*",
              "name": "147",
              "nameS": "147",
              "number": "147",
              "icoX": 0,
              "cls": 8,
              "oprX": 0,
              "prodCtx": {
                "name": "     147",
                "num": "35633",
                "line": "147",
                "matchId": "147",
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
                "num": "16444",
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
                "num": "31656",
                "line": "120",
                "matchId": "120",
                "catOut": "Bus     ",
                "catOutS": "B",
                "catOutL": "Bus     ",
                "catIn": "B",
                "catCode": "3",
                "admin": "BVB---"
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
              "code": "text.occup.loc.max.11",
              "icoX": 4,
              "txtN": "low occupancy expected"
            },
            {
              "type": "A",
              "code": "text.occup.loc.max.13",
              "icoX": 7,
              "txtN": "high occupancy expected"
            },
            {
              "type": "A",
              "code": "FK",
              "prio": 260,
              "icoX": 8,
              "txtN": "Bicycle conveyance"
            },
            {
              "type": "A",
              "code": "text.occup.loc.max.12",
              "icoX": 9,
              "txtN": "medium occupancy expected"
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
              "res": "rt_cnf"
            },
            {
              "res": "occup_fig_low"
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
              "res": "rt_ont",
              "txtA": "on time"
            },
            {
              "res": "occup_fig_high"
            },
            {
              "res": "attr_info"
            },
            {
              "res": "occup_fig_mid"
            }
          ],
          "tcocL": [
            {
              "c": "SECOND",
              "r": 11
            },
            {
              "c": "SECOND",
              "r": 13
            },
            {
              "c": "SECOND",
              "r": 12
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
              "mode": "DLT",
              "fg": {
                "r": 212,
                "g": 12,
                "b": 66
              }
            },
            {
              "mode": "CNT",
              "icoX": 3
            },
            {
              "mode": "HIDE"
            },
            {
              "mode": "ABS",
              "icoX": 6
            }
          ]
        },
        "type": "DEP",
        "jnyL": [
          {
            "jid": "1|31746|4|86|7022024",
            "date": "20240207",
            "prodX": 0,
            "dirTxt": "S+U Jungfernheide",
            "dirFlg": "1",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 13,
              "dProdX": 0,
              "dTimeS": "205300",
              "dTimeR": "205500",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFR": {
                "styleX": 1,
                "txtA": "2 minutes late"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
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
              "x": 13374133,
              "y": 52547874
            },
            "subscr": "F",
            "prodL": [
              {
                "prodX": 0,
                "fLocX": 0,
                "tLocX": 1,
                "fIdx": 13,
                "tIdx": 31
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240207"
          },
          {
            "jid": "1|31823|3|86|7022024",
            "date": "20240207",
            "prodX": 7,
            "dirTxt": "S+U Pankow",
            "dirFlg": "2",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 15,
              "dProdX": 7,
              "dTimeS": "205300",
              "dTimeR": "205500",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFR": {
                "styleX": 1,
                "txtA": "2 minutes late"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
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
              "x": 13362204,
              "y": 52538111
            },
            "subscr": "F",
            "prodL": [
              {
                "prodX": 7,
                "fLocX": 0,
                "tLocX": 2,
                "fIdx": 15,
                "tIdx": 28
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240207"
          },
          {
            "jid": "1|45903|0|86|7022024",
            "date": "20240207",
            "prodX": 8,
            "dirTxt": "Tempelhof",
            "dirFlg": "1",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 6,
              "dProdX": 8,
              "dPltfS": {
                "type": "PL",
                "txt": "1"
              },
              "dTimeS": "205300",
              "dTimeR": "205300",
              "dTimeFS": {
                "styleX": 3
              },
              "dTimeFR": {
                "styleX": 4,
                "txtA": "20:53 on time"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
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
              "x": 13364973,
              "y": 52543046
            },
            "msgL": [
              {
                "type": "REM",
                "remX": 2,
                "sty": "I",
                "fLocX": 0,
                "tLocX": 3,
                "tagL": [
                  "RES_JNY_DTL"
                ],
                "sort": 839385088
              }
            ],
            "subscr": "F",
            "prodL": [
              {
                "prodX": 8,
                "fLocX": 0,
                "tLocX": 3,
                "fIdx": 6,
                "tIdx": 18
              }
            ],
            "sumLDrawStyleX": 2,
            "resLDrawStyleX": 3,
            "trainStartDate": "20240207"
          },
          {
            "jid": "1|21226|0|86|7022024",
            "date": "20240207",
            "prodX": 9,
            "dirTxt": "U Leopoldplatz via S+U Wedding",
            "dirFlg": "1",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 18,
              "dProdX": 9,
              "dTimeS": "205400",
              "dTimeR": "205400",
              "dTimeFS": {
                "styleX": 3
              },
              "dTimeFR": {
                "styleX": 4,
                "txtA": "20:54 on time"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
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
              "x": 13363472,
              "y": 52538705
            },
            "subscr": "F",
            "prodL": [
              {
                "prodX": 9,
                "fLocX": 0,
                "tLocX": 4,
                "fIdx": 18,
                "tIdx": 21
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240207"
          },
          {
            "jid": "1|46002|1|86|7022024",
            "date": "20240207",
            "prodX": 10,
            "dirTxt": "Kurt-Schumacher-Platz",
            "dirFlg": "2",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 17,
              "dProdX": 10,
              "dPltfS": {
                "type": "PL",
                "txt": "2"
              },
              "dTimeS": "205500",
              "dTimeR": "205600",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFR": {
                "styleX": 1,
                "txtA": "1 minutes late"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
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
              "x": 13387419,
              "y": 52525544
            },
            "msgL": [
              {
                "type": "REM",
                "remX": 2,
                "sty": "I",
                "fLocX": 0,
                "tLocX": 5,
                "tagL": [
                  "RES_JNY_DTL"
                ],
                "sort": 839385088
              }
            ],
            "subscr": "F",
            "prodL": [
              {
                "prodX": 10,
                "fLocX": 0,
                "tLocX": 5,
                "fIdx": 17,
                "tIdx": 23
              }
            ],
            "sumLDrawStyleX": 2,
            "resLDrawStyleX": 3,
            "trainStartDate": "20240207"
          },
          {
            "jid": "1|45891|1|86|7022024",
            "date": "20240207",
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
              "dTimeS": "205800",
              "dTimeR": "205800",
              "dTimeFS": {
                "styleX": 3
              },
              "dTimeFR": {
                "styleX": 4,
                "txtA": "20:58 on time"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
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
              "x": 13334661,
              "y": 52560171
            },
            "msgL": [
              {
                "type": "REM",
                "remX": 2,
                "sty": "I",
                "fLocX": 0,
                "tLocX": 6,
                "tagL": [
                  "RES_JNY_DTL"
                ],
                "sort": 839385088
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
            "trainStartDate": "20240207"
          },
          {
            "jid": "1|21270|0|86|7022024",
            "date": "20240207",
            "prodX": 12,
            "dirTxt": "U Märkisches Museum via S+U Hauptbahnhof",
            "dirFlg": "2",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 3,
              "dProdX": 12,
              "dTimeS": "205900",
              "dTimeR": "205900",
              "dTimeFS": {
                "styleX": 3
              },
              "dTimeFR": {
                "styleX": 4,
                "txtA": "20:59 on time"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
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
            "subscr": "F",
            "prodL": [
              {
                "prodX": 12,
                "fLocX": 0,
                "tLocX": 7,
                "fIdx": 3,
                "tIdx": 21
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240207"
          },
          {
            "jid": "1|45951|0|86|7022024",
            "date": "20240207",
            "prodX": 13,
            "dirTxt": "Kurt-Schumacher-Platz",
            "dirFlg": "2",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 17,
              "dProdX": 13,
              "dPltfS": {
                "type": "PL",
                "txt": "2"
              },
              "dTimeS": "210000",
              "dTimeR": "210500",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFR": {
                "styleX": 1,
                "txtA": "5 minutes late"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
              "dTrnCmpSX": {
                "tcocX": [
                  2
                ]
              },
              "msgL": [
                {
                  "type": "REM",
                  "remX": 3,
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
              "x": 13386709,
              "y": 52490891
            },
            "msgL": [
              {
                "type": "REM",
                "remX": 2,
                "sty": "I",
                "fLocX": 0,
                "tLocX": 5,
                "tagL": [
                  "RES_JNY_DTL"
                ],
                "sort": 839385088
              }
            ],
            "subscr": "F",
            "prodL": [
              {
                "prodX": 13,
                "fLocX": 0,
                "tLocX": 5,
                "fIdx": 17,
                "tIdx": 23
              }
            ],
            "sumLDrawStyleX": 2,
            "resLDrawStyleX": 3,
            "trainStartDate": "20240207"
          },
          {
            "jid": "1|18141|3|86|7022024",
            "date": "20240207",
            "prodX": 14,
            "dirTxt": "Märkisches Viertel, Wilhelmsruher Damm",
            "dirFlg": "2",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 6,
              "dProdX": 14,
              "dTimeS": "210100",
              "dTimeR": "210100",
              "dTimeFS": {
                "styleX": 3
              },
              "dTimeFR": {
                "styleX": 4,
                "txtA": "21:01 on time"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
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
            "subscr": "F",
            "prodL": [
              {
                "prodX": 14,
                "fLocX": 0,
                "tLocX": 8,
                "fIdx": 6,
                "tIdx": 39
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240207"
          }
        ],
        "fpB": "20240128",
        "fpE": "20241214",
        "planrtTS": "1707335504",
        "sD": "20240207",
        "sT": "205201",
        "locRefL": [
          0
        ]
      }
    }
  ]
}


# Another response
{
  "ver": "1.44",
  "ext": "BVG.1",
  "lang": "eng",
  "id": "8f28mju6m6x3824k",
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
              "gidL": [
                "A×de:11000:900092202"
              ],
              "chgTime": "000300"
            },
            {
              "lid": "A=1@O=U Alt-Mariendorf (Berlin) [Bus Frieden-/Reißeckstr@X=13387608@Y=52438969@U=86@L=900070702@",
              "type": "S",
              "name": "U Alt-Mariendorf (Berlin) [Bus Frieden-/Reißeckstr",
              "icoX": 0,
              "extId": "900070702",
              "state": "F",
              "crd": {
                "x": 13387608,
                "y": 52438969,
                "floor": 0
              },
              "pCls": 8,
              "gidL": [
                "A×de:11000:900070702"
              ],
              "chgTime": "000300"
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
              "gidL": [
                "A×de:11000:900003201"
              ],
              "chgTime": "001000"
            }
          ],
          "prodL": [
            {
              "pid": "L::3::Bus::B3041857784::Bus_3041857784_N20::*",
              "name": "N20",
              "nameS": "N20",
              "number": "N20",
              "icoX": 0,
              "cls": 8,
              "oprX": 0,
              "prodCtx": {
                "name": "     N20",
                "num": "56358",
                "line": "N20",
                "matchId": "N20",
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
              "pid": "L::3::Bus::B3041857784::Bus_3041857784_N6::*",
              "name": "N6",
              "nameS": "N6",
              "number": "N6",
              "icoX": 0,
              "cls": 8,
              "oprX": 0,
              "prodCtx": {
                "name": "      N6",
                "num": "28175",
                "line": "N6",
                "matchId": "N6",
                "catOut": "Bus     ",
                "catOutS": "B",
                "catOutL": "Bus     ",
                "catIn": "B",
                "catCode": "3",
                "admin": "BVB---"
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
                "num": "1533",
                "line": "120",
                "matchId": "120",
                "catOut": "Bus     ",
                "catOutS": "B",
                "catOutL": "Bus     ",
                "catIn": "B",
                "catCode": "3",
                "admin": "BVB---"
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
              "code": "jh",
              "prio": 350,
              "icoX": 5,
              "txtN": "ab U Alt-Mariendorf weiter als Bus M77 Richtung Waldsassener Str."
            },
            {
              "type": "A",
              "code": "text.occup.loc.max.11",
              "icoX": 7,
              "txtN": "low occupancy expected"
            }
          ],
          "himL": [
            {
              "hid": "214098",
              "act": true,
              "head": "On Friday, 2 February: Strike at 15 transport operators",
              "text": "In Berlin at BVG from approx. 3 a.m. to 10 a.m., in Brandenburg all day. This affects underground trains, trams and buses. Long-distance and regional trains as well as S-Bahn trains will run as scheduled. Info: <a href=\"http://www.vbb.de/streik\">www.vbb.de/streik</a> and <a href=\"http://www.bvg.de\">www.bvg.de</a>",
              "icoX": 4,
              "prio": 100,
              "prod": 16383,
              "src": 50,
              "sDate": "20240129",
              "sTime": "165300",
              "eDate": "20240202",
              "eTime": "235900",
              "sDaily": "000000",
              "eDaily": "235900",
              "comp": "VBB",
              "catRefL": [
                0
              ],
              "pubChL": [
                {
                  "name": "TIMETABLE",
                  "fDate": "20240130",
                  "fTime": "082600",
                  "tDate": "20240202",
                  "tTime": "235900"
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
              "res": "rt_cnf"
            },
            {
              "res": "HIM2"
            },
            {
              "res": "attr_info"
            },
            {
              "res": "rt_ont",
              "txtA": "on time"
            },
            {
              "res": "occup_fig_low"
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
              "r": 11
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
            }
          ],
          "timeStyleL": [
            {
              "mode": "ABS"
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
              "mode": "CNT",
              "icoX": 3
            },
            {
              "mode": "HIDE"
            },
            {
              "mode": "ABS",
              "icoX": 6
            }
          ]
        },
        "type": "DEP",
        "jnyL": [
          {
            "jid": "1|36334|0|86|1022024",
            "date": "20240201",
            "prodX": 0,
            "dirTxt": "Frohnau, Hainbuchenstr.",
            "dirFlg": "2",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 6,
              "dProdX": 0,
              "dTimeS": "01005300",
              "dTimeR": "01005600",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFR": {
                "styleX": 1,
                "txtA": "3 minutes late"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
              "type": "N"
            },
            "pos": {
              "x": 13368200,
              "y": 52541383
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
                "tIdx": 48
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240201"
          },
          {
            "jid": "1|36193|0|86|1022024",
            "date": "20240201",
            "prodX": 7,
            "dirTxt": "U Alt-Mariendorf -> M77 Richtung Marienfelde",
            "dirFlg": "1",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 9,
              "dProdX": 7,
              "dTimeS": "01010000",
              "dTimeR": "01010100",
              "dTimeFS": {
                "styleX": 0
              },
              "dTimeFR": {
                "styleX": 1,
                "txtA": "1 minutes late"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
              "type": "N"
            },
            "pos": {
              "x": 13352190,
              "y": 52550292
            },
            "msgL": [
              {
                "type": "REM",
                "remX": 0,
                "sty": "I",
                "fLocX": 0,
                "tLocX": 2,
                "tagL": [
                  "RES_JNY_DTL"
                ],
                "sort": 851181568
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
                "prodX": 7,
                "fLocX": 0,
                "tLocX": 2,
                "fIdx": 9,
                "tIdx": 35
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240201"
          },
          {
            "jid": "1|62262|2|86|1022024",
            "date": "20240201",
            "prodX": 8,
            "dirTxt": "S+U Hauptbahnhof",
            "dirFlg": "1",
            "status": "P",
            "isRchbl": true,
            "stbStop": {
              "locX": 0,
              "idx": 40,
              "dProdX": 8,
              "dTimeS": "01010200",
              "dTimeR": "01010200",
              "dTimeFS": {
                "styleX": 3
              },
              "dTimeFR": {
                "styleX": 4,
                "txtA": "01:02 on time"
              },
              "dTimeFC": {
                "styleX": 2
              },
              "dProgType": "CALCULATED",
              "dTrnCmpSX": {
                "tcocX": [
                  0
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
              "x": 13350734,
              "y": 52551065
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
                "fIdx": 40,
                "tIdx": 45
              }
            ],
            "sumLDrawStyleX": 0,
            "resLDrawStyleX": 1,
            "trainStartDate": "20240201"
          }
        ],
        "fpB": "20240128",
        "fpE": "20241214",
        "planrtTS": "1706831708",
        "sD": "20240202",
        "sT": "005531",
        "locRefL": [
          0
        ]
      }
    }
  ]
}
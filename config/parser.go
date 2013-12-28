package config

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode/utf8"
)

// generated from /usr/share/X11/rgb.txt
// TODO actually query the X server for color names
var colors = map[string]int{
	"snow":                   0xfffafa,
	"ghost white":            0xf8f8ff,
	"ghostwhite":             0xf8f8ff,
	"white smoke":            0xf5f5f5,
	"whitesmoke":             0xf5f5f5,
	"gainsboro":              0xdcdcdc,
	"floral white":           0xfffaf0,
	"floralwhite":            0xfffaf0,
	"old lace":               0xfdf5e6,
	"oldlace":                0xfdf5e6,
	"linen":                  0xfaf0e6,
	"antique white":          0xfaebd7,
	"antiquewhite":           0xfaebd7,
	"papaya whip":            0xffefd5,
	"papayawhip":             0xffefd5,
	"blanched almond":        0xffebcd,
	"blanchedalmond":         0xffebcd,
	"bisque":                 0xffe4c4,
	"peach puff":             0xffdab9,
	"peachpuff":              0xffdab9,
	"navajo white":           0xffdead,
	"navajowhite":            0xffdead,
	"moccasin":               0xffe4b5,
	"cornsilk":               0xfff8dc,
	"ivory":                  0xfffff0,
	"lemon chiffon":          0xfffacd,
	"lemonchiffon":           0xfffacd,
	"seashell":               0xfff5ee,
	"honeydew":               0xf0fff0,
	"mint cream":             0xf5fffa,
	"mintcream":              0xf5fffa,
	"azure":                  0xf0ffff,
	"alice blue":             0xf0f8ff,
	"aliceblue":              0xf0f8ff,
	"lavender":               0xe6e6fa,
	"lavender blush":         0xfff0f5,
	"lavenderblush":          0xfff0f5,
	"misty rose":             0xffe4e1,
	"mistyrose":              0xffe4e1,
	"white":                  0xffffff,
	"black":                  0x0,
	"dark slate gray":        0x2f4f4f,
	"darkslategray":          0x2f4f4f,
	"dark slate grey":        0x2f4f4f,
	"darkslategrey":          0x2f4f4f,
	"dim gray":               0x696969,
	"dimgray":                0x696969,
	"dim grey":               0x696969,
	"dimgrey":                0x696969,
	"slate gray":             0x708090,
	"slategray":              0x708090,
	"slate grey":             0x708090,
	"slategrey":              0x708090,
	"light slate gray":       0x778899,
	"lightslategray":         0x778899,
	"light slate grey":       0x778899,
	"lightslategrey":         0x778899,
	"gray":                   0xbebebe,
	"grey":                   0xbebebe,
	"light grey":             0xd3d3d3,
	"lightgrey":              0xd3d3d3,
	"light gray":             0xd3d3d3,
	"lightgray":              0xd3d3d3,
	"midnight blue":          0x191970,
	"midnightblue":           0x191970,
	"navy":                   0x80,
	"navy blue":              0x80,
	"navyblue":               0x80,
	"cornflower blue":        0x6495ed,
	"cornflowerblue":         0x6495ed,
	"dark slate blue":        0x483d8b,
	"darkslateblue":          0x483d8b,
	"slate blue":             0x6a5acd,
	"slateblue":              0x6a5acd,
	"medium slate blue":      0x7b68ee,
	"mediumslateblue":        0x7b68ee,
	"light slate blue":       0x8470ff,
	"lightslateblue":         0x8470ff,
	"medium blue":            0xcd,
	"mediumblue":             0xcd,
	"royal blue":             0x4169e1,
	"royalblue":              0x4169e1,
	"blue":                   0xff,
	"dodger blue":            0x1e90ff,
	"dodgerblue":             0x1e90ff,
	"deep sky blue":          0xbfff,
	"deepskyblue":            0xbfff,
	"sky blue":               0x87ceeb,
	"skyblue":                0x87ceeb,
	"light sky blue":         0x87cefa,
	"lightskyblue":           0x87cefa,
	"steel blue":             0x4682b4,
	"steelblue":              0x4682b4,
	"light steel blue":       0xb0c4de,
	"lightsteelblue":         0xb0c4de,
	"light blue":             0xadd8e6,
	"lightblue":              0xadd8e6,
	"powder blue":            0xb0e0e6,
	"powderblue":             0xb0e0e6,
	"pale turquoise":         0xafeeee,
	"paleturquoise":          0xafeeee,
	"dark turquoise":         0xced1,
	"darkturquoise":          0xced1,
	"medium turquoise":       0x48d1cc,
	"mediumturquoise":        0x48d1cc,
	"turquoise":              0x40e0d0,
	"cyan":                   0xffff,
	"light cyan":             0xe0ffff,
	"lightcyan":              0xe0ffff,
	"cadet blue":             0x5f9ea0,
	"cadetblue":              0x5f9ea0,
	"medium aquamarine":      0x66cdaa,
	"mediumaquamarine":       0x66cdaa,
	"aquamarine":             0x7fffd4,
	"dark green":             0x6400,
	"darkgreen":              0x6400,
	"dark olive green":       0x556b2f,
	"darkolivegreen":         0x556b2f,
	"dark sea green":         0x8fbc8f,
	"darkseagreen":           0x8fbc8f,
	"sea green":              0x2e8b57,
	"seagreen":               0x2e8b57,
	"medium sea green":       0x3cb371,
	"mediumseagreen":         0x3cb371,
	"light sea green":        0x20b2aa,
	"lightseagreen":          0x20b2aa,
	"pale green":             0x98fb98,
	"palegreen":              0x98fb98,
	"spring green":           0xff7f,
	"springgreen":            0xff7f,
	"lawn green":             0x7cfc00,
	"lawngreen":              0x7cfc00,
	"green":                  0xff00,
	"chartreuse":             0x7fff00,
	"medium spring green":    0xfa9a,
	"mediumspringgreen":      0xfa9a,
	"green yellow":           0xadff2f,
	"greenyellow":            0xadff2f,
	"lime green":             0x32cd32,
	"limegreen":              0x32cd32,
	"yellow green":           0x9acd32,
	"yellowgreen":            0x9acd32,
	"forest green":           0x228b22,
	"forestgreen":            0x228b22,
	"olive drab":             0x6b8e23,
	"olivedrab":              0x6b8e23,
	"dark khaki":             0xbdb76b,
	"darkkhaki":              0xbdb76b,
	"khaki":                  0xf0e68c,
	"pale goldenrod":         0xeee8aa,
	"palegoldenrod":          0xeee8aa,
	"light goldenrod yellow": 0xfafad2,
	"lightgoldenrodyellow":   0xfafad2,
	"light yellow":           0xffffe0,
	"lightyellow":            0xffffe0,
	"yellow":                 0xffff00,
	"gold":                   0xffd700,
	"light goldenrod":        0xeedd82,
	"lightgoldenrod":         0xeedd82,
	"goldenrod":              0xdaa520,
	"dark goldenrod":         0xb8860b,
	"darkgoldenrod":          0xb8860b,
	"rosy brown":             0xbc8f8f,
	"rosybrown":              0xbc8f8f,
	"indian red":             0xcd5c5c,
	"indianred":              0xcd5c5c,
	"saddle brown":           0x8b4513,
	"saddlebrown":            0x8b4513,
	"sienna":                 0xa0522d,
	"peru":                   0xcd853f,
	"burlywood":              0xdeb887,
	"beige":                  0xf5f5dc,
	"wheat":                  0xf5deb3,
	"sandy brown":            0xf4a460,
	"sandybrown":             0xf4a460,
	"tan":                    0xd2b48c,
	"chocolate":              0xd2691e,
	"firebrick":              0xb22222,
	"brown":                  0xa52a2a,
	"dark salmon":            0xe9967a,
	"darksalmon":             0xe9967a,
	"salmon":                 0xfa8072,
	"light salmon":           0xffa07a,
	"lightsalmon":            0xffa07a,
	"orange":                 0xffa500,
	"dark orange":            0xff8c00,
	"darkorange":             0xff8c00,
	"coral":                  0xff7f50,
	"light coral":            0xf08080,
	"lightcoral":             0xf08080,
	"tomato":                 0xff6347,
	"orange red":             0xff4500,
	"orangered":              0xff4500,
	"red":                    0xff0000,
	"hot pink":               0xff69b4,
	"hotpink":                0xff69b4,
	"deep pink":              0xff1493,
	"deeppink":               0xff1493,
	"pink":                   0xffc0cb,
	"light pink":             0xffb6c1,
	"lightpink":              0xffb6c1,
	"pale violet red":        0xdb7093,
	"palevioletred":          0xdb7093,
	"maroon":                 0xb03060,
	"medium violet red":      0xc71585,
	"mediumvioletred":        0xc71585,
	"violet red":             0xd02090,
	"violetred":              0xd02090,
	"magenta":                0xff00ff,
	"violet":                 0xee82ee,
	"plum":                   0xdda0dd,
	"orchid":                 0xda70d6,
	"medium orchid":          0xba55d3,
	"mediumorchid":           0xba55d3,
	"dark orchid":            0x9932cc,
	"darkorchid":             0x9932cc,
	"dark violet":            0x9400d3,
	"darkviolet":             0x9400d3,
	"blue violet":            0x8a2be2,
	"blueviolet":             0x8a2be2,
	"purple":                 0xa020f0,
	"medium purple":          0x9370db,
	"mediumpurple":           0x9370db,
	"thistle":                0xd8bfd8,
	"snow1":                  0xfffafa,
	"snow2":                  0xeee9e9,
	"snow3":                  0xcdc9c9,
	"snow4":                  0x8b8989,
	"seashell1":              0xfff5ee,
	"seashell2":              0xeee5de,
	"seashell3":              0xcdc5bf,
	"seashell4":              0x8b8682,
	"antiquewhite1":          0xffefdb,
	"antiquewhite2":          0xeedfcc,
	"antiquewhite3":          0xcdc0b0,
	"antiquewhite4":          0x8b8378,
	"bisque1":                0xffe4c4,
	"bisque2":                0xeed5b7,
	"bisque3":                0xcdb79e,
	"bisque4":                0x8b7d6b,
	"peachpuff1":             0xffdab9,
	"peachpuff2":             0xeecbad,
	"peachpuff3":             0xcdaf95,
	"peachpuff4":             0x8b7765,
	"navajowhite1":           0xffdead,
	"navajowhite2":           0xeecfa1,
	"navajowhite3":           0xcdb38b,
	"navajowhite4":           0x8b795e,
	"lemonchiffon1":          0xfffacd,
	"lemonchiffon2":          0xeee9bf,
	"lemonchiffon3":          0xcdc9a5,
	"lemonchiffon4":          0x8b8970,
	"cornsilk1":              0xfff8dc,
	"cornsilk2":              0xeee8cd,
	"cornsilk3":              0xcdc8b1,
	"cornsilk4":              0x8b8878,
	"ivory1":                 0xfffff0,
	"ivory2":                 0xeeeee0,
	"ivory3":                 0xcdcdc1,
	"ivory4":                 0x8b8b83,
	"honeydew1":              0xf0fff0,
	"honeydew2":              0xe0eee0,
	"honeydew3":              0xc1cdc1,
	"honeydew4":              0x838b83,
	"lavenderblush1":         0xfff0f5,
	"lavenderblush2":         0xeee0e5,
	"lavenderblush3":         0xcdc1c5,
	"lavenderblush4":         0x8b8386,
	"mistyrose1":             0xffe4e1,
	"mistyrose2":             0xeed5d2,
	"mistyrose3":             0xcdb7b5,
	"mistyrose4":             0x8b7d7b,
	"azure1":                 0xf0ffff,
	"azure2":                 0xe0eeee,
	"azure3":                 0xc1cdcd,
	"azure4":                 0x838b8b,
	"slateblue1":             0x836fff,
	"slateblue2":             0x7a67ee,
	"slateblue3":             0x6959cd,
	"slateblue4":             0x473c8b,
	"royalblue1":             0x4876ff,
	"royalblue2":             0x436eee,
	"royalblue3":             0x3a5fcd,
	"royalblue4":             0x27408b,
	"blue1":                  0xff,
	"blue2":                  0xee,
	"blue3":                  0xcd,
	"blue4":                  0x8b,
	"dodgerblue1":            0x1e90ff,
	"dodgerblue2":            0x1c86ee,
	"dodgerblue3":            0x1874cd,
	"dodgerblue4":            0x104e8b,
	"steelblue1":             0x63b8ff,
	"steelblue2":             0x5cacee,
	"steelblue3":             0x4f94cd,
	"steelblue4":             0x36648b,
	"deepskyblue1":           0xbfff,
	"deepskyblue2":           0xb2ee,
	"deepskyblue3":           0x9acd,
	"deepskyblue4":           0x688b,
	"skyblue1":               0x87ceff,
	"skyblue2":               0x7ec0ee,
	"skyblue3":               0x6ca6cd,
	"skyblue4":               0x4a708b,
	"lightskyblue1":          0xb0e2ff,
	"lightskyblue2":          0xa4d3ee,
	"lightskyblue3":          0x8db6cd,
	"lightskyblue4":          0x607b8b,
	"slategray1":             0xc6e2ff,
	"slategray2":             0xb9d3ee,
	"slategray3":             0x9fb6cd,
	"slategray4":             0x6c7b8b,
	"lightsteelblue1":        0xcae1ff,
	"lightsteelblue2":        0xbcd2ee,
	"lightsteelblue3":        0xa2b5cd,
	"lightsteelblue4":        0x6e7b8b,
	"lightblue1":             0xbfefff,
	"lightblue2":             0xb2dfee,
	"lightblue3":             0x9ac0cd,
	"lightblue4":             0x68838b,
	"lightcyan1":             0xe0ffff,
	"lightcyan2":             0xd1eeee,
	"lightcyan3":             0xb4cdcd,
	"lightcyan4":             0x7a8b8b,
	"paleturquoise1":         0xbbffff,
	"paleturquoise2":         0xaeeeee,
	"paleturquoise3":         0x96cdcd,
	"paleturquoise4":         0x668b8b,
	"cadetblue1":             0x98f5ff,
	"cadetblue2":             0x8ee5ee,
	"cadetblue3":             0x7ac5cd,
	"cadetblue4":             0x53868b,
	"turquoise1":             0xf5ff,
	"turquoise2":             0xe5ee,
	"turquoise3":             0xc5cd,
	"turquoise4":             0x868b,
	"cyan1":                  0xffff,
	"cyan2":                  0xeeee,
	"cyan3":                  0xcdcd,
	"cyan4":                  0x8b8b,
	"darkslategray1":         0x97ffff,
	"darkslategray2":         0x8deeee,
	"darkslategray3":         0x79cdcd,
	"darkslategray4":         0x528b8b,
	"aquamarine1":            0x7fffd4,
	"aquamarine2":            0x76eec6,
	"aquamarine3":            0x66cdaa,
	"aquamarine4":            0x458b74,
	"darkseagreen1":          0xc1ffc1,
	"darkseagreen2":          0xb4eeb4,
	"darkseagreen3":          0x9bcd9b,
	"darkseagreen4":          0x698b69,
	"seagreen1":              0x54ff9f,
	"seagreen2":              0x4eee94,
	"seagreen3":              0x43cd80,
	"seagreen4":              0x2e8b57,
	"palegreen1":             0x9aff9a,
	"palegreen2":             0x90ee90,
	"palegreen3":             0x7ccd7c,
	"palegreen4":             0x548b54,
	"springgreen1":           0xff7f,
	"springgreen2":           0xee76,
	"springgreen3":           0xcd66,
	"springgreen4":           0x8b45,
	"green1":                 0xff00,
	"green2":                 0xee00,
	"green3":                 0xcd00,
	"green4":                 0x8b00,
	"chartreuse1":            0x7fff00,
	"chartreuse2":            0x76ee00,
	"chartreuse3":            0x66cd00,
	"chartreuse4":            0x458b00,
	"olivedrab1":             0xc0ff3e,
	"olivedrab2":             0xb3ee3a,
	"olivedrab3":             0x9acd32,
	"olivedrab4":             0x698b22,
	"darkolivegreen1":        0xcaff70,
	"darkolivegreen2":        0xbcee68,
	"darkolivegreen3":        0xa2cd5a,
	"darkolivegreen4":        0x6e8b3d,
	"khaki1":                 0xfff68f,
	"khaki2":                 0xeee685,
	"khaki3":                 0xcdc673,
	"khaki4":                 0x8b864e,
	"lightgoldenrod1":        0xffec8b,
	"lightgoldenrod2":        0xeedc82,
	"lightgoldenrod3":        0xcdbe70,
	"lightgoldenrod4":        0x8b814c,
	"lightyellow1":           0xffffe0,
	"lightyellow2":           0xeeeed1,
	"lightyellow3":           0xcdcdb4,
	"lightyellow4":           0x8b8b7a,
	"yellow1":                0xffff00,
	"yellow2":                0xeeee00,
	"yellow3":                0xcdcd00,
	"yellow4":                0x8b8b00,
	"gold1":                  0xffd700,
	"gold2":                  0xeec900,
	"gold3":                  0xcdad00,
	"gold4":                  0x8b7500,
	"goldenrod1":             0xffc125,
	"goldenrod2":             0xeeb422,
	"goldenrod3":             0xcd9b1d,
	"goldenrod4":             0x8b6914,
	"darkgoldenrod1":         0xffb90f,
	"darkgoldenrod2":         0xeead0e,
	"darkgoldenrod3":         0xcd950c,
	"darkgoldenrod4":         0x8b6508,
	"rosybrown1":             0xffc1c1,
	"rosybrown2":             0xeeb4b4,
	"rosybrown3":             0xcd9b9b,
	"rosybrown4":             0x8b6969,
	"indianred1":             0xff6a6a,
	"indianred2":             0xee6363,
	"indianred3":             0xcd5555,
	"indianred4":             0x8b3a3a,
	"sienna1":                0xff8247,
	"sienna2":                0xee7942,
	"sienna3":                0xcd6839,
	"sienna4":                0x8b4726,
	"burlywood1":             0xffd39b,
	"burlywood2":             0xeec591,
	"burlywood3":             0xcdaa7d,
	"burlywood4":             0x8b7355,
	"wheat1":                 0xffe7ba,
	"wheat2":                 0xeed8ae,
	"wheat3":                 0xcdba96,
	"wheat4":                 0x8b7e66,
	"tan1":                   0xffa54f,
	"tan2":                   0xee9a49,
	"tan3":                   0xcd853f,
	"tan4":                   0x8b5a2b,
	"chocolate1":             0xff7f24,
	"chocolate2":             0xee7621,
	"chocolate3":             0xcd661d,
	"chocolate4":             0x8b4513,
	"firebrick1":             0xff3030,
	"firebrick2":             0xee2c2c,
	"firebrick3":             0xcd2626,
	"firebrick4":             0x8b1a1a,
	"brown1":                 0xff4040,
	"brown2":                 0xee3b3b,
	"brown3":                 0xcd3333,
	"brown4":                 0x8b2323,
	"salmon1":                0xff8c69,
	"salmon2":                0xee8262,
	"salmon3":                0xcd7054,
	"salmon4":                0x8b4c39,
	"lightsalmon1":           0xffa07a,
	"lightsalmon2":           0xee9572,
	"lightsalmon3":           0xcd8162,
	"lightsalmon4":           0x8b5742,
	"orange1":                0xffa500,
	"orange2":                0xee9a00,
	"orange3":                0xcd8500,
	"orange4":                0x8b5a00,
	"darkorange1":            0xff7f00,
	"darkorange2":            0xee7600,
	"darkorange3":            0xcd6600,
	"darkorange4":            0x8b4500,
	"coral1":                 0xff7256,
	"coral2":                 0xee6a50,
	"coral3":                 0xcd5b45,
	"coral4":                 0x8b3e2f,
	"tomato1":                0xff6347,
	"tomato2":                0xee5c42,
	"tomato3":                0xcd4f39,
	"tomato4":                0x8b3626,
	"orangered1":             0xff4500,
	"orangered2":             0xee4000,
	"orangered3":             0xcd3700,
	"orangered4":             0x8b2500,
	"red1":                   0xff0000,
	"red2":                   0xee0000,
	"red3":                   0xcd0000,
	"red4":                   0x8b0000,
	"deeppink1":              0xff1493,
	"deeppink2":              0xee1289,
	"deeppink3":              0xcd1076,
	"deeppink4":              0x8b0a50,
	"hotpink1":               0xff6eb4,
	"hotpink2":               0xee6aa7,
	"hotpink3":               0xcd6090,
	"hotpink4":               0x8b3a62,
	"pink1":                  0xffb5c5,
	"pink2":                  0xeea9b8,
	"pink3":                  0xcd919e,
	"pink4":                  0x8b636c,
	"lightpink1":             0xffaeb9,
	"lightpink2":             0xeea2ad,
	"lightpink3":             0xcd8c95,
	"lightpink4":             0x8b5f65,
	"palevioletred1":         0xff82ab,
	"palevioletred2":         0xee799f,
	"palevioletred3":         0xcd6889,
	"palevioletred4":         0x8b475d,
	"maroon1":                0xff34b3,
	"maroon2":                0xee30a7,
	"maroon3":                0xcd2990,
	"maroon4":                0x8b1c62,
	"violetred1":             0xff3e96,
	"violetred2":             0xee3a8c,
	"violetred3":             0xcd3278,
	"violetred4":             0x8b2252,
	"magenta1":               0xff00ff,
	"magenta2":               0xee00ee,
	"magenta3":               0xcd00cd,
	"magenta4":               0x8b008b,
	"orchid1":                0xff83fa,
	"orchid2":                0xee7ae9,
	"orchid3":                0xcd69c9,
	"orchid4":                0x8b4789,
	"plum1":                  0xffbbff,
	"plum2":                  0xeeaeee,
	"plum3":                  0xcd96cd,
	"plum4":                  0x8b668b,
	"mediumorchid1":          0xe066ff,
	"mediumorchid2":          0xd15fee,
	"mediumorchid3":          0xb452cd,
	"mediumorchid4":          0x7a378b,
	"darkorchid1":            0xbf3eff,
	"darkorchid2":            0xb23aee,
	"darkorchid3":            0x9a32cd,
	"darkorchid4":            0x68228b,
	"purple1":                0x9b30ff,
	"purple2":                0x912cee,
	"purple3":                0x7d26cd,
	"purple4":                0x551a8b,
	"mediumpurple1":          0xab82ff,
	"mediumpurple2":          0x9f79ee,
	"mediumpurple3":          0x8968cd,
	"mediumpurple4":          0x5d478b,
	"thistle1":               0xffe1ff,
	"thistle2":               0xeed2ee,
	"thistle3":               0xcdb5cd,
	"thistle4":               0x8b7b8b,
	"gray0":                  0x0,
	"grey0":                  0x0,
	"gray1":                  0x30303,
	"grey1":                  0x30303,
	"gray2":                  0x50505,
	"grey2":                  0x50505,
	"gray3":                  0x80808,
	"grey3":                  0x80808,
	"gray4":                  0xa0a0a,
	"grey4":                  0xa0a0a,
	"gray5":                  0xd0d0d,
	"grey5":                  0xd0d0d,
	"gray6":                  0xf0f0f,
	"grey6":                  0xf0f0f,
	"gray7":                  0x121212,
	"grey7":                  0x121212,
	"gray8":                  0x141414,
	"grey8":                  0x141414,
	"gray9":                  0x171717,
	"grey9":                  0x171717,
	"gray10":                 0x1a1a1a,
	"grey10":                 0x1a1a1a,
	"gray11":                 0x1c1c1c,
	"grey11":                 0x1c1c1c,
	"gray12":                 0x1f1f1f,
	"grey12":                 0x1f1f1f,
	"gray13":                 0x212121,
	"grey13":                 0x212121,
	"gray14":                 0x242424,
	"grey14":                 0x242424,
	"gray15":                 0x262626,
	"grey15":                 0x262626,
	"gray16":                 0x292929,
	"grey16":                 0x292929,
	"gray17":                 0x2b2b2b,
	"grey17":                 0x2b2b2b,
	"gray18":                 0x2e2e2e,
	"grey18":                 0x2e2e2e,
	"gray19":                 0x303030,
	"grey19":                 0x303030,
	"gray20":                 0x333333,
	"grey20":                 0x333333,
	"gray21":                 0x363636,
	"grey21":                 0x363636,
	"gray22":                 0x383838,
	"grey22":                 0x383838,
	"gray23":                 0x3b3b3b,
	"grey23":                 0x3b3b3b,
	"gray24":                 0x3d3d3d,
	"grey24":                 0x3d3d3d,
	"gray25":                 0x404040,
	"grey25":                 0x404040,
	"gray26":                 0x424242,
	"grey26":                 0x424242,
	"gray27":                 0x454545,
	"grey27":                 0x454545,
	"gray28":                 0x474747,
	"grey28":                 0x474747,
	"gray29":                 0x4a4a4a,
	"grey29":                 0x4a4a4a,
	"gray30":                 0x4d4d4d,
	"grey30":                 0x4d4d4d,
	"gray31":                 0x4f4f4f,
	"grey31":                 0x4f4f4f,
	"gray32":                 0x525252,
	"grey32":                 0x525252,
	"gray33":                 0x545454,
	"grey33":                 0x545454,
	"gray34":                 0x575757,
	"grey34":                 0x575757,
	"gray35":                 0x595959,
	"grey35":                 0x595959,
	"gray36":                 0x5c5c5c,
	"grey36":                 0x5c5c5c,
	"gray37":                 0x5e5e5e,
	"grey37":                 0x5e5e5e,
	"gray38":                 0x616161,
	"grey38":                 0x616161,
	"gray39":                 0x636363,
	"grey39":                 0x636363,
	"gray40":                 0x666666,
	"grey40":                 0x666666,
	"gray41":                 0x696969,
	"grey41":                 0x696969,
	"gray42":                 0x6b6b6b,
	"grey42":                 0x6b6b6b,
	"gray43":                 0x6e6e6e,
	"grey43":                 0x6e6e6e,
	"gray44":                 0x707070,
	"grey44":                 0x707070,
	"gray45":                 0x737373,
	"grey45":                 0x737373,
	"gray46":                 0x757575,
	"grey46":                 0x757575,
	"gray47":                 0x787878,
	"grey47":                 0x787878,
	"gray48":                 0x7a7a7a,
	"grey48":                 0x7a7a7a,
	"gray49":                 0x7d7d7d,
	"grey49":                 0x7d7d7d,
	"gray50":                 0x7f7f7f,
	"grey50":                 0x7f7f7f,
	"gray51":                 0x828282,
	"grey51":                 0x828282,
	"gray52":                 0x858585,
	"grey52":                 0x858585,
	"gray53":                 0x878787,
	"grey53":                 0x878787,
	"gray54":                 0x8a8a8a,
	"grey54":                 0x8a8a8a,
	"gray55":                 0x8c8c8c,
	"grey55":                 0x8c8c8c,
	"gray56":                 0x8f8f8f,
	"grey56":                 0x8f8f8f,
	"gray57":                 0x919191,
	"grey57":                 0x919191,
	"gray58":                 0x949494,
	"grey58":                 0x949494,
	"gray59":                 0x969696,
	"grey59":                 0x969696,
	"gray60":                 0x999999,
	"grey60":                 0x999999,
	"gray61":                 0x9c9c9c,
	"grey61":                 0x9c9c9c,
	"gray62":                 0x9e9e9e,
	"grey62":                 0x9e9e9e,
	"gray63":                 0xa1a1a1,
	"grey63":                 0xa1a1a1,
	"gray64":                 0xa3a3a3,
	"grey64":                 0xa3a3a3,
	"gray65":                 0xa6a6a6,
	"grey65":                 0xa6a6a6,
	"gray66":                 0xa8a8a8,
	"grey66":                 0xa8a8a8,
	"gray67":                 0xababab,
	"grey67":                 0xababab,
	"gray68":                 0xadadad,
	"grey68":                 0xadadad,
	"gray69":                 0xb0b0b0,
	"grey69":                 0xb0b0b0,
	"gray70":                 0xb3b3b3,
	"grey70":                 0xb3b3b3,
	"gray71":                 0xb5b5b5,
	"grey71":                 0xb5b5b5,
	"gray72":                 0xb8b8b8,
	"grey72":                 0xb8b8b8,
	"gray73":                 0xbababa,
	"grey73":                 0xbababa,
	"gray74":                 0xbdbdbd,
	"grey74":                 0xbdbdbd,
	"gray75":                 0xbfbfbf,
	"grey75":                 0xbfbfbf,
	"gray76":                 0xc2c2c2,
	"grey76":                 0xc2c2c2,
	"gray77":                 0xc4c4c4,
	"grey77":                 0xc4c4c4,
	"gray78":                 0xc7c7c7,
	"grey78":                 0xc7c7c7,
	"gray79":                 0xc9c9c9,
	"grey79":                 0xc9c9c9,
	"gray80":                 0xcccccc,
	"grey80":                 0xcccccc,
	"gray81":                 0xcfcfcf,
	"grey81":                 0xcfcfcf,
	"gray82":                 0xd1d1d1,
	"grey82":                 0xd1d1d1,
	"gray83":                 0xd4d4d4,
	"grey83":                 0xd4d4d4,
	"gray84":                 0xd6d6d6,
	"grey84":                 0xd6d6d6,
	"gray85":                 0xd9d9d9,
	"grey85":                 0xd9d9d9,
	"gray86":                 0xdbdbdb,
	"grey86":                 0xdbdbdb,
	"gray87":                 0xdedede,
	"grey87":                 0xdedede,
	"gray88":                 0xe0e0e0,
	"grey88":                 0xe0e0e0,
	"gray89":                 0xe3e3e3,
	"grey89":                 0xe3e3e3,
	"gray90":                 0xe5e5e5,
	"grey90":                 0xe5e5e5,
	"gray91":                 0xe8e8e8,
	"grey91":                 0xe8e8e8,
	"gray92":                 0xebebeb,
	"grey92":                 0xebebeb,
	"gray93":                 0xededed,
	"grey93":                 0xededed,
	"gray94":                 0xf0f0f0,
	"grey94":                 0xf0f0f0,
	"gray95":                 0xf2f2f2,
	"grey95":                 0xf2f2f2,
	"gray96":                 0xf5f5f5,
	"grey96":                 0xf5f5f5,
	"gray97":                 0xf7f7f7,
	"grey97":                 0xf7f7f7,
	"gray98":                 0xfafafa,
	"grey98":                 0xfafafa,
	"gray99":                 0xfcfcfc,
	"grey99":                 0xfcfcfc,
	"gray100":                0xffffff,
	"grey100":                0xffffff,
	"dark grey":              0xa9a9a9,
	"darkgrey":               0xa9a9a9,
	"dark gray":              0xa9a9a9,
	"darkgray":               0xa9a9a9,
	"dark blue":              0x8b,
	"darkblue":               0x8b,
	"dark cyan":              0x8b8b,
	"darkcyan":               0x8b8b,
	"dark magenta":           0x8b008b,
	"darkmagenta":            0x8b008b,
	"dark red":               0x8b0000,
	"darkred":                0x8b0000,
	"light green":            0x90ee90,
	"lightgreen":             0x90ee90,
}

type Gap struct {
	Top, Bottom, Left, Right int
}
type ClientSpec struct {
	Name  string
	Class string
}
type KeySpec struct {
	Mods string
	Key  string
}

func (k KeySpec) ToXGB() string {
	var out []string
	for _, c := range k.Mods {
		switch c {
		case 'C':
			out = append(out, "Control")
		case 'M':
			out = append(out, "Mod1")
		case 'S':
			out = append(out, "Shift")
		case '4':
			out = append(out, "Mod4")
		}
	}
	out = append(out, k.Key)
	return strings.Join(out, "-")
}

type Config struct {
	BorderWidth int
	Snapdist    int
	Colors      map[string]int
	Gap         Gap
	Autogroups  map[ClientSpec]int
	Binds       map[KeySpec]string
	Commands    map[string]string
	Font        string // FIXME will we support Xft?
	Ignores     []string
	MouseBinds  map[string]KeySpec
	MoveAmount  int // default: 1
	Sticky      bool
}

type parseDecl struct {
	num int
	fn  func(cfg *Config, in []string) error
}

var parseMap = map[string]parseDecl{
	"autogroup": {2, func(cfg *Config, in []string) error {
		group, err := strconv.Atoi(in[0])
		if err != nil {
			return err
		}
		parts := strings.SplitN(in[1], ".", 2)
		var cs ClientSpec
		switch len(parts) {
		case 1:
			cs = ClientSpec{Class: parts[0]}
		case 2:
			cs = ClientSpec{Name: parts[0], Class: parts[1]}
		default:
			return fmt.Errorf("invalid clientspec %q", in[1])
		}
		cfg.Autogroups[cs] = group
		return nil
	}},

	"bind": {2, func(cfg *Config, in []string) error {
		parts := strings.SplitN(in[0], "-", 2)
		var key KeySpec
		switch len(parts) {
		case 1:
			key = KeySpec{Key: parts[0]}
		case 2:
			// TODO make sure the mod is valid
			key = KeySpec{Mods: parts[0], Key: parts[1]}
		default:
			return fmt.Errorf("invalid keyspec %q", in[0])
		}
		if in[1] == "unmap" {
			delete(cfg.Binds, key)
		} else {
			cfg.Binds[key] = in[1]
		}
		return nil
	}},

	"borderwidth": {1, func(cfg *Config, in []string) error {
		i, err := strconv.Atoi(in[0])
		if err != nil {
			return err
		}
		cfg.BorderWidth = i
		return nil
	}},

	"color": {2, func(cfg *Config, in []string) error {
		color, ok := colors[strings.ToLower(in[1])]
		if !ok {
			return fmt.Errorf("unknown color name %q", in[1])
		}
		cfg.Colors[in[0]] = color
		return nil
	}},

	"command": {2, func(cfg *Config, in []string) error {
		cfg.Commands[in[0]] = in[1]
		return nil
	}},

	"fontname": {1, func(cfg *Config, in []string) error {
		cfg.Font = in[0]
		return nil
	}},

	"gap": {4, func(cfg *Config, in []string) error {
		i, err := parseInts(in)
		if err != nil {
			return err
		}
		cfg.Gap.Top = i[0]
		cfg.Gap.Bottom = i[1]
		cfg.Gap.Left = i[2]
		cfg.Gap.Right = i[3]
		return nil
	}},

	"ignore": {1, func(cfg *Config, in []string) error {
		cfg.Ignores = append(cfg.Ignores, in[0])
		return nil
	}},

	"mousebind": {2, func(cfg *Config, in []string) error {
		parts := strings.SplitN(in[0], "-", 2)
		var key KeySpec
		switch len(parts) {
		case 1:
			key = KeySpec{Key: parts[0]}
		case 2:
			key = KeySpec{Mods: parts[0], Key: parts[1]}
		default:
			return fmt.Errorf("invalid mousepec %q", in[0])
		}
		if in[1] == "unmap" {
			for k, v := range cfg.MouseBinds {
				if v == key {
					delete(cfg.MouseBinds, k)
					break
				}
			}
		} else {
			cfg.MouseBinds[in[1]] = key
		}
		return nil
	}},

	"moveamount": {1, func(cfg *Config, in []string) error {
		i, err := strconv.Atoi(in[0])
		if err != nil {
			return err
		}
		cfg.MoveAmount = i
		return nil
	}},

	"snapdist": {1, func(cfg *Config, in []string) error {
		i, err := strconv.Atoi(in[0])
		if err != nil {
			return err
		}
		cfg.Snapdist = i
		return nil
	}},

	"sticky": {1, func(cfg *Config, in []string) error {
		switch in[0] {
		case "yes":
			cfg.Sticky = true
		case "no":
			cfg.Sticky = false
		default:
			return fmt.Errorf("invalid value %q for sticky", in[0])
		}
		return nil
	}},
}

func Parse(r io.Reader) (*Config, error) {
	cfg := &Config{}
	cfg.Autogroups = make(map[ClientSpec]int)
	cfg.Binds = make(map[KeySpec]string)
	cfg.Colors = make(map[string]int)
	cfg.Commands = make(map[string]string)
	cfg.MouseBinds = make(map[string]KeySpec)
	cfg.MoveAmount = 1

	cnt, _ := ioutil.ReadAll(r)
	_, ch := lex(string(cnt))
	for {
		command, ok := <-ch
		if !ok {
			return cfg, errors.New("internal error")
		}

		if command.typ == itemEOF {
			return cfg, nil
		}
		if command.typ != itemString {
			return cfg, errors.New("unexpected token " + command.String())
		}
		decl, ok := parseMap[command.val]
		if !ok {
			return cfg, errors.New("unknown option " + command.val)
		}
		in, err := expect(ch, decl.num)
		if err != nil {
			return cfg, err
		}
		err = decl.fn(cfg, in)
		if err != nil {
			return cfg, err
		}
	}
}

func parseInts(in []string) ([]int, error) {
	out := make([]int, len(in))
	var err error
	for i, s := range in {
		out[i], err = strconv.Atoi(s)
		if err != nil {
			return out, err
		}
	}
	return out, nil
}

func expect(ch chan item, num int) ([]string, error) {
	var ret []string
	for i := 0; i < num; i++ {
		val := <-ch
		if val.typ == itemError {
			return ret, errors.New(val.val)
		}

		if val.typ == itemTerminator || val.typ == itemEOF {
			return ret, io.EOF
		}

		ret = append(ret, val.val)
	}

	val := <-ch
	if val.typ != itemTerminator {
		return ret, errors.New("unexpected token " + val.typ.String())
	}

	return ret, nil
}

type lexer struct {
	input             string
	start             int
	pos               int
	width             int
	items             chan item
	lastWasTerminator bool
}
type itemType int

const (
	itemError itemType = iota
	itemString
	itemTerminator
	itemEOF
)

func (i itemType) String() string {
	switch i {
	case itemError:
		return "error"
	case itemString:
		return "string"
	case itemTerminator:
		return "terminator"
	case itemEOF:
		return "eof"
	default:
		return ""
	}
}

const eof = -1

type item struct {
	typ itemType
	val string
}

func (i item) String() string {
	switch i.typ {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.val
	}
	return fmt.Sprintf("(%s) %q", i.typ, i.val)
}

type stateFn func(*lexer) stateFn

func lex(input string) (*lexer, chan item) {
	l := &lexer{
		input: input,
		items: make(chan item),
	}
	go l.run()
	return l, l.items
}

func (l *lexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func (l *lexer) emit(t itemType) {
	l.lastWasTerminator = t == itemTerminator
	l.items <- item{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) next() (rune rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	rune, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return rune
}

func (l *lexer) ignore() {
	l.start = l.pos
}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}

func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, fmt.Sprintf(format, args...)}
	return nil
}

func lexText(l *lexer) stateFn {
	for {
		r := l.next()
		if r == eof {
			break
		}

		if r == '#' {
			return lexComment
		}

		// TODO is this right?
		if r == ' ' || r == '\t' {
			l.ignore()
			continue
		}

		if r == '\n' {
			if l.lastWasTerminator {
				l.ignore()
			} else {
				l.emit(itemTerminator)
			}
		}

		return lexString

	}
	l.emit(itemEOF)
	return nil
}

func lexString(l *lexer) stateFn {
	defer func() {
		if l.input[l.start:l.pos] != "" {
			l.emit(itemString)
		}
	}()
	quoted := false
	if l.input[l.pos-1] == '"' {
		quoted = true
	}
	escape := false
	multiline := false

	var r rune
loop:
	for r != eof {
		r = l.next()
		switch r {
		case '\\':
			if quoted {
				escape = !escape
			} else {
				multiline = true
			}
			// TODO multiline string
		case '"':
			if quoted && !escape {
				break loop
			}
		case ' ', '\t':
			if !quoted {
				l.backup()
				break loop
			}
		case '\n':
			if quoted || multiline {
				multiline = false
			} else {
				l.backup()
				break loop
			}
		case '#':
			if !quoted {
				l.backup()

				return lexComment
			}
		}
	}

	return lexText
}

func lexComment(l *lexer) stateFn {
	for {
		r := l.next()
		if r == eof || r == '\n' {
			l.backup()
			// TODO can comments be multiline?
			break
		}
	}
	l.ignore()
	return lexText
}

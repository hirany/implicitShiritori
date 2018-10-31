package main

// すでにあるかどうかの判定
func userdata(username string) bool {
	if username == "cpu" {
		return true
	}
	return false
}

func voc(username string, uservoc [][]string) {
	if username == "cpu" {
		uservoc[0][0] = "a"
		uservoc[0][1] = "ab"
		uservoc[1][0] = "b"
		uservoc[1][1] = "bc"
		uservoc[2][0] = "c"
		uservoc[2][1] = "cd"
		uservoc[3][0] = "d"
		uservoc[3][1] = "de"
		uservoc[4][0] = "e"
		uservoc[4][1] = "ef"
		uservoc[5][0] = "f"
		uservoc[5][1] = "fg"
		uservoc[6][0] = "g"
		uservoc[6][1] = "gh"
		uservoc[7][0] = "h"
		uservoc[7][1] = "hi"
		uservoc[8][0] = "i"
		uservoc[8][1] = "ij"
		uservoc[9][0] = "j"
		uservoc[9][1] = "jk"
		uservoc[10][0] = "k"
		uservoc[10][1] = "kl"
		uservoc[11][0] = "l"
		uservoc[11][1] = "lm"
		uservoc[12][0] = "m"
		uservoc[12][1] = "mn"
		uservoc[13][0] = "n"
		uservoc[13][1] = "no"
		uservoc[14][0] = "o"
		uservoc[14][1] = "op"
		uservoc[15][0] = "p"
		uservoc[15][1] = "pq"
		uservoc[16][0] = "q"
		uservoc[16][1] = "qr"
		uservoc[17][0] = "r"
		uservoc[17][1] = "rs"
		uservoc[18][0] = "s"
		uservoc[18][1] = "st"
		uservoc[19][0] = "t"
		uservoc[19][1] = "tu"
		uservoc[20][0] = "u"
		uservoc[20][1] = "uv"
		uservoc[21][0] = "v"
		uservoc[21][1] = "vw"
		uservoc[22][0] = "w"
		uservoc[22][1] = "wx"
		uservoc[23][0] = "x"
		uservoc[23][1] = "xy"
		uservoc[24][0] = "y"
		uservoc[24][1] = "yz"
		uservoc[25][0] = "z"
		uservoc[25][1] = "za"
	}
}

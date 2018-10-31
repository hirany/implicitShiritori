package main

import (
	"fmt"
	"math/rand"
	"time"
)

type playerData struct {
	iscpu bool
	name  string
	words [][]string
}

// しりとりの内部処理 var.人
// 単語の入力・判定を行い、次に始めるアルファベットを返す。
func humancalc(sw int, seacretalpha, lastalpha byte, newword string) byte {

	// アルファベットの語尾が秘密の文字なら負け
	if newword[len(newword)-1] == seacretalpha {
		fmt.Printf("秘密の文字は %c でした。\n", seacretalpha)
		fmt.Println("ボクの勝ちだ!")
		return '0'
	}

	// 次に始めるアルファベットを返す
	return newword[len(newword)-1]

}

// しりとりの内部処理 var.cpu
// 単語の入力・判定を行い、次に始めるアルファベットを返す。
func cpucalc(sw int, seacretalpha, lastalpha byte) string {

	var newword string

	// 単語の選択
	newword = choiseword(lastalpha, sw)
	if newword == "" {
		fmt.Println("もう思いつかないよ～")
		fmt.Println("ボクの負けだ")
		return "0"
	}

	// アルファベットの語尾が秘密の文字なら負け
	if newword[len(newword)-1] == seacretalpha {
		fmt.Printf("秘密の文字は %c でした。\n", seacretalpha)
		fmt.Println("キミの勝ちだ")
		return "0"
	}

	// cpuが選んだ言葉を返す
	return newword
}

// cpuが言葉を探す
func choiseword(lastalpha byte, sw int) string {

	headalpha := alnum(lastalpha)
	alphalen := wordsLen(playerdata[sw].words[headalpha])
	if alphalen == 0 {
		return ""
	}

	compleetstring := make([]int, alphalen)
	var randnum int
	var flag bool

	rand.Seed(time.Now().UTC().UnixNano())
	randnum = rand.Intn(alphalen)
	compleetstring[randnum] = 1

	for existInUsed(playerdata[sw].words[headalpha][randnum]) {
		flag = true
		rand.Seed(time.Now().UTC().UnixNano())
		randnum = rand.Intn(alphalen)
		if compleetstring[randnum] == 1 {
			continue
		}
		compleetstring[randnum] = 1
		for i := 0; i < alphalen; i++ {
			if compleetstring[i] != 1 {
				flag = false
				break
			}
		}
		if flag {
			return ""
		}
	}

	return playerdata[sw].words[headalpha][randnum]

}

// 使われた単語を追加
// 単語数が上限まで入っている場合falseを返す
func addWord(newword string) bool {

	headalpha := alnum(newword[0])
	alphalen := wordsLen(usedword[headalpha])

	if alphalen > 1000 {
		fmt.Println("flag")
		return false
	}

	usedword[headalpha][alphalen] = newword

	return true

}

// 単語が既に使われているかどうかの判定
// 既に使われている場合trueを返す
func existInUsed(newword string) bool {

	headalpha := alnum(newword[0])
	alphalen := wordsLen(usedword[headalpha])

	for i := 0; i < alphalen; i++ {
		if usedword[headalpha][i] == newword {
			return true
		}
	}

	return false

}

// 単語数を数えて返却
func wordsLen(words []string) int {

	num := 0

	for words[num] != "" {
		num++
	}

	return num
}

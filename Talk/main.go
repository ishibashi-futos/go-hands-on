package main

import (
	"fmt"
)

type MessageTo int

const (
	MESSAGE_TO_KENT    MessageTo = 0
	MESSAGE_TO_YUKPIZ  MessageTo = 1
	MESSAGE_END_ME     MessageTo = 3
	MESSAGE_END_ME_TOO MessageTo = 4
)

func main() {
	ch := make(chan MessageTo) // 相手にボールを渡す為のチャネル
	done := make(chan bool)    // 全ての会話完了を検知するチャネル

	go kent(ch, done)
	go yukpiz(ch, done)

	// 会話が終わるまでブロック
	<-done

	// 閉じているかを確認
	<-ch
	fmt.Println("会話完了")
}

func kent(ch chan MessageTo, done chan bool) {
	texts := []string{"あなたの名前を教えてください", "私はけんとです", "生ハム食べに行きましょう"}
	for v := range ch {
		switch v {
		case MESSAGE_END_ME_TOO:
			// 終了通知が来たらチャネルを閉じる
			close(ch)
			close(done)
		case MESSAGE_END_ME:
			if len(texts) == 0 {
				// セリフがなくなったら完了
				ch <- MESSAGE_END_ME_TOO
				break
			}
			fallthrough
		case MESSAGE_TO_KENT:
			if len(texts) > 0 {
				fmt.Println("kent:" + texts[0]) // 次のセリフを発言する
				texts = texts[1:]               // 発言済みのセリフを消す
				ch <- MESSAGE_TO_YUKPIZ         // 相手にボールを渡す
			} else {
				ch <- MESSAGE_END_ME
			}
		case MESSAGE_TO_YUKPIZ:
			ch <- MESSAGE_TO_YUKPIZ
		}
	}
}

func yukpiz(ch chan MessageTo, done chan bool) {
	texts := []string{"私はゆくぴずです", "よろしくおねがいします", "行きましょう"}
	ch <- MESSAGE_TO_KENT
	for v := range ch {
		switch v {
		case MESSAGE_END_ME_TOO:
			close(ch)
			close(done)
		case MESSAGE_END_ME:
			if len(texts) == 0 {
				ch <- MESSAGE_END_ME_TOO
				break
			}
			fallthrough
		case MESSAGE_TO_KENT:
			ch <- MESSAGE_TO_KENT
		case MESSAGE_TO_YUKPIZ:
			if len(texts) > 0 {
				fmt.Println("yukpiz:" + texts[0])
				texts = texts[1:]
				ch <- MESSAGE_TO_KENT
			} else {
				ch <- MESSAGE_END_ME
			}
		}
	}
}

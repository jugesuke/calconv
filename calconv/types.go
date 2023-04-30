package calconv

import "time"

type Event struct {
	Subject     string    //予定の名前
	StartTime   time.Time //予定の開始日時
	EndTime     time.Time //予定の終了日時
	AllDayEvent bool      //終日の予定であるかどうかを指定します 終日の予定の場合は True,そうでない場合は False と入力します.
	Description string    //予定の説明やメモ
	Location    string    //予定の場所
}

type ClassMetaData struct {
	Date    Date
	Hour    int
	Subject string
	Room    string
	LMSLink string
}

type ClassDay struct {
	Day  time.Time
	Date Date
}

type Date int

const (
	Sun Date = iota
	Mon
	Tue
	Wed
	Thu
	Fri
	Sat
)

package model

// RHeader 收到消息头部信息
type RHeader struct {
	Length 		uint32
	LHead    	uint32
	Type   		uint32
	Normal   	uint32
}

// Population 人气值
type Population struct {
	Value 		uint32
}

// CMD cmd消息
type CMD struct {
	Info [][]interface{} `json:"info"`
	Cmd  string          `json:"cmd"`
}

// Danmaku 弹幕
type Danmaku struct {
	Info []string `json:"info"`
}

// Gift 礼物
type Gift struct {
	Data struct {
		GiftName  string `json:"giftName"`
		Num       uint   `json:"num"`
		Uname     string `json:"uname"`
		UID       uint   `json:"uid"`
		Remain    uint   `json:"remain"`
		CoinType  string `json:"coin_type"`
		TotalCoin uint   `json:"total_coin"`
	} `json:"data"`
}

// Guard 舰队
type Guard struct {
	Data struct {
		UID        uint   `json:"uid"`
		Username   string `json:"username"`
		GuardLevel int    `json:"guard_level"`
		Num        uint   `json:"num"`
		Price      uint   `json:"price"`
		GiftID     int    `json:"gift_id"`
		GiftName   string `json:"gift_name"`
		StartTime  int    `json:"start_time"`
		EndTime    int    `json:"end_time"`
	} `json:"data"`
}
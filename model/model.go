package model

// RHeader 收到消息头部信息
type RHeader struct {
	Length uint32
	LHead  uint32
	Type   uint32
	Normal uint32
}

// HeartBeat 人气值
type HeartBeat struct {
	Value uint32
}

// CMD cmd消息
type CMD struct {
	Info [][]interface{} `json:"info"`
	Cmd  string          `json:"cmd"`
}

// Infomation 弹幕
type Infomation struct {
	Cmd  string        `json:"cmd"`
	Info []interface{} `json:"info"`
}

// GiftCmd cmd
type GiftCmd struct {
	Cmd  string `json:"cmd"`
	Data Gift   `json:"data"`
}

// Gift 礼物
type Gift struct {
	GiftName          string        `json:"giftName"`
	Num               int           `json:"num"`
	Uname             string        `json:"uname"`
	Face              string        `json:"face"`
	GuardLevel        int           `json:"guard_level"`
	Rcost             int           `json:"rcost"`
	UID               int           `json:"uid"`
	TopList           []interface{} `json:"top_list"`
	Timestamp         int           `json:"timestamp"`
	GiftID            int           `json:"giftId"`
	GiftType          int           `json:"giftType"`
	Action            string        `json:"action"`
	Super             int           `json:"super"`
	SuperGiftNum      int           `json:"super_gift_num"`
	SuperBatchGiftNum int           `json:"super_batch_gift_num"`
	BatchComboID      string        `json:"batch_combo_id"`
	Price             int           `json:"price"`
	Rnd               string        `json:"rnd"`
	NewMedal          int           `json:"newMedal"`
	NewTitle          int           `json:"newTitle"`
	Medal             []interface{} `json:"medal"`
	Title             string        `json:"title"`
	BeatID            string        `json:"beatId"`
	BizSource         string        `json:"biz_source"`
	Metadata          string        `json:"metadata"`
	Remain            int           `json:"remain"`
	Gold              int           `json:"gold"`
	Silver            int           `json:"silver"`
	EventScore        int           `json:"eventScore"`
	EventNum          int           `json:"eventNum"`
	SmalltvMsg        []interface{} `json:"smalltv_msg"`
	SpecialGift       interface{}   `json:"specialGift"`
	NoticeMsg         []interface{} `json:"notice_msg"`
	SmallTVCountFlag  bool          `json:"smallTVCountFlag"`
	Capsule           interface{}   `json:"capsule"`
	AddFollow         int           `json:"addFollow"`
	EffectBlock       int           `json:"effect_block"`
	CoinType          string        `json:"coin_type"`
	TotalCoin         int           `json:"total_coin"`
	Effect            int           `json:"effect"`
	BroadcastID       int           `json:"broadcast_id"`
	Draw              int           `json:"draw"`
	CritProb          int           `json:"crit_prob"`
	TagImage          string        `json:"tag_image"`
	SendMaster        interface{}   `json:"send_master"`
	IsFirst           bool          `json:"is_first"`
	Demarcation       int           `json:"demarcation"`
	ComboStayTime     int           `json:"combo_stay_time"`
	ComboTotalCoin    int           `json:"combo_total_coin"`
}

// GuardCmd cmd
type GuardCmd struct {
	Cmd  string `json:"cmd"`
	Data Guard  `json:"data"`
}

// Guard 舰队
type Guard struct {
	UID        uint   `json:"uid"`
	Username   string `json:"username"`
	GuardLevel int    `json:"guard_level"`
	Num        uint   `json:"num"`
	Price      uint   `json:"price"`
	GiftID     int    `json:"gift_id"`
	GiftName   string `json:"gift_name"`
	StartTime  int    `json:"start_time"`
	EndTime    int    `json:"end_time"`
}

package message

const (
	DANMU_MSG         = "DANMU_MSG"
	SEND_GIFT         = "SEND_GIFT"
	ROOM_RANK         = "ROOM_RANK"
	INTERACT_WORD     = "INTERACT_WORD"
	ONLINE_RANK_V2    = "ONLINE_RANK_V2"
	ONLINE_RANK_TOP3  = "ONLINE_RANK_TOP3"
	ONLINE_RANK_COUNT = "ONLINE_RANK_COUNT"
	COMBO_SEND        = "COMBO_SEND"
	WIDGET_BANNER     = "WIDGET_BANNER"
	ENTRY_EFFECT      = "ENTRY_EFFECT"
)

type Danmaku struct {
	Content  string
	Username string
	UID      uint
}

type RoomRank struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Roomid    int    `json:"roomid"`
		RankDesc  string `json:"rank_desc"`
		Color     string `json:"color"`
		H5URL     string `json:"h5_url"`
		WebURL    string `json:"web_url"`
		Timestamp int    `json:"timestamp"`
	} `json:"data"`
}

type SendGift struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Draw              int         `json:"draw"`
		Gold              int         `json:"gold"`
		Silver            int         `json:"silver"`
		Num               int         `json:"num"`
		TotalCoin         int         `json:"total_coin"`
		Effect            int         `json:"effect"`
		BroadcastID       int         `json:"broadcast_id"`
		CritProb          int         `json:"crit_prob"`
		GuardLevel        int         `json:"guard_level"`
		Rcost             int         `json:"rcost"`
		UID               int         `json:"uid"`
		Timestamp         int         `json:"timestamp"`
		GiftID            int         `json:"giftId"`
		GiftType          int         `json:"giftType"`
		Super             int         `json:"super"`
		SuperGiftNum      int         `json:"super_gift_num"`
		SuperBatchGiftNum int         `json:"super_batch_gift_num"`
		Remain            int         `json:"remain"`
		Price             int         `json:"price"`
		BeatID            string      `json:"beatId"`
		BizSource         string      `json:"biz_source"`
		Action            string      `json:"action"`
		CoinType          string      `json:"coin_type"`
		Uname             string      `json:"uname"`
		Face              string      `json:"face"`
		BatchComboID      string      `json:"batch_combo_id"`
		Rnd               string      `json:"rnd"`
		GiftName          string      `json:"giftName"`
		ComboSend         interface{} `json:"combo_send"`
		BatchComboSend    interface{} `json:"batch_combo_send"`
		TagImage          string      `json:"tag_image"`
		TopList           interface{} `json:"top_list"`
		SendMaster        interface{} `json:"send_master"`
		IsFirst           bool        `json:"is_first"`
		Demarcation       int         `json:"demarcation"`
		ComboStayTime     int         `json:"combo_stay_time"`
		ComboTotalCoin    int         `json:"combo_total_coin"`
		Tid               string      `json:"tid"`
		EffectBlock       int         `json:"effect_block"`
		IsSpecialBatch    int         `json:"is_special_batch"`
		ComboResourcesID  int         `json:"combo_resources_id"`
		Magnification     float64     `json:"magnification"`
		NameColor         string      `json:"name_color"`
		MedalInfo         struct {
			TargetID         int    `json:"target_id"`
			Special          string `json:"special"`
			IconID           int    `json:"icon_id"`
			AnchorUname      string `json:"anchor_uname"`
			AnchorRoomid     int    `json:"anchor_roomid"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int    `json:"medal_color"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			IsLighted        int    `json:"is_lighted"`
			GuardLevel       int    `json:"guard_level"`
		} `json:"medal_info"`
		SvgaBlock int `json:"svga_block"`
	} `json:"data"`
}

type InteractWord struct {
	Cmd string `json:"cmd"`

	Data struct {
		UID        int    `json:"uid"`
		Uname      string `json:"uname"`
		UnameColor string `json:"uname_color"`
		Identities []int  `json:"identities"`
		MsgType    int    `json:"msg_type"`
		Roomid     int    `json:"roomid"`
		Timestamp  int    `json:"timestamp"`
		Score      int64  `json:"score"`
		FansMedal  struct {
			TargetID         int    `json:"target_id"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int    `json:"medal_color"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			IsLighted        int    `json:"is_lighted"`
			GuardLevel       int    `json:"guard_level"`
			Special          string `json:"special"`
			IconID           int    `json:"icon_id"`
			AnchorRoomid     int    `json:"anchor_roomid"`
			Score            int    `json:"score"`
		} `json:"fans_medal"`
		IsSpread     int    `json:"is_spread"`
		SpreadInfo   string `json:"spread_info"`
		Contribution struct {
			Grade int `json:"grade"`
		} `json:"contribution"`
		SpreadDesc string `json:"spread_desc"`
		TailIcon   int    `json:"tail_icon"`
	} `json:"data"`
}

type OnlineRankV2 struct {
	Cmd  string `json:"cmd"`
	Data struct {
		List []struct {
			UID        int    `json:"uid"`
			Face       string `json:"face"`
			Score      string `json:"score"`
			Uname      string `json:"uname"`
			Rank       int    `json:"rank"`
			GuardLevel int    `json:"guard_level"`
		} `json:"list"`
		RankType string `json:"rank_type"`
	} `json:"data"`
}

type OnlineRankTOP3 struct {
	Cmd  string `json:"cmd"`
	Data struct {
		List []struct {
			Msg  string `json:"msg"`
			Rank int    `json:"rank"`
		} `json:"list"`
	} `json:"data"`
}

type ComboSend struct {
	Cmd  string `json:"cmd"`
	Data struct {
		UID           int         `json:"uid"`
		Ruid          int         `json:"ruid"`
		Uname         string      `json:"uname"`
		RUname        string      `json:"r_uname"`
		ComboNum      int         `json:"combo_num"`
		GiftID        int         `json:"gift_id"`
		GiftNum       int         `json:"gift_num"`
		BatchComboNum int         `json:"batch_combo_num"`
		GiftName      string      `json:"gift_name"`
		Action        string      `json:"action"`
		ComboID       string      `json:"combo_id"`
		BatchComboID  string      `json:"batch_combo_id"`
		IsShow        int         `json:"is_show"`
		SendMaster    interface{} `json:"send_master"`
		NameColor     string      `json:"name_color"`
		TotalNum      int         `json:"total_num"`
		MedalInfo     struct {
			TargetID         int    `json:"target_id"`
			Special          string `json:"special"`
			IconID           int    `json:"icon_id"`
			AnchorUname      string `json:"anchor_uname"`
			AnchorRoomid     int    `json:"anchor_roomid"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int    `json:"medal_color"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			IsLighted        int    `json:"is_lighted"`
			GuardLevel       int    `json:"guard_level"`
		} `json:"medal_info"`
		ComboTotalCoin int `json:"combo_total_coin"`
	} `json:"data"`
}

type WidgetBanner struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Timestamp  int `json:"timestamp"`
		WidgetList struct {
			Num7 struct {
				Type    int    `json:"type"`
				BandID  int    `json:"band_id"`
				SubKey  string `json:"sub_key"`
				SubData string `json:"sub_data"`
			} `json:"7"`
		} `json:"widget_list"`
	} `json:"data"`
}

type EntryEffect struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ID               int           `json:"id"`
		UID              int           `json:"uid"`
		TargetID         int           `json:"target_id"`
		MockEffect       int           `json:"mock_effect"`
		Face             string        `json:"face"`
		PrivilegeType    int           `json:"privilege_type"`
		CopyWriting      string        `json:"copy_writing"`
		CopyColor        string        `json:"copy_color"`
		HighlightColor   string        `json:"highlight_color"`
		Priority         int           `json:"priority"`
		BasemapURL       string        `json:"basemap_url"`
		ShowAvatar       int           `json:"show_avatar"`
		EffectiveTime    int           `json:"effective_time"`
		WebBasemapURL    string        `json:"web_basemap_url"`
		WebEffectiveTime int           `json:"web_effective_time"`
		WebEffectClose   int           `json:"web_effect_close"`
		WebCloseTime     int           `json:"web_close_time"`
		Business         int           `json:"business"`
		CopyWritingV2    string        `json:"copy_writing_v2"`
		IconList         []interface{} `json:"icon_list"`
		MaxDelayTime     int           `json:"max_delay_time"`
	} `json:"data"`
}

type OnlineRankCount struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Count int `json:"count"`
	} `json:"data"`
}

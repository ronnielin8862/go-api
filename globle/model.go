package globle

type (
	ChatHistory struct {
		Id         int64  `json:"id" gorm:"primaryKey;autoIncrement"` // bigint(20) unsigned NOT NULL AUTO_INCREMENT,
		Uid        int    `json:"uid"`                                //int(11) DEFAULT NULL COMMENT '用户id',
		RoomId     int    `json:"room_id"`                            //varchar(100) DEFAULT NULL COMMENT '房间id',
		CreateTime int64  `json:"create_time"`                        //datetime DEFAULT NULL,
		Content    string `json:"content"`
	}

	SendGiftReq struct {
		Anchorid int   `json:"anchorid" validate:"required"`
		Giftid   int   `json:"giftid" validate:"required"`
		Liveid   int64 `json:"liveid"`
		Count    int   `json:"count"`
		Uid      int
	}

	RoomScoreLiveMessage struct {
		SentMessageStruct
		ScoreLive RoomScoreLive `json:"score_live"`
	}

	RoomScoreLive struct {
		Id          int   `json:"match_id"`
		Status      int   `json:"match_status"`
		HomeScore   Score `json:"home_score"`
		AwayScore   Score `json:"away_score"`
		KickOutTime int64 `json:"kick_out_time"`
	}

	Score struct {
		Score        int `json:"score"`
		HalfScore    int `json:"half_score"`
		RedCard      int `json:"red_card"`
		YellowCard   int `json:"yellow_card"`
		CornerKick   int `json:"corner_kick"`
		OTScore      int `json:"ot_score"`
		PenaltyScore int `json:"penalty_score"`
	}

	FootballStatsLiveMessage struct {
		Id   int64 `json:"match_id"`
		Type int   `json:"type"`
		Home int   `json:"home"`
		Away int   `json:"away"`
	}

	BasketStatsLiveMessage struct {
		Id   int64   `json:"match_id"`
		Type int     `json:"type"`
		Home float64 `json:"home"`
		Away float64 `json:"away"`
	}

	RoomTextLiveMessage struct {
		SentMessageStruct
		Id         int64  `json:"match_id"`    // 赛事id
		Time       string `json:"time"`        // 事件时间
		Type       int8   `json:"type"`        // 事件类型
		Data       string `json:"data"`        // 事件文本
		Position   int8   `json:"position"`    // 事件發生方， 0-中立 1-主队 2-客队
		Main       int8   `json:"main"`        // 是否重要事件 0-否 1-是
		CreateTime int64  `json:"create_time"` // 創建時間
	}

	TextLiveStruct struct {
		Id         int    `json:"match_id"`    // 赛事id
		Time       string `json:"time"`        // 事件时间
		Type       int8   `json:"type"`        // 事件类型
		Data       string `json:"data"`        // 事件文本
		Position   int8   `json:"position"`    // 事件發生方， 0-中立 1-主队 2-客队
		Main       int8   `json:"main"`        // 是否重要事件 0-否 1-是
		AwayScore  int    `json:"away_score"`  // 客队比分
		HomeScore  int    `json:"home_score"`  // 主队比分
		CreateTime int64  `json:"create_time"` // 創建時間
	}
	SentMessageStruct struct {
		Type    string `json:"type"`
		Message string `default:"" json:"message,omitempty"`
	}

	Lineup struct {
		MatchId       int          `json:"match_id"`       // 比賽ID
		Confirmed     int8         `json:"confirmed"`      // 正式阵容，1-是、0-不是
		HomeFormation string       `json:"home_formation"` // 主队阵型 3-4-3
		AwayFormation string       `json:"away_formation"` // 客队阵型 4-4-1-1
		Home          []LineupItem `json:"home"`           // 主隊球員信息
		Away          []LineupItem `json:"away"`           // 客隊球員信息
	}
	LineupItem struct {
		LineupId     int         `json:"id"`            // 球员id
		TeamId       int         `json:"team_id"`       // 球队id
		First        int8        `json:"first"`         // 是否首发，1-是、0-否
		Captain      int8        `json:"captain"`       // 是否队长，1-是、0-否
		Name         string      `json:"name"`          // 球员名称
		Logo         string      `json:"logo"`          // 球员logo
		NationalLogo string      `json:"national_logo"` // 球员logo(国家队)
		ShirtNumber  int         `json:"shirt_number"`  // 球衣号
		Position     string      `json:"position"`      // 球员位置，F前锋、M中场、D后卫、G守门员、其他为未知
		X            int8        `json:"x"`             // 阵容x坐标，总共100
		Y            int8        `json:"y"`             // 阵容y坐标，总共100
		Rating       string      `json:"rating"`        // 评分，10为满分
		Incidents    []Incidents `json:"Incidents"`     // 球员事件列表，有事件存在，默认不存在
	}
	Incidents struct {
		Type       int    `json:"type"`        // 事件类型，详见状态码->技术统计
		Time       string `json:"time"`        // 事件发生时间（含加时时间，'A+B':A-比赛时间,B-加时时间）
		Belong     int8   `json:"belong"`      // 发生方，0-中立、1-主队、2-客队
		HomeScore  int    `json:"home_score"`  // 主队比分
		AwayScore  int    `json:"away_score"`  // 客队比分
		ReasonType int8   `json:"reason_type"` // 红黄牌、换人事件原因，详见状态码->事件原因（红黄牌、换人事件存在）
		Player     Player `json:"player"`      // player-相关球员
		Assist1    Player `json:"assist1"`     // assist1-助攻球员1
		Assist2    Player `json:"assist2"`     // assist1-助攻球员2
		InPlayer   Player `json:"in_player"`   // in_player-换上球员
		OutPlayer  Player `json:"out_player"`  // out_player-换下球员
	}
	Player struct {
		PlayerId int    `json:"id"`   // 球员id
		Name     string `json:"name"` // 中文名称
	}
	BasketballScore struct {
		MatchId     int   `json:"match_id"`     // 比赛id
		MatchStatus int   `json:"match_status"` // 比赛状态，详见状态码->比赛状态
		TimeLeft    int   `json:"time_left"`    // 小节剩余时间(秒)
		AwayScore   []int `json:"away_score"`   // 客队比分
		HomeScore   []int `json:"home_score"`   // 主队比分 [23, 32, 22, 34, 0 ], 数组下标 （0 第一节比分 1 第二节比分，2 第三节比分，3 //第4节分数 4 //加时分数）
	}

	BasketballRecord struct {
		MatchId      int          `json:"match_id"`       // 比赛ID
		Home         []PlayerInfo `json:"home"`           // 主队阵容
		Away         []PlayerInfo `json:"away"`           // 客队真人
		HomeTimeLine TimeLineItem `json:"home_time_line"` // 主队在场持续时间统计
		AwayTimeLine TimeLineItem `json:"away_time_line"` // 客队在场持续时间统计
	}
	PlayerInfo struct {
		Id       int64        `json:"id"`       // 球员id
		NameZh   string       `json:"name_zh"`  // 球员中文名称
		NameZht  string       `json:"name_zht"` // 球员粤语名称
		NameEn   string       `json:"name_en"`  // 球员英文名称
		Logo     string       `json:"logo"`     // 球员logo
		Number   string       `json:"number"`   // 球衣号
		TimeLine TimeLineItem `json:"timeline"` // 在场持续时间统计
	}
	TimeLineItem struct {
		HitCount          int `json:"hit_count"`          // 命中次数
		ShotCount         int `json:"shot_count"`         // 投篮次数
		Goal3Score        int `json:"goal_3_score"`       // 三分球投篮命中次数
		Goal3ScoreCount   int `json:"goal_3_score_count"` // 三分投篮次数
		FQHitCount        int `json:"fq_hit_count"`       // 罚球命中次数
		FQCount           int `json:"fq_count"`           // 罚球投篮次数
		OffensiveRebounds int `json:"offensive_rebounds"` // 进攻篮板
		DefensiveRebounds int `json:"defensive_rebounds"` // 防守篮板
		TotalRebounds     int `json:"total_rebounds"`     // 总的篮板
		Assists           int `json:"assists"`            // 助攻数
		Steals            int `json:"steals"`             // 抢断数
		Caps              int `json:"caps"`               // 盖帽数
		MistakeCount      int `json:"mistake_count"`      // 失误次数
		Fouls             int `json:"fouls"`              // 个人犯规次数
	}

	LiveStream struct {
		RoomId   int    `json:"id"`       // 房间id
		Function string `json:"function"` // 功能类型，直播、点播
		Uid      int    `json:"uid"`
		Content  string `json:"message"` // 消息
	}

	QuizContent struct {
		ItemNum       int   `json:"ItemNum"`
		Status        int   `json:"Status"`
		PeopleNumList []int `json:"PeopleNumList"`
		AmountList    []int `json:"AmountList"`
		MinBet        int   `json:"MinBet"`
	}

	IndexNats struct {
		Index
		ExponentialScore string `json:"exponential_score"` // 进球比分 or 角球比(cr)，主队-客队
		MatchId          int    `json:"match_id"`          // 比赛id
		MatchStatus      int    `json:"match_status"`      // 比赛状态，详见状态码->比赛状态 - int
		OddsStatus       int    `json:"odds_status"`       // 是否封盘：1-封盘,0-未封盘
		CompetitionTime  string `json:"competition_time"`  // 比赛进行时间
	}

	Index struct {
		CompanyId       int    `json:"company_id"`            // 指数公司id，详见状态码->指数公司ID
		ExponentialType string `json:"exponential_type"`      // 指数类型：asia-亚盘、eu-欧赔、bs-大小球、cr-角球
		Odds1           string `json:"odds_1"`                // 主胜/大球/大
		Odds2           string `json:"odds_2"`                // 和局/盘口
		Odds3           string `json:"odds_3"`                // 客胜/小球/小
		ChangeTime      int64  `json:"change_time,omitempty"` // 变化时间
	}

	GetTextLiveReq struct {
		RoomId int `json:"id"`
		LastId int `json:"last_id"`
		Size   int `json:"size" default:"20"`
	}
)

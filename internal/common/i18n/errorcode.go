package i18n

type errorCode struct {
	En string
	Tw string
	Cn string
}

/*
	錯誤代碼設計分為三個區段
	由左至右每4個代碼代表一個區段
*/
var errorCodeMap = map[string]errorCode{ // TODO 需要修改語系字串
	/*
		通用的類別
		00010001[0001-9999]
	*/
	"000100010001": errorCode{En: "錯誤代碼尚未定義", Tw: "錯誤代碼尚未定義", Cn: "錯誤代碼尚未定義"},
	"000100010002": errorCode{En: "JSON格式錯誤", Tw: "JSON格式錯誤", Cn: "JSON格式錯誤"},
	"000100010003": errorCode{En: "API權限不足", Tw: "API權限不足", Cn: "API權限不足"},
	"000100010004": errorCode{En: "缺少傳入參數", Tw: "缺少傳入參數", Cn: "缺少傳入參數"},
	"000100010005": errorCode{En: "無效的 username", Tw: "無效的 username", Cn: "無效的 username"},
	"000100010006": errorCode{En: "無效的 password", Tw: "無效的 password", Cn: "無效的 password"},
	"000100010007": errorCode{En: "無效的 alias", Tw: "無效的 alias", Cn: "無效的 alias"},
	"000100010008": errorCode{En: "無效的使用者階層", Tw: "無效的使用者階層", Cn: "無效的使用者階層"},
	"000100010009": errorCode{En: "無法建立 db連線", Tw: "無法建立 db連線", Cn: "無法建立 db連線"},
	"000100010010": errorCode{En: "無法初始化資料表", Tw: "無法建立初始化資料表", Cn: "無法建立初始化資料表"},
	"000100010011": errorCode{En: "資料解析錯誤", Tw: "資料解析錯誤", Cn: "資料解析錯誤"},
	"000100010012": errorCode{En: "資料加密失敗", Tw: "資料加密失敗", Cn: "資料加密失敗"},
	"000100010013": errorCode{En: "資料無法寫入 db", Tw: "資料無法寫入 db", Cn: "資料無法寫入 db"},
	"000100010014": errorCode{En: "db連線取得失敗", Tw: "db連線取得失敗", Cn: "db連線取得失敗"},

	/*
		身份驗證類別
		00010002[0001-9999]
	*/
	"000100020001": errorCode{En: "帳號註冊成功", Tw: "帳號註冊成功", Cn: "帳號註冊成功"},
	"000100020002": errorCode{En: "帳號名稱已經存在", Tw: "帳號名稱已經存在", Cn: "帳號名稱已經存在"},
	"000100020003": errorCode{En: "帳號註冊失敗", Tw: "帳號註冊失敗", Cn: "帳號註冊失敗"},
	"000100020004": errorCode{En: "密碼錯誤", Tw: "密碼錯誤", Cn: "密碼錯誤"},
}

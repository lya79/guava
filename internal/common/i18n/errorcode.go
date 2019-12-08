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
var errorCodeMap = map[string]errorCode{
	/*
		通用的類別
		00010001[0001-9999]
	*/
	"000100010001": errorCode{En: "錯誤代碼尚未定義", Tw: "錯誤代碼尚未定義", Cn: "錯誤代碼尚未定義"},
	"000100010002": errorCode{En: "JSON格式錯誤", Tw: "JSON格式錯誤", Cn: "JSON格式錯誤"},
	"000100010003": errorCode{En: "API權限不足", Tw: "API權限不足", Cn: "API權限不足"},

	/*
		身份驗證類別
		00010002[0001-9999]
	*/
	"000100020001": errorCode{En: "帳號註冊成功", Tw: "帳號註冊成功", Cn: "帳號註冊成功"},
	"000100020002": errorCode{En: "帳號名稱已經存在", Tw: "帳號名稱已經存在", Cn: "帳號名稱已經存在"},
}

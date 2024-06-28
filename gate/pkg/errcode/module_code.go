package errcode

var (
	ErrorOutOfRange = BadRequestErrorData(20000001, "超出分頁範圍")

	ErrorUploadFileFail = NewErrorData(20010001, "上載文件失敗")

	ErrorUserNotExist        = BadRequestErrorData(20020001, "用戶不存在")
	ErrorConflictUser        = BadRequestErrorData(20020002, "已存在此用戶")
	ErrorUserDomainFail      = BadRequestErrorData(20020003, "Email網域不合法")
	ErrorUserGoogleTokenFail = BadRequestErrorData(20020004, "Google 驗證失敗")

	ErrorGetProjectListFail = BadRequestErrorData(20030001, "取得專案列表失敗")
	ErrorCreateProjectFail  = BadRequestErrorData(20030002, "創建專案失敗")
	ErrorUpdateProjectFail  = BadRequestErrorData(20030003, "更新專案失敗")
	ErrorDeleteProjectFail  = BadRequestErrorData(20030004, "刪除專案失敗")
	ErrorProjectNotExist    = BadRequestErrorData(20030005, "專案不存在")

	ErrorGetPublishVersionListFail = BadRequestErrorData(20040001, "取得版本列表失敗")
	ErrorCreatePublishVersionFail  = BadRequestErrorData(20040002, "創建版本失敗")
	ErrorUpdatePublishVersionFail  = BadRequestErrorData(20040003, "更新版本失敗")
	ErrorDeletePublishVersionFail  = BadRequestErrorData(20040004, "刪除版本失敗")
	ErrorPublishVersionNotExist    = BadRequestErrorData(20040005, "版本不存在")
	ErrorConflictPublishVersion    = BadRequestErrorData(20040006, "版本重複")

	ErrorGetKeyListFail = BadRequestErrorData(20050001, "取得Key列表失敗")
	ErrorCreateKeyFail  = BadRequestErrorData(20050002, "創建Key失敗")
	ErrorUpdateKeyFail  = BadRequestErrorData(20050003, "更新Key失敗")
	ErrorDeleteKeyFail  = BadRequestErrorData(20050004, "刪除Key失敗")
	ErrorKeyNotExist    = BadRequestErrorData(20050005, "Key不存在")
	ErrorConflictKey    = BadRequestErrorData(20050006, "Key重複")

	ErrorExportNoDataAvailable = BadRequestErrorData(20060001, "No data available")
)

package init

import (
	"lvbu/models/env"
	men "lvbu/models/env"
	mac "lvbu/models/machine"
	mper "lvbu/models/permission"
	mpro "lvbu/models/project"
	"lvbu/models/user"
	"lvbu/utils"

	"github.com/astaxie/beego"
)

func InitTemplate() {
	beego.AddFuncMap("Getuser", user.Getuser)
	beego.AddFuncMap("Getposname", user.Getposname)
	beego.AddFuncMap("Getpospers", mper.Getpospers)
	beego.AddFuncMap("Isperitem", mper.Isperitem)
	beego.AddFuncMap("Isuserper", mper.Isuserper)

	beego.AddFuncMap("Getposition", user.Getposition)
	beego.AddFuncMap("Getenv", env.Getenv)
	beego.AddFuncMap("Getenvidlist", men.GetEnvIdList)

	beego.AddFuncMap("Getmacforenv", mac.GetMacforenv)

	beego.AddFuncMap("Getproject", mpro.Getproject)
	beego.AddFuncMap("Getprofornodecount", mpro.Getprofornodecount)

	beego.AddFuncMap("Getnode", mpro.Getnode)
	//	beego.AddFuncMap("getProjectname", projects.GetProjectName)
	//	beego.AddFuncMap("getLeaveProcess", leaves.ListLeaveApproverProcessHtml)
	//	beego.AddFuncMap("getExpenseProcess", expenses.ListExpenseApproverProcessHtml)
	//	beego.AddFuncMap("getBusinesstripProcess", businesstrips.ListBusinesstripApproverProcessHtml)
	//	beego.AddFuncMap("getGooutProcess", goouts.ListGooutApproverProcessHtml)
	//	beego.AddFuncMap("getOagoodProcess", oagoods.ListOagoodApproverProcessHtml)
	//	beego.AddFuncMap("getOvertimeProcess", overtimes.ListOvertimeApproverProcessHtml)

	beego.AddFuncMap("getDate", utils.GetDate)
	beego.AddFuncMap("getDateMH", utils.GetDateMH)
	//beego.AddFuncMap("getNeedsStatus", utils.GetNeedsStatus)
	beego.AddFuncMap("getNeedsSource", utils.GetNeedsSource)
	beego.AddFuncMap("getNeedsStage", utils.GetNeedsStage)
	beego.AddFuncMap("getTaskStatus", utils.GetTaskStatus)
	beego.AddFuncMap("getTaskType", utils.GetTaskType)
	beego.AddFuncMap("getTestStatus", utils.GetTestStatus)
	beego.AddFuncMap("getLeaveType", utils.GetLeaveType)

	beego.AddFuncMap("getOs", utils.GetOs)
	beego.AddFuncMap("getBrowser", utils.GetBrowser)
	beego.AddFuncMap("getAvatarSource", utils.GetAvatarSource)
	beego.AddFuncMap("getAvatar", utils.GetAvatar)
	//beego.AddFuncMap("getAvatarUserid", users.GetAvatarUserid)
	//	beego.AddFuncMap("getPositionsName", users.GetPositionsNameForUserid)
	//	beego.AddFuncMap("getDepartmentsName", users.GetDepartmentsNameForUserid)
	beego.AddFuncMap("Getsex", utils.Getsex)
	beego.AddFuncMap("getEdu", utils.GetEdu)
	beego.AddFuncMap("getWorkYear", utils.GetWorkYear)
	beego.AddFuncMap("getResumeStatus", utils.GetResumeStatus)

	beego.AddFuncMap("getCheckworkType", utils.GetCheckworkType)

	beego.AddFuncMap("getMessageType", utils.GetMessageType)
	beego.AddFuncMap("getMessageSubtype", utils.GetMessageSubtype)

}

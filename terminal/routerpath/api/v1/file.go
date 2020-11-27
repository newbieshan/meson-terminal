package v1

import (
	"github.com/daqnext/meson-common/common"
	"github.com/daqnext/meson-common/common/commonmsg"
	"github.com/daqnext/meson-common/common/downloadtaskmgr"
	"github.com/daqnext/meson-common/common/resp"
	"github.com/daqnext/meson-terminal/terminal/manager/filemgr"
	"github.com/gin-gonic/gin"
)

func init() {
	common.AutoConfigRouter()

	// /api/v1/file/save
	common.GetMyRouter().POST("/save", saveNewFileHandler)
	// /api/v1/file/delete
	common.GetMyRouter().POST("/delete", deleteFileHandler)

	// /api/v1/file/deletefolder
	common.GetMyRouter().POST("/deletefolder", deleteFolderHandler)
}

func saveNewFileHandler(ctx *gin.Context) {
	//get cmd msg
	var downloadCmd commonmsg.DownLoadFileCmdMsg
	if err := ctx.ShouldBindJSON(&downloadCmd); err != nil {
		resp.ErrorResp(ctx, resp.ErrMalParams)
		return
	}

	//check file exist or not
	//filePath := global.FileDirPath + "/" + downloadCmd.BindNameHash + "/" + downloadCmd.FileNameHash
	//if utils.Exists(filePath) {
	//	//report success
	//	resp.SuccessResp(ctx, nil)
	//
	//	//post finish msg to server
	//	payload := commonmsg.TerminalDownloadFinishMsg{
	//		TransferTag:  downloadCmd.TransferTag,
	//		FileNameHash: downloadCmd.FileNameHash,
	//		BindNameHash: downloadCmd.BindNameHash,
	//		Continent:    downloadCmd.Continent,
	//		Country:      downloadCmd.Country,
	//		Area:         downloadCmd.Area,
	//	}
	//	header := map[string]string{
	//		"Content-Type":  "application/json",
	//		"Authorization": "Bearer " + accountmgr.Token,
	//	}
	//	_, err := httputils.Request("POST", global.ReportDownloadFinishUrl, payload, header)
	//	if err != nil {
	//		logger.Error("send downloadfinish msg to server error", "err", err)
	//	}
	//
	//	return
	//}

	//if not exist, start download
	err := downloadtaskmgr.AddTask(
		downloadCmd.DownloadUrl,
		downloadCmd.TransferTag,
		downloadCmd.Continent,
		downloadCmd.Country,
		downloadCmd.Area,
		downloadCmd.BindNameHash,
		downloadCmd.FileNameHash,
		0,
	)
	if err != nil {
		resp.ErrorResp(ctx, resp.ErrAddDownloadTaskFailed)
		return
	}
	resp.SuccessResp(ctx, nil)
}

func deleteFileHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": 0,
		"msg":    "deleteFileHandler",
	})
}

func deleteFolderHandler(ctx *gin.Context) {
	var msg commonmsg.DeleteFolderCmdMsg
	if err := ctx.ShouldBindJSON(&msg); err != nil {
		resp.ErrorResp(ctx, resp.ErrMalParams)
		return
	}

	err := filemgr.DeleteFolder(msg.FolderName)
	if err != nil {
		resp.ErrorResp(ctx, resp.ErrInternalError)
		return
	}

	resp.SuccessResp(ctx, nil)
}

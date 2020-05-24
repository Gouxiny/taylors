import service from '@/utils/request'
// @Tags api
// @Summary 分页获取TOP列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取TOP列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
// {
//  page     int
//	pageSize int
// }
export const getTopList = (data) => {
    return service({
        url: "/api/top/list",
        method: 'post',
        data
    })
}

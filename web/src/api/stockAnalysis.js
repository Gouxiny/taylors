import service from '@/utils/request'

//获取Analysis列表
export const getAnalysisList = (data) => {
    return service({
        url: "/stock/analysis/list",
        method: 'post',
        data
    })
}
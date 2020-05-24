import service from '@/utils/request'

//获取TOP列表
export const getTopList = (data) => {
    return service({
        url: "/stock/top/list",
        method: 'post',
        data
    })
}
import service from '@/utils/request'

//获取All列表
export const getAllList = (data) => {
    return service({
        url: "/stock/all/list",
        method: 'post',
        data
    })
}
import service from '@/utils/request'

//获取Monitor列表
export const getMonitorList = (data) => {
    return service({
        url: "/stock/monitor/list",
        method: 'post',
        data
    })
}
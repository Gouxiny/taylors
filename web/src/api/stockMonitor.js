import service from '@/utils/request'

//获取Monitor
export const getMonitorOne = (data) => {
    return service({
        url: "/stock/monitor/one",
        method: 'post',
        data
    })
}

//获取Monitor列表
export const getMonitorList = (data) => {
    return service({
        url: "/stock/monitor/list",
        method: 'post',
        data
    })
}

//新增Monitor
export const addMonitor = (data) => {
    return service({
        url: "/stock/monitor/add",
        method: 'post',
        data
    })
}

//删除Monitor
export const delMonitor = (data) => {
    return service({
        url: "/stock/monitor/del",
        method: 'post',
        data
    })
}

//更新Monitor
export const updateMonitor = (data) => {
    return service({
        url: "/stock/monitor/update",
        method: 'post',
        data
    })
}
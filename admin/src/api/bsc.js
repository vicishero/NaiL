import service from '@/utils/request'

export const createBscContractConfig = (data) => {
  return service({
    url: '/bsc/createBscContractConfig',
    method: 'post',
    data
  })
}

export const deleteBscContractConfig = (params) => {
  return service({
    url: '/bsc/deleteBscContractConfig',
    method: 'delete',
    params
  })
}

export const updateBscContractConfig = (data) => {
  return service({
    url: '/bsc/updateBscContractConfig',
    method: 'put',
    data
  })
}

export const findBscContractConfig = (params) => {
  return service({
    url: '/bsc/findBscContractConfig',
    method: 'get',
    params
  })
}

export const getBscContractConfigList = (params) => {
  return service({
    url: '/bsc/getBscContractConfigList',
    method: 'get',
    params
  })
}

// BSC Sync Info APIs
export const createBscSyncInfo = (data) => {
  return service({
    url: '/bscSyncInfo/createBscSyncInfo',
    method: 'post',
    data
  })
}

export const deleteBscSyncInfo = (params) => {
  return service({
    url: '/bscSyncInfo/deleteBscSyncInfo',
    method: 'delete',
    params
  })
}

export const updateBscSyncInfo = (data) => {
  return service({
    url: '/bscSyncInfo/updateBscSyncInfo',
    method: 'put',
    data
  })
}

export const findBscSyncInfo = (params) => {
  return service({
    url: '/bscSyncInfo/findBscSyncInfo',
    method: 'get',
    params
  })
}

export const getBscSyncInfoList = (params) => {
  return service({
    url: '/bscSyncInfo/getBscSyncInfoList',
    method: 'get',
    params
  })
}

// BSC Synced Event APIs
export const getBscSyncedEventList = (params) => {
  return service({
    url: '/bscSyncedEvent/getBscSyncedEventList',
    method: 'get',
    params
  })
}

export const updateBscSyncedEventStatus = (data) => {
  return service({
    url: '/bscSyncedEvent/updateStatus',
    method: 'put',
    data
  })
}

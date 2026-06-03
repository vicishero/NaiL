import service from '@/utils/request'

// --- H5 User Management ---
export const getH5UserList = (params) => service({ url: '/h5Admin/userList', method: 'get', params })
export const getH5User = (params) => service({ url: '/h5Admin/user', method: 'get', params })
export const updateH5User = (data) => service({ url: '/h5Admin/user', method: 'put', data })
export const deleteH5User = (data) => service({ url: '/h5Admin/user', method: 'delete', data })

// --- H5 Post Management ---
export const getH5PostList = (params) => service({ url: '/h5Admin/postList', method: 'get', params })
export const getH5Post = (params) => service({ url: '/h5Admin/post', method: 'get', params })
export const updateH5Post = (data) => service({ url: '/h5Admin/post', method: 'put', data })
export const deleteH5Post = (data) => service({ url: '/h5Admin/post', method: 'delete', data })

// --- H5 Comment Management ---
export const getH5CommentList = (params) => service({ url: '/h5Admin/commentList', method: 'get', params })
export const deleteH5Comment = (data) => service({ url: '/h5Admin/comment', method: 'delete', data })

// --- H5 Tag Management ---
export const getH5TagList = (params) => service({ url: '/h5Admin/tagList', method: 'get', params })
export const updateH5Tag = (data) => service({ url: '/h5Admin/tag', method: 'put', data })
export const deleteH5Tag = (data) => service({ url: '/h5Admin/tag', method: 'delete', data })

// --- H5 Collection Management ---
export const getH5CollectionList = (params) => service({ url: '/h5Admin/collectionList', method: 'get', params })
export const deleteH5Collection = (data) => service({ url: '/h5Admin/collection', method: 'delete', data })

// --- H5 Following Management ---
export const getH5FollowingList = (params) => service({ url: '/h5Admin/followingList', method: 'get', params })
export const deleteH5Following = (data) => service({ url: '/h5Admin/following', method: 'delete', data })

import api from './api'

export const userDashboardService = {
  // Dashboard
  getDashboard(userId) {
    return api.get(`/dashboard/${userId}`)
  },
  
  // Listings
  getListings(userId) {
    return api.get(`/dashboard/${userId}/listings`)
  },
  
  createListing(data) {
    return api.post('/listings', data)
  },
  
  // Messages
  getMessages(userId) {
    return api.get(`/dashboard/${userId}/messages`)
  },
  
  getConversations(userId) {
    return api.get(`/dashboard/${userId}/conversations`)
  },
  
  getConversationMessages(conversationId) {
    return api.get(`/dashboard/conversations/${conversationId}/messages`)
  },
  
  sendMessage(data) {
    return api.post('/dashboard/messages', data)
  },
  
  markConversationRead(conversationId) {
    return api.put(`/dashboard/conversations/${conversationId}/read`)
  },
  
  // Sales
  getSales(userId) {
    return api.get(`/dashboard/${userId}/sales`)
  },
  
  // Analytics
  getAnalytics(userId, range = '30d') {
    return api.get(`/dashboard/${userId}/analytics?range=${range}`)
  }
}

export const userService = {
  getAll() {
    return api.get('/users')
  },
  
  getById(id) {
    return api.get(`/users/${id}`)
  },
  
  create(userData) {
    return api.post('/users', userData)
  },
  
  updateProfile(userId, data) {
    return api.put(`/users/${userId}`, data)
  },
  
  uploadAvatar(formData) {
    return api.post('/users/avatar', formData, {
      headers: { 'Content-Type': 'multipart/form-data' }
    })
  },
  
  changePassword(data) {
    return api.put('/users/password', data)
  }
}

export const listingService = {
  getAll() { return api.get('/listings') },
  getById(id) { return api.get(`/listings/${id}`) },
  create(data) { return api.post('/listings', data) },
  update(id, data) { return api.put(`/listings/${id}`, data) },
  delete(id) { return api.delete(`/listings/${id}`) }
}

export const categoryService = {
  getAll() {
    return api.get('/categories')
  },
  
  getByType(type) {
    return api.get(`/categories?type=${type}`)
  }
}

export default api
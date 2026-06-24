import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:3000/api',
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  },
  timeout: 5000
})

export const userService = {
  getAll() {
    return api.get('/users')
  },
  getById(id) {
    return api.get(`/users/${id}`)
  },
  create(userData) {
    return api.post('/users', userData)
  }
}

export const listingService = {
  getAll() {
    return api.get('/listings')
  },
  getById(id) {
    return api.get(`/listings/${id}`)
  },
  create(listingData) {
    return api.post('/listings', listingData)
  }
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
import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '../views/Landing_page.vue'
import MarketPlace from '../views/Market_place.vue'
import RegisterLoginView from '../views/Register_login.vue'

const routes = [
  // Public routes
  {
    path: '/',
    name: 'landing',
    component: LandingPage
  },
  {
    path: '/register-login',
    name: 'register-login',
    component: RegisterLoginView
  },
  {
    path: '/market',
    name: 'market',
    component: MarketPlace
  },
  
  // Admin routes
  {
    path: '/admin',
    name: 'admin-dashboard',
    component: () => import('../views/admin/admin_dashboard.vue'),
    meta: { requiresAuth: true, requiresAdmin: true }
  },
  
  // User routes - FLAT, no nesting, each page handles its own layout
  {
    path: '/dashboard',
    name: 'user-dashboard',
    component: () => import('../views/users/user_dashboard.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/dashboard/listings',
    name: 'my-listings',
    component: () => import('../views/users/my_listings.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/dashboard/sales',
    name: 'sales',
    component: () => import('../views/users/sales.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/dashboard/analytics',
    name: 'analytics',
    component: () => import('../views/users/analytics.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/dashboard/messages',
    name: 'messages',
    component: () => import('../views/users/messages.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/settings',
    name: 'profile-settings',
    component: () => import('../views/users/profile_settings.vue'),
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (to.hash) {
      return {
        el: to.hash,
        behavior: 'smooth',
      }
    }
    return { top: 0 }
  }
})

// Navigation guard
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  const userStr = localStorage.getItem('user')
  let userRole = null
  
  if (userStr) {
    try {
      const user = JSON.parse(userStr)
      userRole = user.role
    } catch (e) {
      console.error('Error parsing user data', e)
    }
  }

  // If route requires authentication
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!token) {
      next({
        name: 'register-login',
        query: { redirect: to.fullPath }
      })
      return
    }
    
    if (to.meta.requiresAdmin && userRole !== 'admin') {
      next({ name: 'user-dashboard' })
      return
    }
  }
  
  // Redirect logged-in users away from login page
  if (to.name === 'register-login' && token) {
    if (userRole === 'admin') {
      next({ name: 'admin-dashboard' })
    } else {
      next({ name: 'user-dashboard' })
    }
    return
  }
  
  next()
})

export default router
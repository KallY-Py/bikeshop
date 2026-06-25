import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '../views/Landing_page.vue'
import MarketPlace from '../views/Market_place.vue'
import RegisterLoginView from '../views/Register_login.vue'

import AdminDashboard from '../views/admin/admin_dashboard.vue'

// import UserDashboard from '../views/users/user_dashboard.vue'

const routes = [
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
  {
    path: '/admin/dashboard',
    name: 'admin-dashboard',
    component: AdminDashboard,
    meta: { requiresAuth: true, requiresAdmin: true }
  },
  // {
  //   path: '/user/dashboard',
  //   name: 'user-dashboard',
  //   component: UserDashboard,
  //   meta: { requiresAuth: true }
  // }
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

  if (to.meta.requiresAuth) {
    if (!token) {
      next('/register-login')
    } else if (to.meta.requiresAdmin && userRole !== 'admin') {
      next('/user/dashboard')
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router
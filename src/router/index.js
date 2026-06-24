import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '../views/Landing_page.vue'
import MarketPlace from '../views/Market_place.vue'
import RegisterLoginView from '../views/Register_login.vue'

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

export default router
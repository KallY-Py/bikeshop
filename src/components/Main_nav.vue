<template>
  <nav class="fixed top-0 left-0 right-0 bg-gray-900/95 backdrop-blur-sm border-b border-gray-800 z-50">
    <div class="max-w-7xl mx-auto px-2 sm:px-4 lg:px-6">
      <div class="flex items-center justify-between h-16">
        <!-- Logo - Top Left -->
        <div class="flex-shrink-0">
          <router-link to="/" class="text-2xl font-bold text-orange-500 tracking-wider">
            PADYAK<span class="text-white"> ALL</span>
          </router-link>
        </div>
        
        <!-- Desktop Menu -->
        <div class="hidden md:block">
          <div class="ml-10 flex items-baseline space-x-4">
            <router-link 
              to="/" 
              class="nav-link"
              active-class="active-link"
              exact-active-class="active-link"
            >
              Home
            </router-link>
            <router-link 
              to="/market" 
              class="nav-link"
              active-class="active-link"
            >
              Market Place
            </router-link>
            <router-link 
              to="/#road" 
              class="nav-link"
              active-class="active-link"
            >
              Road
            </router-link>
            <router-link 
              to="/#mtb" 
              class="nav-link"
              active-class="active-link"
            >
              MTB
            </router-link>
            <router-link 
              to="/#bmx" 
              class="nav-link"
              active-class="active-link"
            >
              BMX
            </router-link>
            <router-link 
              to="/#gravel" 
              class="nav-link"
              active-class="active-link"
            >
              Gravel
            </router-link>
          </div>
        </div>

        <!-- Search Bar -->
        <div class="hidden md:block flex-1 max-w-md mx-4">
          <div class="relative">
            <input 
              type="text" 
              placeholder="Search bikes..." 
              class="w-full bg-gray-800 text-white pl-10 pr-4 py-2 rounded-lg border border-gray-700 focus:outline-none focus:border-orange-500 transition-colors"
            />
            <svg class="absolute left-3 top-2.5 h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
            </svg>
          </div>
        </div>

        <!-- Auth Buttons -->
        <div class="flex items-center space-x-3">
          <router-link to="/register-login">
            <button class="text-white hover:text-orange-500 transition-colors text-sm font-medium">
              Log In
            </button>
          </router-link>
          <router-link to="/register-login">
            <button class="bg-orange-500 hover:bg-orange-600 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors">
              Sign Up
            </button>
          </router-link>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

// Function to scroll to section if hash is present
const scrollToSection = () => {
  const hash = route.hash
  if (hash) {
    const element = document.querySelector(hash)
    if (element) {
      setTimeout(() => {
        element.scrollIntoView({ behavior: 'smooth' })
      }, 200)
    }
  }
}

onMounted(() => {
  // Add scroll listener for navbar background
  const nav = document.querySelector('nav')
  window.addEventListener('scroll', () => {
    if (window.scrollY > 50) {
      nav?.classList.add('shadow-lg')
    } else {
      nav?.classList.remove('shadow-lg')
    }
  })

  // Scroll to section on initial load if hash exists
  scrollToSection()
})

// Watch for route changes to handle hash navigation
watch(
  () => route.path + route.hash,
  () => {
    scrollToSection()
  }
)
</script>

<style scoped>
/* Navigation link styles */
.nav-link {
  @apply px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200;
  color: #9ca3af; /* text-gray-400 */
}

.nav-link:hover {
  color: #f97316; /* text-orange-500 */
}

/* Active link styles */
.active-link {
  color: #f97316 !important; /* text-orange-500 */
  background-color: rgba(249, 115, 22, 0.1); /* orange-500/10 */
  position: relative;
}

/* Optional: Add an underline indicator for active link */
.active-link::after {
  content: '';
  position: absolute;
  bottom: 2px;
  left: 50%;
  transform: translateX(-50%);
  width: 20px;
  height: 2px;
  background-color: #f97316;
  border-radius: 2px;
}

/* Ensure the link stays relative for the underline */
.nav-link {
  position: relative;
}

/* Custom scrollbar */
::-webkit-scrollbar {
  width: 10px;
}

::-webkit-scrollbar-track {
  background: #111827;
}

::-webkit-scrollbar-thumb {
  background: #f97316;
  border-radius: 5px;
}

::-webkit-scrollbar-thumb:hover {
  background: #ea580c;
}

/* Smooth scroll */
html {
  scroll-behavior: smooth;
}
</style>
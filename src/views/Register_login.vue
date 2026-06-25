<template>
  <div class="min-h-screen bg-gray-900 text-white p-4 md:p-8 flex items-center justify-center">
    <div class="w-full max-w-md relative">
      
      <transition name="flip-switch" mode="out-in">
        
        <div v-if="!isFlipped" key="register" class="bg-gray-800 rounded-2xl shadow-2xl p-8 border border-gray-700 w-full">
          <h2 class="text-2xl font-bold text-center mb-6 text-orange-500">Create Account</h2>
          <form @submit.prevent="handleRegister" class="space-y-4">
            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-sm font-medium text-gray-300 mb-1">First Name</label>
                <input v-model="registerForm.first_name" type="text" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-300 mb-1">Last Name</label>
                <input v-model="registerForm.last_name" type="text" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
              </div>
            </div>
            
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-1">Username</label>
              <input v-model="registerForm.username" type="text" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-300 mb-1">Email</label>
              <input v-model="registerForm.email" type="email" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-300 mb-1">Phone Number</label>
              <input v-model="registerForm.phone_number" type="tel" class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-300 mb-1">Password</label>
              <input v-model="registerForm.password" type="password" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-300 mb-1">Confirm Password</label>
              <input v-model="registerForm.confirm_password" type="password" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
            </div>

            <button type="submit" class="w-full bg-orange-500 hover:bg-orange-600 text-white font-semibold py-2.5 rounded-lg transition-colors mt-2">
              Create Account
            </button>
          </form>
          
          <p class="text-center text-sm text-gray-400 mt-6">
            Already have an account? 
            <button @click="isFlipped = true" class="text-orange-500 font-semibold hover:underline">Login account</button>
          </p>
        </div>

        <div v-else key="login" class="bg-gray-800 rounded-2xl shadow-2xl p-6 border border-gray-700 w-full">
          <h2 class="text-2xl font-bold text-center mb-4 text-orange-500">Login Account</h2>
          <form @submit.prevent="handleLogin" class="space-y-4">
            <div>
              <label class="block text-sm font-medium text-gray-300 mb-1">Email or Username</label>
              <input v-model="loginForm.identifier" type="text" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-300 mb-1">Password</label>
              <input v-model="loginForm.password" type="password" required class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-orange-500 transition-colors" />
            </div>

            <button type="submit" :disabled="loading" class="w-full bg-orange-500 hover:bg-orange-600 text-white font-semibold py-2.5 rounded-lg transition-colors mt-4 disabled:opacity-50 disabled:cursor-not-allowed">
              {{ loading ? 'Logging in...' : 'Log In' }}
            </button>
          </form>
          
          <p class="text-center text-sm text-gray-400 mt-6">
            Don't have an account? 
            <button @click="isFlipped = false" class="text-orange-500 font-semibold hover:underline">Sign up</button>
          </p>
        </div>

      </transition>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'RegisterLoginView',
  
  data() {
    return {
      isFlipped: false, 
      registerForm: {
        first_name: '',
        last_name: '',
        username: '',
        email: '',
        phone_number: '',
        password: '',
        confirm_password: ''
      },
      loginForm: {
        identifier: '',
        password: ''
      },
      connected: false,
      users: [],
      categories: [],
      error: null,
      loading: false
    }
  },
  
  async mounted() {
    await this.checkConnection()
  },
  
  methods: {
    async handleRegister() {
      if (this.registerForm.password !== this.registerForm.confirm_password) {
        alert('Passwords do not match!')
        return
      }

      try {
        const response = await axios.post('http://localhost:3000/api/auth/register', {
          first_name: this.registerForm.first_name,
          last_name: this.registerForm.last_name,
          username: this.registerForm.username,
          email: this.registerForm.email,
          phone_number: this.registerForm.phone_number || '',
          password: this.registerForm.password
        })

        console.log('Registration successful:', response.data)
        alert('Registration successful! Please login.')
        
        // Reset form
        this.registerForm = {
          first_name: '',
          last_name: '',
          username: '',
          email: '',
          phone_number: '',
          password: '',
          confirm_password: ''
        }
        
        // Switch to login view after successful registration
        this.isFlipped = true
        
      } catch (error) {
        console.error('Registration error:', error)
        let errorMessage = 'Registration failed. Please try again.'
        
        if (error.response && error.response.data) {
          errorMessage = error.response.data.message || error.response.data.error || errorMessage
        } else if (error.message) {
          errorMessage = error.message
        }
        
        alert(errorMessage)
      }
    },
    
    async handleLogin() {
      this.loading = true
      this.error = null
      
      try {
        const response = await axios.post('http://localhost:3000/api/auth/login', {
          identifier: this.loginForm.identifier,
          password: this.loginForm.password
        })

        console.log('Login successful:', response.data)
        
        // Store token and user data
        localStorage.setItem('token', response.data.token)
        localStorage.setItem('user', JSON.stringify(response.data.user))
        
        // Role-based redirection
        const userRole = response.data.user.role
        
        if (userRole === 'admin') {
          this.$router.push('/admin/dashboard')
        } else {
          // Default to user dashboard for 'user' role or any other role
          this.$router.push('/user/dashboard')
        }
        
      } catch (error) {
        console.error('Login error:', error)
        let errorMessage = 'Login failed. Please check your credentials.'
        
        if (error.response && error.response.data) {
          errorMessage = error.response.data.message || error.response.data.error || errorMessage
        }
        
        alert(errorMessage)
      } finally {
        this.loading = false
      }
    },

    async checkConnection() {
      this.error = null
      try {
        const response = await axios.get('http://localhost:3000/api/health')
        this.connected = response.data.status === 'ok'
        console.log('Connection check:', response.data)
      } catch (err) {
        this.connected = false
        this.error = `Connection failed: ${err.message}`
        console.error('Connection error:', err)
      }
    },
    
    async loadUsers() {
      this.error = null
      this.loading = true
      try {
        const response = await axios.get('http://localhost:3000/api/users')
        this.users = response.data || []
        console.log('Users loaded:', this.users)
      } catch (err) {
        this.error = `Failed to load users: ${err.message}`
        this.users = []
        console.error('Users error:', err)
      } finally {
        this.loading = false
      }
    },
    
    async loadCategories() {
      this.error = null
      this.loading = true
      try {
        const response = await axios.get('http://localhost:3000/api/categories')
        this.categories = response.data || []
        console.log('Categories loaded:', this.categories)
      } catch (err) {
        this.error = `Failed to load categories: ${err.message}`
        this.categories = []
        console.error('Categories error:', err)
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

<style scoped>
/* Custom Flip/Switch Animation for Separate Containers */
.flip-switch-enter-active,
.flip-switch-leave-active {
  transition: all 0.4s ease;
}

.flip-switch-enter-from {
  opacity: 0;
  transform: scale(0.95) rotateY(-10deg);
}

.flip-switch-leave-to {
  opacity: 0;
  transform: scale(0.95) rotateY(10deg);
}

.flip-switch-enter-to,
.flip-switch-leave-from {
  opacity: 1;
  transform: scale(1) rotateY(0);
}
</style>
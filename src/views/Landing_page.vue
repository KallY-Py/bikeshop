<template>
  <div class="min-h-screen bg-gray-900 text-white">
    <!-- Navigation Bar Component -->
    <Main_nav />

    <!-- Hero Section -->
    <section id="home" class="relative pt-32 pb-20 lg:pt-48 lg:pb-32 overflow-hidden">
      <div class="absolute inset-0 bg-gradient-to-br from-gray-900 via-gray-900 to-orange-900/20"></div>
      <div class="relative max-w-7xl mx-auto px-2 sm:px-4 lg:px-6 text-center">
        <div class="inline-block mb-4 px-4 py-1 rounded-full border border-orange-500/30 bg-orange-500/10">
          <span class="text-orange-500 text-sm font-semibold">🚴 The #1 Cycling Marketplace</span>
        </div>
        <h1 class="text-5xl md:text-7xl font-bold mb-6 leading-tight">
          Find Your Perfect <span class="text-orange-500">Ride</span>
        </h1>
        <p class="text-xl text-gray-400 mb-10 max-w-2xl mx-auto">
          Buy and sell bikes, parts, and accessories. Thousands of listings from passionate cyclists near you.
        </p>
        
        <!-- Search Box -->
        <div class="max-w-2xl mx-auto mb-12">
          <div class="flex bg-gray-800 rounded-lg p-2 border border-gray-700">
            <input 
              type="text" 
              placeholder="Search bikes, brands, parts..." 
              class="flex-1 bg-transparent text-white px-4 py-3 focus:outline-none"
            />
            <button class="bg-orange-500 hover:bg-orange-600 text-white px-8 py-3 rounded-md font-semibold transition-colors">
              Search
            </button>
          </div>
        </div>

        <!-- Stats -->
        <div class="flex justify-center items-center space-x-6 text-sm">
          <div class="text-center">
            <span class="block text-2xl font-bold text-orange-500">500+</span>
            <span class="text-gray-400">Listings</span>
          </div>
          <div class="text-center">
            <span class="block text-2xl font-bold text-orange-500">200+</span>
            <span class="text-gray-400">Sellers</span>
          </div>
          <div class="text-center">
            <span class="block text-2xl font-bold text-orange-500">4.9</span>
            <span class="text-gray-400">Avg Rating</span>
          </div>
        </div>
      </div>
    </section>

    <!-- Unleash Potential Section -->
    <section class="py-20 bg-gray-800/50">
      <div class="max-w-7xl mx-auto px-2 sm:px-4 lg:px-6">
        <div class="grid md:grid-cols-2 gap-12 items-center">
          <div>
            <h2 class="text-4xl font-bold mb-4">
              UNLEASH YOUR <span class="text-orange-500">POTENTIAL</span>
            </h2>
            <p class="text-gray-400 mb-8 text-lg">
              The ultimate marketplace for high-performance bicycles, precision components, and elite gear. Buy, sell, and dominate the road.
            </p>
            <div class="flex space-x-4">
              <button class="bg-orange-500 hover:bg-orange-600 text-white px-8 py-3 rounded-lg font-semibold transition-colors">
                Shop Now
              </button>
              <button class="border border-gray-600 hover:border-orange-500 text-white px-8 py-3 rounded-lg font-semibold transition-colors">
                Sell Your Bike
              </button>
            </div>
          </div>
          <div class="relative">
            <div class="absolute inset-0 bg-orange-500/20 blur-3xl rounded-full"></div>
            <img 
              src="https://images.unsplash.com/photo-1485965120184-e224f723d1e1?w=800&h=600&fit=crop" 
              alt="Premium Bike" 
              class="relative rounded-lg shadow-2xl"
            />
          </div>
        </div>
      </div>
    </section>

    <!-- Featured Listings -->
    <section id="bikes" class="py-20">
      <div class="max-w-7xl mx-auto px-2 sm:px-4 lg:px-6">
        <div class="flex justify-between items-center mb-12">
          <div>
            <h2 class="text-3xl font-bold mb-2">FEATURED LISTINGS</h2>
            <p class="text-gray-400">Top-of-the-line handpicked for performance</p>
          </div>
          <button class="text-orange-500 hover:text-orange-400 font-semibold">View All →</button>
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
          <!-- Bike Card 1 -->
          <div class="bg-gray-800 rounded-lg overflow-hidden border border-gray-700 hover:border-orange-500/50 transition-all hover:transform hover:scale-105">
            <img 
              src="https://images.unsplash.com/photo-1532298229144-0ec0c57515c7?w=600&h=400&fit=crop" 
              alt="Mountain Bike" 
              class="w-full h-48 object-cover"
            />
            <div class="p-5">
              <h3 class="text-xl font-bold mb-2">S-Works Stumpjumper EVO</h3>
              <p class="text-gray-400 text-sm mb-4">2024 • Mountain Bike • Carbon Frame</p>
              <div class="flex justify-between items-center">
                <span class="text-2xl font-bold text-orange-500">$8,500</span>
                <button class="text-sm text-gray-400 hover:text-white">Details</button>
              </div>
            </div>
          </div>

          <!-- Bike Card 2 -->
          <div class="bg-gray-800 rounded-lg overflow-hidden border border-gray-700 hover:border-orange-500/50 transition-all hover:transform hover:scale-105">
            <img 
              src="https://images.unsplash.com/photo-1571068316344-75bc76f77890?w=600&h=400&fit=crop" 
              alt="Road Bike" 
              class="w-full h-48 object-cover"
            />
            <div class="p-5">
              <h3 class="text-xl font-bold mb-2">Trek Domane SL 7</h3>
              <p class="text-gray-400 text-sm mb-4">2024 • Road Bike • Endurance</p>
              <div class="flex justify-between items-center">
                <span class="text-2xl font-bold text-orange-500">$6,200</span>
                <button class="text-sm text-gray-400 hover:text-white">Details</button>
              </div>
            </div>
          </div>

          <!-- Bike Card 3 -->
          <div class="bg-gray-800 rounded-lg overflow-hidden border border-gray-700 hover:border-orange-500/50 transition-all hover:transform hover:scale-105">
            <img 
              src="https://images.unsplash.com/photo-1576435728678-38d01d52e0a5?w=600&h=400&fit=crop" 
              alt="Gravel Bike" 
              class="w-full h-48 object-cover"
            />
            <div class="p-5">
              <h3 class="text-xl font-bold mb-2">Cannondale Topstone</h3>
              <p class="text-gray-400 text-sm mb-4">2024 • Gravel Bike • Adventure</p>
              <div class="flex justify-between items-center">
                <span class="text-2xl font-bold text-orange-500">$3,800</span>
                <button class="text-sm text-gray-400 hover:text-white">Details</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Categories Section - The Armory -->
    <section class="py-20 bg-gray-800/50">
      <div class="max-w-7xl mx-auto px-2 sm:px-4 lg:px-6">
        <div class="text-center mb-12">
          <h2 class="text-3xl font-bold mb-4">THE ARMORY</h2>
          <p class="text-gray-400">Select your weapon. We have the finest machines for every terrain.</p>
        </div>

        <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
          <!-- Category 1 -->
          <div class="group relative overflow-hidden rounded-lg cursor-pointer">
            <img 
              src="https://images.unsplash.com/photo-1517649763962-0c623066013b?w=400&h=500&fit=crop" 
              alt="Road Bikes" 
              class="w-full h-80 object-cover group-hover:scale-110 transition-transform duration-300"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>
            <div class="absolute bottom-0 left-0 right-0 p-4">
              <h3 class="text-xl font-bold mb-2">ROAD</h3>
              <p class="text-gray-300 text-sm">Speed & Performance</p>
            </div>
          </div>

          <!-- Category 2 -->
          <div class="group relative overflow-hidden rounded-lg cursor-pointer">
            <img 
              src="https://images.unsplash.com/photo-1544197150-b99a580bb7a8?w=400&h=500&fit=crop" 
              alt="Mountain Bikes" 
              class="w-full h-80 object-cover group-hover:scale-110 transition-transform duration-300"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>
            <div class="absolute bottom-0 left-0 right-0 p-4">
              <h3 class="text-xl font-bold mb-2">MTB</h3>
              <p class="text-gray-300 text-sm">Trail & Downhill</p>
            </div>
          </div>

          <!-- Category 3 -->
          <div class="group relative overflow-hidden rounded-lg cursor-pointer">
            <img 
              src="https://images.unsplash.com/photo-1507035895480-2b3156c31fc8?w=400&h=500&fit=crop" 
              alt="Gravel Bikes" 
              class="w-full h-80 object-cover group-hover:scale-110 transition-transform duration-300"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>
            <div class="absolute bottom-0 left-0 right-0 p-4">
              <h3 class="text-xl font-bold mb-2">GRAVEL</h3>
              <p class="text-gray-300 text-sm">Adventure & Exploration</p>
            </div>
          </div>

          <!-- Category 4 -->
          <div class="group relative overflow-hidden rounded-lg cursor-pointer">
            <img 
              src="https://images.unsplash.com/photo-1558521766-522130451a7a?w=400&h=500&fit=crop" 
              alt="BMX Bikes" 
              class="w-full h-80 object-cover group-hover:scale-110 transition-transform duration-300"
            />
            <div class="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>
            <div class="absolute bottom-0 left-0 right-0 p-4">
              <h3 class="text-xl font-bold mb-2">BMX</h3>
              <p class="text-gray-300 text-sm">Tricks & Street</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA Section -->
    <section class="py-20">
      <div class="max-w-5xl mx-auto px-2 sm:px-4 lg:px-6">
        <div class="bg-gradient-to-r from-orange-600 to-orange-500 rounded-2xl px-8 py-12 text-center">
          <h2 class="text-4xl font-bold mb-4">Ready to Start Your Journey?</h2>
          <p class="text-xl mb-8 text-orange-100">Join thousands of cyclists who trust PADYAK ALL for their cycling needs.</p>
          <div class="flex justify-center space-x-4">
            <button class="bg-white text-orange-600 px-8 py-3 rounded-lg font-bold hover:bg-gray-100 transition-colors">
              Get Started
            </button>
            <button class="border-2 border-white text-white px-8 py-3 rounded-lg font-bold hover:bg-white/10 transition-colors">
              Learn More
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- Footer -->
    <footer class="bg-gray-900 border-t border-gray-800 py-12">
      <div class="max-w-7xl mx-auto px-2 sm:px-4 lg:px-6">
        <div class="grid md:grid-cols-4 gap-6 mb-8">
          <div>
            <h3 class="text-2xl font-bold text-orange-500 mb-4">PADYAK ALL</h3>
            <p class="text-gray-400 text-sm">Your premier destination for buying and selling quality bicycles.</p>
          </div>
          <div>
            <h4 class="font-bold mb-4">Quick Links</h4>
            <ul class="space-y-2 text-gray-400 text-sm">
              <li><a href="#" class="hover:text-orange-500 transition-colors">Browse Bikes</a></li>
              <li><a href="#" class="hover:text-orange-500 transition-colors">Sell Your Bike</a></li>
              <li><a href="#" class="hover:text-orange-500 transition-colors">How It Works</a></li>
            </ul>
          </div>
          <div>
            <h4 class="font-bold mb-4">Categories</h4>
            <ul class="space-y-2 text-gray-400 text-sm">
              <li><a href="#" class="hover:text-orange-500 transition-colors">Road Bikes</a></li>
              <li><a href="#" class="hover:text-orange-500 transition-colors">Mountain Bikes</a></li>
              <li><a href="#" class="hover:text-orange-500 transition-colors">Gravel Bikes</a></li>
            </ul>
          </div>
          <div>
            <h4 class="font-bold mb-4">Contact</h4>
            <ul class="space-y-2 text-gray-400 text-sm">
              <li>support@padyakall.com</li>
              <li>+63 123 456 7890</li>
              <li>Manila, Philippines</li>
            </ul>
          </div>
        </div>
        <div class="border-t border-gray-800 pt-8 text-center text-gray-400 text-sm">
          <p>&copy; 2026 PADYAK ALL. All rights reserved.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import Main_nav from '@/components/Main_nav.vue' 
// Note: If you don't have the '@' alias configured in your vite.config.js or vue.config.js, 
// change the import path to: import Main_nav from '../components/Main_nav.vue'
</script>

<style scoped>
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
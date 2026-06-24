<template>
  <div class="min-h-screen bg-gray-900 text-white">
    <!-- Navigation Bar Component -->
    <Main_nav />

    <!-- Main Content -->
    <main class="flex-grow pt-24 pb-20 md:pb-8 max-w-7xl mx-auto w-full px-2 sm:px-4 lg:px-6 flex gap-6">
      <!-- Sidebar Filter (Desktop) -->
      <aside class="hidden md:flex flex-col w-[280px] flex-shrink-0 sticky top-24 h-[calc(100vh-96px)] overflow-y-auto custom-scrollbar pr-4">
        <div class="flex items-center justify-between mb-6">
          <h2 class="text-2xl font-bold text-white">Filters</h2>
          <button class="text-orange-500 text-sm font-medium hover:underline">Clear all</button>
        </div>
        
        <!-- Categories -->
        <div class="mb-8 border-b border-gray-700 pb-6">
          <h3 class="text-sm font-semibold text-gray-400 mb-4 uppercase tracking-wider">Category</h3>
          <ul class="space-y-3 text-base text-white">
            <li>
              <label class="flex items-center gap-3 cursor-pointer group">
                <input checked class="text-orange-500 focus:ring-orange-500 bg-gray-800 border-gray-700 rounded-full h-4 w-4" name="category" type="radio"/>
                <span class="group-hover:text-orange-500 transition-colors font-medium">All Bikes</span>
              </label>
            </li>
            <li>
              <label class="flex items-center gap-3 cursor-pointer group">
                <input class="text-orange-500 focus:ring-orange-500 bg-gray-800 border-gray-700 rounded-full h-4 w-4" name="category" type="radio"/>
                <span class="group-hover:text-orange-500 transition-colors font-medium">Road</span>
              </label>
            </li>
            <li>
              <label class="flex items-center gap-3 cursor-pointer group">
                <input class="text-orange-500 focus:ring-orange-500 bg-gray-800 border-gray-700 rounded-full h-4 w-4" name="category" type="radio"/>
                <span class="group-hover:text-orange-500 transition-colors font-medium">MTB</span>
              </label>
            </li>
            <li>
              <label class="flex items-center gap-3 cursor-pointer group">
                <input class="text-orange-500 focus:ring-orange-500 bg-gray-800 border-gray-700 rounded-full h-4 w-4" name="category" type="radio"/>
                <span class="group-hover:text-orange-500 transition-colors font-medium">Gravel</span>
              </label>
            </li>
            <li>
              <label class="flex items-center gap-3 cursor-pointer group">
                <input class="text-orange-500 focus:ring-orange-500 bg-gray-800 border-gray-700 rounded-full h-4 w-4" name="category" type="radio"/>
                <span class="group-hover:text-orange-500 transition-colors font-medium">BMX</span>
              </label>
            </li>
          </ul>
        </div>

        <!-- Price Range -->
        <div class="mb-8 border-b border-gray-700 pb-6">
          <h3 class="text-sm font-semibold text-gray-400 mb-4 uppercase tracking-wider">Price Range</h3>
          <div class="flex items-center gap-2">
            <div class="relative w-full">
              <span class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 text-sm">$</span>
              <input class="w-full bg-gray-800 border border-gray-700 rounded-lg pl-8 pr-2 py-2.5 text-white text-base focus:border-orange-500 focus:ring-1 focus:ring-orange-500 outline-none transition-colors" placeholder="Min" type="number"/>
            </div>
            <span class="text-gray-400">-</span>
            <div class="relative w-full">
              <span class="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 text-sm">$</span>
              <input class="w-full bg-gray-800 border border-gray-700 rounded-lg pl-8 pr-2 py-2.5 text-white text-base focus:border-orange-500 focus:ring-1 focus:ring-orange-500 outline-none transition-colors" placeholder="Max" type="number"/>
            </div>
          </div>
        </div>

        <!-- Condition -->
        <div class="mb-8 border-b border-gray-700 pb-6">
          <h3 class="text-sm font-semibold text-gray-400 mb-4 uppercase tracking-wider">Condition</h3>
          <div class="flex flex-wrap gap-2">
            <button class="px-4 py-2 rounded-lg border border-orange-500 bg-orange-500/10 text-orange-500 text-sm font-medium transition-colors">Any</button>
            <button class="px-4 py-2 rounded-lg border border-gray-700 text-gray-400 hover:border-orange-500 hover:text-orange-500 text-sm font-medium transition-colors">New</button>
            <button class="px-4 py-2 rounded-lg border border-gray-700 text-gray-400 hover:border-orange-500 hover:text-orange-500 text-sm font-medium transition-colors">Used - Like New</button>
            <button class="px-4 py-2 rounded-lg border border-gray-700 text-gray-400 hover:border-orange-500 hover:text-orange-500 text-sm font-medium transition-colors">Used - Good</button>
          </div>
        </div>

        <!-- Frame Size -->
        <div class="mb-8">
          <h3 class="text-sm font-semibold text-gray-400 mb-4 uppercase tracking-wider">Frame Size</h3>
          <div class="grid grid-cols-3 gap-2">
            <button class="py-2.5 border border-gray-700 rounded-lg text-gray-400 hover:border-orange-500 hover:text-orange-500 text-sm font-medium transition-colors">XS</button>
            <button class="py-2.5 border border-gray-700 rounded-lg text-gray-400 hover:border-orange-500 hover:text-orange-500 text-sm font-medium transition-colors">S</button>
            <button class="py-2.5 border border-orange-500 bg-orange-500/5 rounded-lg text-orange-500 text-sm font-medium transition-colors">M</button>
            <button class="py-2.5 border border-gray-700 rounded-lg text-gray-400 hover:border-orange-500 hover:text-orange-500 text-sm font-medium transition-colors">L</button>
            <button class="py-2.5 border border-gray-700 rounded-lg text-gray-400 hover:border-orange-500 hover:text-orange-500 text-sm font-medium transition-colors">XL</button>
            <button class="py-2.5 border border-gray-700 rounded-lg text-gray-400 hover:border-orange-500 hover:text-orange-500 text-sm font-medium transition-colors">XXL</button>
          </div>
        </div>

        <button class="w-full bg-orange-500 hover:bg-orange-600 text-white font-semibold py-3 rounded-lg transition-colors mt-4">
          Apply Filters
        </button>
      </aside>

      <!-- Product Grid -->
      <div class="flex-grow pt-4 md:pt-8 w-full">
        <!-- Sorting & Results count -->
        <div class="flex justify-between items-center mb-8">
          <h1 class="text-3xl font-bold text-white">Explore Bikes</h1>
          <div class="flex items-center gap-2 text-gray-400 text-base">
            <span class="hidden sm:inline font-medium">Sort by:</span>
            <select class="bg-gray-800 border border-gray-700 rounded-lg focus:ring-0 text-white cursor-pointer px-3 py-2 focus:border-orange-500 outline-none transition-colors">
              <option>Featured</option>
              <option>Price: Low to High</option>
              <option>Price: High to Low</option>
              <option>Newest Arrivals</option>
            </select>
          </div>
        </div>

        <!-- Grid -->
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <!-- Card 1 (Featured - Large) -->
          <div class="col-span-1 sm:col-span-2 lg:col-span-2 group bg-gray-800 rounded-xl overflow-hidden border border-gray-700 hover:border-orange-500/50 transition-all flex flex-col sm:flex-row relative">
            <div class="absolute top-4 right-4 z-10">
              <button class="w-10 h-10 rounded-full bg-gray-900/80 backdrop-blur-sm flex items-center justify-center text-gray-400 hover:text-orange-500 transition-all">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" /></svg>
              </button>
            </div>
            <div class="w-full sm:w-1/2 aspect-video sm:aspect-auto relative bg-gray-900">
              <img class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" src="https://images.unsplash.com/photo-1532298229144-0ec0c57515c7?w=800&h=600&fit=crop" alt="Road Bike">
              <div class="absolute top-4 left-4">
                <span class="bg-green-500/20 text-green-400 text-xs font-semibold px-2.5 py-1 rounded-md uppercase tracking-wider backdrop-blur-sm border border-green-500/30">Verified</span>
              </div>
            </div>
            <div class="w-full sm:w-1/2 p-6 flex flex-col justify-between">
              <div>
                <div class="flex justify-between items-start mb-2">
                  <span class="text-orange-500 text-xs font-semibold uppercase tracking-wider">Road Bike</span>
                  <span class="text-gray-400 text-xs font-medium flex items-center gap-1">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" /></svg>
                    Portland, OR
                  </span>
                </div>
                <h3 class="text-xl font-bold text-white mb-2 group-hover:text-orange-500 transition-colors">Specialized Tarmac SL7 Pro - 2023</h3>
                <p class="text-gray-400 text-base line-clamp-2 mb-4">Pristine condition. Upgraded with Roval Rapide CLX wheels. Ultegra Di2 12-speed group. Never crashed.</p>
              </div>
              <div class="flex items-end justify-between border-t border-gray-700 pt-4">
                <div>
                  <span class="text-gray-400 text-sm line-through block mb-1">$7,500</span>
                  <span class="text-2xl font-bold text-orange-500">$6,200</span>
                </div>
                <button class="bg-gray-700 hover:bg-orange-500 hover:text-white text-white px-4 py-2 rounded-lg font-semibold text-sm transition-colors">View Details</button>
              </div>
            </div>
          </div>

          <!-- Card 2 -->
          <div class="col-span-1 group bg-gray-800 rounded-xl overflow-hidden border border-gray-700 hover:border-orange-500/50 transition-all flex flex-col relative">
            <div class="absolute top-4 right-4 z-10">
              <button class="w-8 h-8 rounded-full bg-gray-900/80 backdrop-blur-sm flex items-center justify-center text-gray-400 hover:text-orange-500 transition-all">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" /></svg>
              </button>
            </div>
            <div class="aspect-[4/3] w-full relative bg-gray-900 overflow-hidden">
              <img class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" src="https://images.unsplash.com/photo-1571068316344-75bc76f77890?w=600&h=400&fit=crop" alt="MTB">
            </div>
            <div class="p-5 flex flex-col flex-grow">
              <div class="flex justify-between items-center mb-2">
                <span class="text-orange-500 text-xs font-semibold uppercase tracking-wider">MTB</span>
                <span class="text-gray-400 text-xs font-medium">Size: L</span>
              </div>
              <h3 class="text-lg font-bold text-white mb-1 group-hover:text-orange-500 transition-colors">Trek Slash 9.8 XT</h3>
              <div class="text-gray-400 text-xs font-medium flex items-center gap-1 mb-4">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" /></svg>
                Denver, CO
              </div>
              <div class="mt-auto flex justify-between items-end">
                <span class="text-2xl font-bold text-orange-500">$4,500</span>
              </div>
            </div>
          </div>

          <!-- Card 3 -->
          <div class="col-span-1 group bg-gray-800 rounded-xl overflow-hidden border border-gray-700 hover:border-orange-500/50 transition-all flex flex-col relative">
            <div class="absolute top-4 right-4 z-10">
              <button class="w-8 h-8 rounded-full bg-gray-900/80 backdrop-blur-sm flex items-center justify-center text-orange-500 transition-all">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="currentColor" viewBox="0 0 24 24"><path d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" /></svg>
              </button>
            </div>
            <div class="aspect-[4/3] w-full relative bg-gray-900 overflow-hidden">
              <img class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" src="https://images.unsplash.com/photo-1576435728678-38d01d52e0a5?w=600&h=400&fit=crop" alt="Gravel">
            </div>
            <div class="p-5 flex flex-col flex-grow">
              <div class="flex justify-between items-center mb-2">
                <span class="text-orange-500 text-xs font-semibold uppercase tracking-wider">Gravel</span>
                <span class="text-gray-400 text-xs font-medium">Size: M</span>
              </div>
              <h3 class="text-lg font-bold text-white mb-1 group-hover:text-orange-500 transition-colors">Canyon Grizl CF SL 8</h3>
              <div class="text-gray-400 text-xs font-medium flex items-center gap-1 mb-4">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" /></svg>
                Austin, TX
              </div>
              <div class="mt-auto flex justify-between items-end">
                <span class="text-2xl font-bold text-orange-500">$2,800</span>
              </div>
            </div>
          </div>

          <!-- Card 4 (Parts) -->
          <div class="col-span-1 group bg-gray-800 rounded-xl overflow-hidden border border-gray-700 hover:border-orange-500/50 transition-all flex flex-col relative">
            <div class="absolute top-4 right-4 z-10">
              <button class="w-8 h-8 rounded-full bg-gray-900/80 backdrop-blur-sm flex items-center justify-center text-gray-400 hover:text-orange-500 transition-all">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" /></svg>
              </button>
            </div>
            <div class="aspect-[4/3] w-full relative bg-gray-900 p-4 flex items-center justify-center overflow-hidden">
              <div class="absolute inset-0 bg-gradient-to-tr from-gray-800 to-transparent opacity-50"></div>
              <img class="w-4/5 h-auto object-contain drop-shadow-2xl group-hover:scale-110 transition-transform duration-500 z-10" src="https://images.unsplash.com/photo-1599482629446-10e20a00a6b5?w=600&h=400&fit=crop" alt="Wheelset">
              <div class="absolute top-4 left-4 z-10">
                <span class="bg-gray-800 text-gray-400 text-xs font-semibold px-2.5 py-1 rounded-md uppercase tracking-wider border border-gray-700">Parts</span>
              </div>
            </div>
            <div class="p-5 flex flex-col flex-grow">
              <div class="flex justify-between items-center mb-2">
                <span class="text-orange-500 text-xs font-semibold uppercase tracking-wider">Components</span>
              </div>
              <h3 class="text-lg font-bold text-white mb-1 group-hover:text-orange-500 transition-colors">Zipp 404 Firecrest Wheelset</h3>
              <div class="text-gray-400 text-xs font-medium flex items-center gap-1 mb-4">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" /></svg>
                Ships from CA
              </div>
              <div class="mt-auto flex justify-between items-end">
                <span class="text-2xl font-bold text-orange-500">$1,200</span>
              </div>
            </div>
          </div>

          <!-- Card 5 -->
          <div class="col-span-1 group bg-gray-800 rounded-xl overflow-hidden border border-gray-700 hover:border-orange-500/50 transition-all flex flex-col relative">
            <div class="absolute top-4 right-4 z-10">
              <button class="w-8 h-8 rounded-full bg-gray-900/80 backdrop-blur-sm flex items-center justify-center text-gray-400 hover:text-orange-500 transition-all">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" /></svg>
              </button>
            </div>
            <div class="aspect-[4/3] w-full relative bg-gray-900 overflow-hidden">
              <img class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-500" src="https://images.unsplash.com/photo-1485965120184-e224f723d1e1?w=600&h=400&fit=crop" alt="Vintage">
            </div>
            <div class="p-5 flex flex-col flex-grow">
              <div class="flex justify-between items-center mb-2">
                <span class="text-orange-500 text-xs font-semibold uppercase tracking-wider">Road / Vintage</span>
                <span class="text-gray-400 text-xs font-medium">Size: 56cm</span>
              </div>
              <h3 class="text-lg font-bold text-white mb-1 group-hover:text-orange-500 transition-colors">Colnago Master Olympic '94</h3>
              <div class="text-gray-400 text-xs font-medium flex items-center gap-1 mb-4">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" /><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" /></svg>
                Brooklyn, NY
              </div>
              <div class="mt-auto flex justify-between items-end">
                <span class="text-2xl font-bold text-orange-500">$3,100</span>
              </div>
            </div>
          </div>
        </div>

        <!-- Load More -->
        <div class="mt-12 flex justify-center">
          <button class="px-8 py-3 rounded-lg border border-gray-700 text-white hover:text-orange-500 hover:border-orange-500 transition-colors font-semibold text-sm bg-gray-800 hover:bg-gray-700">
            Load More Results
          </button>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import Main_nav from '@/components/Main_nav.vue'
// Note: If you don't have the '@' alias configured, change to: import Main_nav from '../components/Main_nav.vue'
</script>

<style scoped>
/* Custom scrollbar for the main page */
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

/* Custom scrollbar for the sidebar */
.custom-scrollbar::-webkit-scrollbar {
    width: 6px;
}
.custom-scrollbar::-webkit-scrollbar-track {
    background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
    background-color: #374151; /* gray-700 */
    border-radius: 20px;
}
</style>
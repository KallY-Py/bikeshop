<template>
  <div class="fixed inset-0 bg-black/50 z-50 flex items-center justify-center p-4">
    <div class="bg-surface-container-high border border-outline-variant rounded-xl p-6 max-w-2xl w-full max-h-[90vh] overflow-y-auto">
      <div class="flex justify-between items-center mb-6">
        <h2 class="font-headline-lg text-headline-lg text-on-surface">{{ listing?.id ? 'Edit Listing' : 'New Listing' }}</h2>
        <button @click="$emit('close')" class="text-on-surface-variant hover:text-on-surface transition-colors">
          <span class="material-symbols-outlined">close</span>
        </button>
      </div>
      
      <form @submit.prevent="handleSubmit" class="space-y-4">
        <!-- Title -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Title *</label>
          <input 
            v-model="form.title" 
            type="text" 
            required 
            placeholder="e.g., Trek Madone SLR 9 2024" 
            class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary focus:outline-none transition-colors" 
          />
        </div>
        
        <!-- Category -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Category *</label>
          <select 
            v-model="form.category_id" 
            required 
            class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary focus:outline-none transition-colors"
          >
            <option value="">Select category</option>
            <option v-for="cat in categories" :key="cat.id" :value="cat.id">{{ cat.name }} ({{ cat.type }})</option>
          </select>
        </div>
        
        <!-- Price -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Price (PHP) *</label>
          <input 
            v-model.number="form.price" 
            type="number" 
            required 
            min="0" 
            step="0.01" 
            placeholder="0.00" 
            class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary focus:outline-none transition-colors" 
          />
        </div>
        
        <!-- Condition -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Condition *</label>
          <select 
            v-model="form.condition_type" 
            required 
            class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary focus:outline-none transition-colors"
          >
            <option value="used">Used</option>
            <option value="brand_new">Brand New</option>
            <option value="refurbished">Refurbished</option>
          </select>
        </div>
        
        <!-- Location -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Location</label>
          <input 
            v-model="form.location" 
            type="text" 
            placeholder="City, State" 
            class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary focus:outline-none transition-colors" 
          />
        </div>
        
        <!-- Image Upload -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">
            Listing Image <span class="text-on-surface-variant text-xs">(JPEG or PNG only, max 5MB)</span>
          </label>
          
          <!-- Drop Zone -->
          <div 
            @click="$refs.fileInput.click()"
            @dragover.prevent="dragOver = true"
            @dragleave.prevent="dragOver = false"
            @drop.prevent="handleDrop"
            class="border-2 border-dashed rounded-lg p-6 text-center cursor-pointer transition-colors"
            :class="dragOver ? 'border-primary bg-primary/5' : 'border-outline-variant hover:border-primary/50'"
          >
            <input 
              ref="fileInput"
              type="file" 
              accept="image/jpeg,image/png"
              @change="handleFileSelect"
              class="hidden"
            />
            
            <!-- No file selected -->
            <div v-if="!imagePreview && !form.image_url">
              <span class="material-symbols-outlined text-4xl text-on-surface-variant mb-2">cloud_upload</span>
              <p class="text-on-surface-variant font-label-md mb-1">Click to upload or drag and drop</p>
              <p class="text-on-surface-variant text-xs">JPEG or PNG only • Max 5MB</p>
            </div>
            
            <!-- Preview -->
            <div v-else class="relative">
              <img :src="imagePreview || form.image_url" class="max-h-48 mx-auto rounded-lg object-cover" />
              <button 
                type="button"
                @click.stop="removeImage"
                class="absolute top-2 right-2 bg-black/60 hover:bg-black/80 text-white rounded-full p-1 transition-colors"
              >
                <span class="material-symbols-outlined text-lg">close</span>
              </button>
              <p v-if="selectedFile" class="text-on-surface-variant text-xs mt-2">{{ selectedFile.name }}</p>
              <p v-else-if="form.image_url" class="text-on-surface-variant text-xs mt-2">Current image</p>
            </div>
          </div>
          
          <!-- Upload progress -->
          <div v-if="uploading" class="mt-2 flex items-center gap-2 text-sm text-primary">
            <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-primary"></div>
            Uploading image...
          </div>
          
          <!-- Upload error -->
          <div v-if="uploadError" class="mt-2 text-sm text-error">
            {{ uploadError }}
          </div>
        </div>
        
        <!-- Description -->
        <div>
          <label class="block font-label-md text-label-md text-on-surface mb-2">Description</label>
          <textarea 
            v-model="form.description" 
            rows="4" 
            placeholder="Describe your bike, its condition, and any upgrades..." 
            class="w-full bg-surface border border-outline-variant rounded-lg py-2.5 px-4 text-on-surface focus:border-primary focus:outline-none transition-colors resize-none"
          ></textarea>
        </div>
        
        <!-- Buttons -->
        <div class="flex gap-3 justify-end pt-4 border-t border-outline-variant">
          <button 
            type="button" 
            @click="$emit('close')" 
            class="px-6 py-2.5 rounded-lg bg-surface-variant text-on-surface hover:bg-surface-bright transition-colors"
          >
            Cancel
          </button>
          <button 
            type="submit" 
            :disabled="saving || uploading" 
            class="px-6 py-2.5 rounded-lg bg-gradient-to-b from-[#FF8A00] to-[#FF6B00] text-black font-label-md hover:brightness-110 transition-all disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2"
          >
            <span v-if="uploading" class="animate-spin rounded-full h-4 w-4 border-b-2 border-black"></span>
            {{ uploading ? 'Uploading...' : saving ? 'Saving...' : (listing?.id ? 'Update Listing' : 'Post Listing') }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'ListingModal',
  props: {
    listing: { type: Object, default: null },
    categories: { type: Array, default: () => [] },
    saving: { type: Boolean, default: false }
  },
  emits: ['close', 'save'],
  data() {
    return {
      selectedFile: null,
      imagePreview: null,
      uploading: false,
      uploadError: null,
      dragOver: false,
      form: {
        title: '',
        category_id: '',
        price: 0,
        condition_type: 'used',
        location: '',
        description: '',
        image_url: ''
      }
    }
  },
  watch: {
    listing: {
      immediate: true,
      handler(val) {
        if (val) {
          this.form = {
            title: val.title || '',
            category_id: val.category_id || '',
            price: val.price || 0,
            condition_type: val.condition_type || 'used',
            location: val.location || '',
            description: val.description || '',
            image_url: val.image_url || ''
          }
          if (val.image_url) {
            this.imagePreview = val.image_url
          }
        }
      }
    }
  },
  methods: {
    handleFileSelect(event) {
      const file = event.target.files[0]
      this.processFile(file)
    },
    
    handleDrop(event) {
      this.dragOver = false
      const file = event.dataTransfer.files[0]
      this.processFile(file)
    },
    
    processFile(file) {
      if (!file) return
      
      this.uploadError = null
      
      // Validate file type - JPEG or PNG only
      const validTypes = ['image/jpeg', 'image/png', 'image/jpg']
      if (!validTypes.includes(file.type)) {
        this.uploadError = 'Please select a JPEG or PNG image only.'
        return
      }
      
      // Validate file size (max 5MB)
      if (file.size > 5 * 1024 * 1024) {
        this.uploadError = 'Image must be less than 5MB.'
        return
      }
      
      this.selectedFile = file
      
      // Show preview
      const reader = new FileReader()
      reader.onload = (e) => {
        this.imagePreview = e.target.result
      }
      reader.readAsDataURL(file)
      
      // Reset file input
      if (this.$refs.fileInput) {
        this.$refs.fileInput.value = ''
      }
    },
    
    removeImage() {
      this.selectedFile = null
      this.imagePreview = null
      this.form.image_url = ''
      this.uploadError = null
      // Reset file input
      if (this.$refs.fileInput) {
        this.$refs.fileInput.value = ''
      }
    },
    
    async uploadImage() {
      if (!this.selectedFile) return this.form.image_url || null
      
      this.uploading = true
      this.uploadError = null
      
      try {
        const formData = new FormData()
        formData.append('image', this.selectedFile)
        
        const response = await axios.post('http://localhost:3000/api/upload', formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          },
          timeout: 30000 // 30 second timeout
        })
        
        if (response.data && response.data.image_url) {
          console.log('Image uploaded successfully:', response.data.image_url)
          return response.data.image_url
        } else {
          throw new Error('No image URL returned from server')
        }
      } catch (err) {
        console.error('Image upload failed:', err)
        
        let errorMessage = 'Failed to upload image. '
        if (err.response) {
          // Server responded with error
          errorMessage += err.response.data?.error || err.response.statusText || 'Server error'
        } else if (err.request) {
          // Request made but no response
          errorMessage += 'No response from server. Please check your connection.'
        } else {
          // Request setup error
          errorMessage += err.message || 'Unknown error'
        }
        
        this.uploadError = errorMessage
        return null
      } finally {
        this.uploading = false
      }
    },
    
    async handleSubmit() {
      // Validate required fields
      if (!this.form.title.trim()) {
        alert('Please enter a title for your listing.')
        return
      }
      
      if (!this.form.category_id) {
        alert('Please select a category.')
        return
      }
      
      if (!this.form.price || this.form.price <= 0) {
        alert('Please enter a valid price.')
        return
      }
      
      // Upload image first if file selected
      if (this.selectedFile) {
        const imageUrl = await this.uploadImage()
        if (imageUrl) {
          this.form.image_url = imageUrl
        } else {
          // Don't submit if upload failed
          return
        }
      }
      
      // Require image for new listings (optional for edit)
      if (!this.listing?.id && !this.form.image_url) {
        alert('Please upload an image for your listing.')
        return
      }
      
      // Create a clean copy of the form data
      const submitData = {
        title: this.form.title.trim(),
        category_id: parseInt(this.form.category_id),
        price: parseFloat(this.form.price) || 0,
        condition_type: this.form.condition_type,
        location: this.form.location.trim(),
        description: this.form.description.trim(),
        image_url: this.form.image_url || ''
      }
      
      console.log('Submitting listing data:', submitData)
      this.$emit('save', submitData)
    }
  }
}
</script>

<style scoped>
/* Smooth scroll for modal */
.modal-content {
  scrollbar-width: thin;
  scrollbar-color: var(--outline-variant) transparent;
}

.modal-content::-webkit-scrollbar {
  width: 6px;
}

.modal-content::-webkit-scrollbar-track {
  background: transparent;
}

.modal-content::-webkit-scrollbar-thumb {
  background: var(--outline-variant);
  border-radius: 3px;
}
</style>
<template>
  <div class="organisms__navbar">
    <!-- Navigation Links -->
    <div class="flex items-center gap-6">
      <NuxtLink to="/" class="text-lg font-bold">
        Admin Panel
      </NuxtLink>
      <nav class="flex items-center gap-4">
        <NuxtLink
            to="/"
            class="nav-link"
            :class="{ 'nav-link-active': $route.name === 'index' }"
        >
          Dashboard
        </NuxtLink>
        <NuxtLink
            to="/users"
            class="nav-link"
            :class="{ 'nav-link-active': $route.name === 'users' }"
        >
          Users
        </NuxtLink>
      </nav>
    </div>

    <div class="flex-grow"></div>

    <!-- User menu when authenticated -->
    <div v-if="useUserStore().isAuthorized && user" class="flex items-center gap-4">
      <span class="text-sm text-gray-600 dark:text-gray-300">
        {{ user.name }}
      </span>
      <UDropdown
          :items="userMenuItems"
          :popper="{ placement: 'bottom-end' }"
      >
        <UAvatar
            :alt="user.name"
            size="sm"
            class="cursor-pointer"
        />
      </UDropdown>
    </div>

    <!-- Login button when not authenticated -->
    <NuxtLink to="/auth" v-else-if="authButtonVisible">
      <UButton>Login</UButton>
    </NuxtLink>

    <UButton v-else-if="!authButtonVisible" @click="fuckingLogout">Logout</UButton>
  </div>
</template>

<script setup lang="ts">


// const { user, isAuthenticated, logout } = useAuth()

const fuckingLogout = () => {
  useUserStore().logout();
}

const authButtonVisible = computed(() => {
  return useRoute().name !== 'auth' && !useUserStore().isAuthorized
})
//
// const userMenuItems = [
//   [{
//     label: 'Profile',
//     icon: 'i-heroicons-user-circle',
//     disabled: true // TODO: Implement profile page
//   }],
//   [{
//     label: 'Logout',
//     icon: 'i-heroicons-arrow-left-on-rectangle',
//     click: async () => {
//       await logout()
//     }
//   }]
// ]
</script>

<style scoped>

@reference "@/assets/css/main.css";

.organisms__navbar {
  @apply w-screen p-5 flex items-center gap-x-10 border-b border-gray-200 dark:border-gray-700
}

.nav-link {
  @apply px-3 py-2 text-sm font-medium text-gray-600 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100 transition-colors duration-200
}

.nav-link-active {
  @apply text-blue-600 dark:text-blue-400
}

</style>
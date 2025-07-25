<template>
  <div class="h-full flex items-center justify-center">
    <AuthForm
        ref="authFormRef"

    />
  </div>
</template>
<script setup lang="ts">
import {AuthForm} from '~/components/organisms/AuthForm'

// Disable auth middleware for this page and use auth layout
definePageMeta({
  middleware: [],
  layout: 'auth'
})

// const { login: authLogin, isLoading } = useAuth()
// const toast = useToast()
// const authFormRef = ref<InstanceType<typeof AuthForm>>()
//
// const handleLogin = async (login: string, password: string) => {
//   console.log('handleLogin called with:', { login, password })
//   try {
//     await authLogin({ login, password })
//     toast.add({
//       title: 'Success',
//       description: 'Logged in successfully',
//       color: 'green'
//     })
//   } catch (error: any) {
//     console.error('Login failed:', error)
//
//     // Show error in the form
//     if (authFormRef.value) {
//       authFormRef.value.setError(error.message || 'Login failed. Please check your credentials.')
//     }
//
//     // Also show toast for error
//     toast.add({
//       title: 'Login Failed',
//       description: error.message || 'Please check your credentials and try again',
//       color: 'red'
//     })
//   }
// }
</script>
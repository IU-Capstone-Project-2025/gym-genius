<template>
  <UCard class="auth__login-form">
    <template #header>
      Log In
    </template>

    <div class="flex flex-col gap-y-4">
      <UFormGroup
          label="Login"
          :error="errors.login"
          class="space-y-1"
      >
        <UInput
            class="w-full"
            placeholder="Enter your login"
            size="xl"
            autocomplete="off"
            v-model="login"
            :disabled="isLoading"
        />
      </UFormGroup>

      <UFormGroup
          label="Password"
          :error="errors.password"
          class="space-y-1"
      >
        <UInput
            class="w-full"
            placeholder="Enter your password"
            size="xl"
            type="password"
            autocomplete="new-password"
            v-model="password"
            :disabled="isLoading"
            @keyup.enter="handleLogin"
        />
      </UFormGroup>

      <UAlert
          v-if="generalError"
          color="red"
          variant="solid"
          :title="generalError"
          class="mb-4"
      />
    </div>

    <template #footer>
      <div class="flex justify-end">
        <UButton
            size="xl"
            color="primary"
            :loading="isLoading"
            :disabled="isLoading"
            @click="handleLogin"
        >
          {{ isLoading ? 'Logging in...' : 'Log in' }}
        </UButton>
      </div>
    </template>
  </UCard>
</template>

<script setup lang="ts">

const props = defineProps<{
  isLoading?: boolean
}>();

const emit = defineEmits<{
  (e: "submit", login: string, password: string): void
}>()

const login = ref("");
const password = ref("");
const errors = ref({
  login: '',
  password: ''
});
const generalError = ref('');

const validateForm = (): boolean => {
  // Clear previous errors
  errors.value.login = '';
  errors.value.password = '';
  generalError.value = '';

  return login.value && password.value;
};

const handleLogin = async () => {
  console.log('handleLogin in AuthForm called')
  generalError.value = '';

  if (!validateForm()) {
    generalError.value = "All fields must be filled";
    console.log('Validation failed');
    return;
  }

  console.log('Emitting submit event with:', login.value, password.value)
  emit('submit', login.value, password.value);

  const userStore = useUserStore();
  const toast = useToast()

  const login_ = async (login: string, password: string) => {
    userStore.login(login, password)
        .then(result => {
          toast.add({
            title: 'Success',
            description: 'Logged in successfully',
            color: 'green'
          })
          navigateTo('/')
        })
        .catch(error => {
          generalError.value = error.statusCode === 404 ? "Invalid credentials" : "Network error. Try again"
          toast.add({
            title: 'Login Failed',
            description: generalError.value,
            color: 'red'
          })
        })
  }

  await login_(login.value, password.value);

}

// Clear general error when user starts typing
watch([login, password], () => {
  generalError.value = '';
});

// Show error passed from parent
defineExpose({
  setError: (error: string) => {
    generalError.value = error;
  }
});

</script>

<style scoped>


.auth__login-form {
  @apply w-[500px]
}

</style>
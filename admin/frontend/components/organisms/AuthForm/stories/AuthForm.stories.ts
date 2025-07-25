// components/organisms/AuthForm/stories/AuthForm.stories.ts
import type {Meta, StoryObj} from '@storybook/vue3';
import AuthForm from "../ui/AuthForm.vue";

const meta = {
    title: "Organisms/AuthForm",
    component: AuthForm,
    tags: ['autodocs'],
    argTypes: {
        isLoading: {
            control: 'boolean',
            description: 'Loading state of the form'
        },
        submit: {
            action: 'submit',
            description: 'Event emitted when submitting the form with login and password'
        }
    },
    parameters: {
        componentSubtitle: 'Authentication form for admin login',
        layout: 'fullscreen'
    }
} satisfies Meta<typeof AuthForm>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
    args: {
        isLoading: false
    }
};

export const Loading: Story = {
    args: {
        isLoading: true
    }
};

// Interactive version that shows form behavior
export const Interactive: Story = {
    parameters: {
        docs: {
            description: {
                story: 'Interactive example showing form validation and submission'
            }
        }
    },
    render: (args) => ({
        components: {AuthForm},
        setup() {
            const isLoading = ref(args.isLoading || false);
            const authFormRef = ref(null);

            async function handleSubmit(login, password) {
                isLoading.value = true;
                args.submit(login, password);
                
                // Simulate API call
                await new Promise(resolve => setTimeout(resolve, 2000));
                
                // Simulate error for demo
                if (login === 'error') {
                    authFormRef.value?.setError('Invalid credentials');
                }
                
                isLoading.value = false;
            }

            return {isLoading, authFormRef, handleSubmit};
        },
        template: `
          <div class="h-screen bg-gray-100 p-4 flex items-center justify-center">
            <AuthForm
                ref="authFormRef"
                :is-loading="isLoading"
                @submit="handleSubmit"
            />
            <div class="fixed bottom-4 left-4 p-4 bg-white rounded shadow">
              <p class="text-sm font-semibold mb-2">Debug Info:</p>
              <p class="text-xs">Loading: {{ isLoading }}</p>
              <p class="text-xs mt-2">Try logging in with any credentials</p>
              <p class="text-xs">Use login "error" to simulate an error</p>
            </div>
          </div>
        `
    })
};

// Story that demonstrates form validation behavior
export const FormValidation: Story = {
    args: {
        isLoading: false
    },
    parameters: {
        docs: {
            description: {
                story: 'Demonstrates form validation with different scenarios'
            }
        }
    },
    render: (args) => ({
        components: {AuthForm},
        setup() {
            const submittedData = ref(null);
            const authFormRef = ref(null);

            function handleSubmit(login, password) {
                submittedData.value = {login, password, timestamp: new Date().toLocaleTimeString()};
                args.submit(login, password);

                // Reset after 3 seconds
                setTimeout(() => {
                    submittedData.value = null;
                }, 3000);
            }

            return {args, submittedData, authFormRef, handleSubmit};
        },
        template: `
          <div class="h-screen bg-gray-100 p-4 flex items-center justify-center">
            <AuthForm
                ref="authFormRef"
                :is-loading="args.isLoading"
                @submit="handleSubmit"
            />
            <div class="fixed bottom-4 left-4 p-4 bg-white rounded shadow">
              <p class="text-sm font-semibold mb-2">Form Submission:</p>
              <div v-if="submittedData" class="text-xs p-2 bg-green-100 rounded">
                <p class="font-semibold text-green-700">Form Submitted Successfully!</p>
                <p class="mt-1">Login: {{ submittedData.login }}</p>
                <p>Password: {{ '•'.repeat(submittedData.password.length) }}</p>
                <p>Time: {{ submittedData.timestamp }}</p>
              </div>
              <p v-else class="text-xs">
                Form will validate on blur and submit.<br>
                Min 3 chars for login, 6 for password.
              </p>
            </div>
          </div>
        `
    })
};

// Dark mode story
export const DarkMode: Story = {
    render: () => ({
        components: { AuthForm },
        setup() {
            const isLoading = ref(false);

            function handleSubmit(login, password) {
                console.log('Login submitted:', login, password);
            }

            return { isLoading, handleSubmit };
        },
        template: `
            <div style="min-height: 100vh; background-color: #1a1a1a; padding: 2rem;" class="dark flex items-center justify-center">
                <AuthForm 
                    :is-loading="isLoading"
                    @submit="handleSubmit"
                />
            </div>
        `
    })
};

// Error state story
export const WithError: Story = {
    render: () => ({
        components: { AuthForm },
        setup() {
            const authFormRef = ref(null);
            const isLoading = ref(false);

            onMounted(() => {
                // Show error after component mounts
                setTimeout(() => {
                    authFormRef.value?.setError('Invalid credentials. Please try again.');
                }, 500);
            });

            function handleSubmit(login, password) {
                console.log('Login submitted:', login, password);
            }

            return { authFormRef, isLoading, handleSubmit };
        },
        template: `
            <div class="h-screen bg-gray-100 p-4 flex items-center justify-center">
                <AuthForm 
                    ref="authFormRef"
                    :is-loading="isLoading"
                    @submit="handleSubmit"
                />
            </div>
        `
    })
};
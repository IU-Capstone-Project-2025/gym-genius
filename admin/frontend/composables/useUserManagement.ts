import {ref, computed} from 'vue'

export type UserStatus = 'active' | 'banned' | 'suspended'
export type SubscriptionPlan = 'free' | 'basic' | 'premium' | 'pro'
export type SubscriptionStatus = 'active' | 'expired' | 'cancelled' | 'pending'

export interface User {
    id: number
    name: string
    email: string
    phone?: string
    avatar?: string
    status: UserStatus
    registrationDate: Date
    lastActivity: Date
    subscription: {
        plan: SubscriptionPlan
        status: SubscriptionStatus
        startDate: Date
        endDate: Date
        autoRenew: boolean
    }
    stats: {
        totalWorkouts: number
        totalTimeSpent: number // in minutes
        averageSessionTime: number // in minutes
        favoriteExercises: string[]
        streak: number // current streak in days
        monthlyActivity: number[] // last 12 months
    }
}

export interface UserFilters {
    search: string
    status: UserStatus | 'all'
    subscriptionPlan: SubscriptionPlan | 'all'
    subscriptionStatus: SubscriptionStatus | 'all'
    registrationDateFrom?: Date
    registrationDateTo?: Date
}

export const useUserManagement = () => {
    const users = ref<User[]>([])
    const isLoading = ref(false)
    const selectedUser = ref<User | null>(null)
    const currentPage = ref(1)
    const pageSize = ref(10)
    const totalUsers = ref(0)

    const filters = ref<UserFilters>({
        search: '',
        status: 'all',
        subscriptionPlan: 'all',
        subscriptionStatus: 'all'
    })

    // Mock data generation
    const generateMockUsers = (count: number): User[] => {
        const names = ['John Doe', 'Jane Smith', 'Mike Johnson', 'Sarah Wilson', 'Chris Brown', 'Emily Davis', 'David Lee', 'Lisa Garcia', 'Tom Anderson', 'Anna Martinez', 'Kevin Taylor', 'Michelle White', 'Robert Thompson', 'Jennifer Lopez', 'Mark Wilson']
        const exercises = ['Push-ups', 'Squats', 'Deadlifts', 'Bench Press', 'Pull-ups', 'Lunges', 'Plank', 'Burpees', 'Rows', 'Overhead Press']
        const statuses: UserStatus[] = ['active', 'banned', 'suspended']
        const plans: SubscriptionPlan[] = ['free', 'basic', 'premium', 'pro']
        const subStatuses: SubscriptionStatus[] = ['active', 'expired', 'cancelled', 'pending']

        return Array.from({length: count}, (_, i) => {
            const registrationDate = new Date(Date.now() - Math.random() * 365 * 24 * 60 * 60 * 1000)
            const lastActivity = new Date(Date.now() - Math.random() * 7 * 24 * 60 * 60 * 1000)
            const subscriptionStart = new Date(registrationDate.getTime() + Math.random() * 30 * 24 * 60 * 60 * 1000)
            const subscriptionEnd = new Date(subscriptionStart.getTime() + 30 * 24 * 60 * 60 * 1000)

            return {
                id: i + 1,
                name: names[i % names.length],
                email: `${names[i % names.length].toLowerCase().replace(' ', '.')}@email.com`,
                phone: Math.random() > 0.5 ? `+1${Math.floor(Math.random() * 9000000000 + 1000000000)}` : undefined,
                avatar: `https://api.dicebear.com/7.x/avataaars/svg?seed=${i}`,
                status: statuses[Math.floor(Math.random() * statuses.length)],
                registrationDate,
                lastActivity,
                subscription: {
                    plan: plans[Math.floor(Math.random() * plans.length)],
                    status: subStatuses[Math.floor(Math.random() * subStatuses.length)],
                    startDate: subscriptionStart,
                    endDate: subscriptionEnd,
                    autoRenew: Math.random() > 0.5
                },
                stats: {
                    totalWorkouts: Math.floor(Math.random() * 200) + 10,
                    totalTimeSpent: Math.floor(Math.random() * 5000) + 100,
                    averageSessionTime: Math.floor(Math.random() * 90) + 30,
                    favoriteExercises: exercises.slice(0, Math.floor(Math.random() * 5) + 1),
                    streak: Math.floor(Math.random() * 30),
                    monthlyActivity: Array.from({length: 12}, () => Math.floor(Math.random() * 25) + 5)
                }
            }
        })
    }

    // Initialize mock data
    const initializeMockData = () => {
        users.value = generateMockUsers(50)
        totalUsers.value = users.value.length
    }

    const filteredUsers = computed(() => {
        let filtered = [...users.value]

        // Search filter
        if (filters.value.search) {
            const search = filters.value.search.toLowerCase()
            filtered = filtered.filter(user =>
                user.name.toLowerCase().includes(search) ||
                user.email.toLowerCase().includes(search) ||
                user.phone?.includes(search)
            )
        }

        // Status filter
        if (filters.value.status !== 'all') {
            filtered = filtered.filter(user => user.status === filters.value.status)
        }

        // Subscription plan filter
        if (filters.value.subscriptionPlan !== 'all') {
            filtered = filtered.filter(user => user.subscription.plan === filters.value.subscriptionPlan)
        }

        // Subscription status filter
        if (filters.value.subscriptionStatus !== 'all') {
            filtered = filtered.filter(user => user.subscription.status === filters.value.subscriptionStatus)
        }

        // Date range filters
        if (filters.value.registrationDateFrom) {
            filtered = filtered.filter(user => user.registrationDate >= filters.value.registrationDateFrom!)
        }
        if (filters.value.registrationDateTo) {
            filtered = filtered.filter(user => user.registrationDate <= filters.value.registrationDateTo!)
        }

        return filtered
    })

    const paginatedUsers = computed(() => {
        const start = (currentPage.value - 1) * pageSize.value
        const end = start + pageSize.value
        return filteredUsers.value.slice(start, end)
    })

    const totalPages = computed(() => {
        return Math.ceil(filteredUsers.value.length / pageSize.value)
    })

    const fetchUsers = async () => {
        isLoading.value = true
        await new Promise(resolve => setTimeout(resolve, 500));
        let response = await $fetch('https://api.говно.site/users', {
            credentials: 'include'
        });
        console.log('USERS:', response);
        isLoading.value = false
    }

    const getUserById = (id: number): User | undefined => {
        return users.value.find(user => user.id === id)
    }

    const updateUserStatus = async (userId: number, status: UserStatus) => {
        const user = users.value.find(u => u.id === userId)
        if (user) {
            user.status = status
            await new Promise(resolve => setTimeout(resolve, 300))
        }
    }

    const updateUserSubscription = async (userId: number, subscription: Partial<User['subscription']>) => {
        const user = users.value.find(u => u.id === userId)
        if (user) {
            user.subscription = {...user.subscription, ...subscription}
            await new Promise(resolve => setTimeout(resolve, 300))
        }
    }

    const deleteUser = async (userId: number) => {
        const index = users.value.findIndex(u => u.id === userId)
        if (index > -1) {
            users.value.splice(index, 1)
            totalUsers.value = users.value.length
            await new Promise(resolve => setTimeout(resolve, 300))
        }
    }

    const clearFilters = () => {
        filters.value = {
            search: '',
            status: 'all',
            subscriptionPlan: 'all',
            subscriptionStatus: 'all'
        }
        currentPage.value = 1
    }

    // Initialize data
    initializeMockData()

    return {
        users: paginatedUsers,
        allUsers: users,
        isLoading,
        selectedUser,
        currentPage,
        pageSize,
        totalUsers: computed(() => filteredUsers.value.length),
        totalPages,
        filters,
        fetchUsers,
        getUserById,
        updateUserStatus,
        updateUserSubscription,
        deleteUser,
        clearFilters
    }
}
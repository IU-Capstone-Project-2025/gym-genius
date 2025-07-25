import {ref, computed} from 'vue'
import type {TimeInterval} from '~/components/molecules/TimeIntervalSelector.vue'
import {useUsersStats} from "~/composables/useUsersStats";

interface ActivityData {
    labels: string[]
    values: number[]
}

export const useUserActivity = () => {
    const selectedInterval = ref<TimeInterval>('1d')
    const isLoading = ref(false)
    
    const {fetchUsersStats, usersStats} = useUsersStats()

    // Transform usersStats data to ActivityData format
    const activityData = computed((): ActivityData => {
        if (!usersStats.value || usersStats.value.length === 0) {
            return { labels: [], values: [] }
        }

        const labels = usersStats.value.map(stat => {
            const date = new Date(stat.start_date)
            return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit' })
        })

        const values = usersStats.value.map(stat => stat.count)

        return { labels, values }
    })

    const fetchActivityData = async () => {
        isLoading.value = true
        await fetchUsersStats(selectedInterval.value);
        isLoading.value = false
    }

    return {
        selectedInterval,
        activityData,
        isLoading,
        fetchActivityData
    }
}
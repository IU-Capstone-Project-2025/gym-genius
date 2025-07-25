import {ref, type Ref} from 'vue'
import {useUserStore} from '~/stores/user'

type UsersStats = {
    start_date: string, // Date like
    end_date: string,
    count: number
}

export const useUsersStats = () => {
    const usersStats: Ref<UsersStats[]> = ref([]);
    const userStore = useUserStore();
    const startDate = '2023-06-01T00:00:00Z';
    const endDate = '2025-06-10T00:00:00Z'


    const fetchUsersStats = async (step: string = '1d') => {
        usersStats.value = await $fetch('https://api.говно.site/statistics/active-users', {
            query: {
                start_date: startDate,
                end_date: endDate,
                step: step
            },
            credentials: 'include'
        })

        console.log(usersStats.value);
    }

    return {
        fetchUsersStats,
        usersStats
    }
}
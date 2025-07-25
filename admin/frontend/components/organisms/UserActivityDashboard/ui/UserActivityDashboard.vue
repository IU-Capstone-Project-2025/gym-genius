<template>
  <UCard>
    <template #header>
      <div class="flex items-center justify-between">
        <h2 class="text-xl font-semibold">User Activity</h2>
        <TimeIntervalSelector
            v-model:selected-interval="selectedInterval"
        />
      </div>
    </template>


    <div class="relative h-96">
      <div v-if="isLoading"
           class="absolute inset-0 flex items-center justify-center bg-white/50 dark:bg-gray-900/50 z-10">
        <UIcon name="i-heroicons-arrow-path" class="animate-spin h-8 w-8 text-primary-500"/>
      </div>
      <ActivityChart
          :data="activityData"
          title=""
      />
    </div>

    <template #footer>
      <div class="flex items-center justify-between text-sm text-gray-500 dark:text-gray-400">
        <div class="flex items-center gap-2">
          <UIcon name="i-heroicons-users" class="h-4 w-4"/>
          <span>Total Active Users: {{ totalUsers }}</span>
        </div>
        <div class="flex items-center gap-2">
          <UIcon name="i-heroicons-arrow-trending-up" class="h-4 w-4"/>
          <span>Average: {{ averageUsers }}</span>
        </div>
      </div>
    </template>
  </UCard>
</template>

<script setup lang="ts">
import {computed, watch} from 'vue'

import {useUserActivity} from '~/composables/useUserActivity'
import {ActivityChart} from "~/components/molecules/ActivityChart";
import {TimeIntervalSelector} from "~/components/molecules/TimeIntervalSelector";


const {selectedInterval, activityData, isLoading, fetchActivityData} = useUserActivity()



const totalUsers = computed(() =>
    activityData.value.values.reduce((sum, val) => sum + val, 0)
)

const averageUsers = computed(() =>
    Math.round(totalUsers.value / activityData.value.values.length)
)

// Fetch data when interval changes
watch(selectedInterval, () => {
  fetchActivityData()
})

// Initial fetch
fetchActivityData()
</script>
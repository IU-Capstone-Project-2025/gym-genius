<template>
  <UButtonGroup size="sm" orientation="horizontal">
    <TimeIntervalButton
        v-for="interval in intervals"
        :key="interval.value"
        :label="interval.label"
        :is-active="selectedInterval === interval.value"
        @select="$emit('update:selectedInterval', interval.value)"
    />
  </UButtonGroup>
</template>

<script setup lang="ts">

import {TimeIntervalButton} from "~/components/atoms/TimeIntervalButton";

export type TimeInterval = '1h' | '1d' | '7d' | '30d'

interface TimeIntervalOption {
  label: string
  value: TimeInterval
}

interface Props {
  selectedInterval: TimeInterval
}

defineProps<Props>()
defineEmits<{
  'update:selectedInterval': [value: TimeInterval]
}>()

const intervals: TimeIntervalOption[] = [
  {label: '1 Hour', value: '1h'},
  {label: '1 Day', value: '1d'},
  {label: '7 Days', value: '7d'},
  {label: '30 Days', value: '30d'}
]
</script>
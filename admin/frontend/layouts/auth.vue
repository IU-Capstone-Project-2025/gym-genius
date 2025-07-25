<template>
  <div class="h-screen flex flex-col">
    <Navbar />
    <div class="flex-grow">
      <NuxtPage />
    </div>
  </div>
</template>

<script setup lang="ts">
import { Navbar } from "~/components/organisms/Navbar"
</script>
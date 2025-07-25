import tailwindcss from "@tailwindcss/vite";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    compatibilityDate: '2025-06-25',
    modules: ["@nuxt/ui", "@nuxtjs/storybook", "@pinia/nuxt"],
    css: ['@/assets/css/main.css'],
    vite: {
        plugins: [
            tailwindcss(),
        ],
    },
    ui: {
        fonts: false
    },
    storybook: {
        // Options
        host: 'http://localhost',
        port: 6006,
    },
})
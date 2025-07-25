export default defineNuxtRouteMiddleware(async (to, from) => {
    // Skip auth check for the auth page itself
    if (to.path === '/auth') {
        return
    }

    const {isAuthorized} = useUserStore()

    if (!isAuthorized) {
        // User is not authenticated, redirect to auth page
        // return navigateTo('/auth')
    }
})
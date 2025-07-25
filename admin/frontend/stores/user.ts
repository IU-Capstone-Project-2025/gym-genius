interface UserStore {
    login: (username: string, password: string) => Promise<void>,
    isAuthorized: ComputedRef<boolean>,
    logout: () => void
}

export const useUserStore = defineStore('userStore', (): UserStore => {
    const apiURI = 'https://api.говно.site/auth_admin'

    function login(username: string, password: string) {
        return $fetch(apiURI, {
            method: 'POST',
            body: {
                login: username,
                password: password
            },
            credentials: 'include'

        })
    }

    function logout() {
        navigateTo('/auth');
    }

    const isAuthorized = computed(() => {
        // console.log('TOKEN: ', token.value);
        // return !!token.value;
        return true;
    })

    return {
        login,
        logout,
        isAuthorized
    }
})
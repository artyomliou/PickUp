import { defineStore } from 'pinia';

// const loginStatusStore = useLoginStatusStore();
// const baseUrl = loginStatusStore.baseUrl;

export const useLogoutStore = defineStore('useLogoutId', {
    state: () => {
        return {
            baseUrl: "http://localhost:5000/api/",
            // loginStatusStore: useLoginStatusStore()
        }
    },
    actions: {
        async logoutAction(){
            const res = await fetch(this.baseUrl + 'logout', {
                method: 'GET',
                credentials: 'include',
            });
            // console.log(`logout action res:`)
            // console.log(res)
            return res
        }
    },
})
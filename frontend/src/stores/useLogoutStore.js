import { defineStore } from 'pinia';
export const useLogoutStore = defineStore('useLogoutId', {
    state: () => {
        return {
            baseUrl: "http://localhost:5000/api/",
            // loginStatusStore: useLoginStatusStore()
        }
    },
    actions: {
        async logoutAction(){
            try {
                const res = await fetch(this.baseUrl + 'logout', {
                    method: 'GET',
                    credentials: 'include',
                });
                return res    
            }
            catch(err){
                console.error(err)
            }
        }
    },
})
import { defineStore } from 'pinia';

// const loginStatusStore = useLoginStatusStore();
// const baseUrl = loginStatusStore.baseUrl;

export const useLoginStore = defineStore('useLoginId', {
    state: () => {
        return {
            baseUrl: "http://localhost:5000/api/",
        }
    },
    actions: {
        async isLoginStatus(){
            const url = this.baseUrl + 'is-logged-in'
            console.log(url)
            const response = await fetch(url, {
                credentials: 'include',
            });
            const apiStatus = await response.json();
            return !!apiStatus.isloggedin
        },
        async loginAction(email, password){
            try {
                const url = this.baseUrl + 'login'
                // console.log(url)
                const res = await fetch(url, {
                    method: 'POST',
                    credentials: 'include',
                    body: JSON.stringify({
                        email: email,
                        password: password
                    }),
                    headers: {
                        "content-type": "application/json"
                    },
                });
                return res
            }
            catch(error) {
                console.log(error)
            }
        }
    },
})
import { defineStore } from 'pinia';
export const useLoginStatusStore = defineStore('useLoginStatusId', {
    state: () => {
        return {
            baseUrl: "http://localhost:5000/api/",
            isLogin: false,
        }
    },
    actions: {
        async isLoginStatus(){
            const url = this.baseUrl + 'is-logged-in'
            // console.log(url)
            const response = await fetch(url, {
                credentials: 'include',
            });
            const apiStatus = await response.json();
            this.isLogin = !!apiStatus.isloggedin
            return !!apiStatus.isloggedin
        },
    },
})
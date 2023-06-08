import { defineStore } from 'pinia'

export const useLoginStore = defineStore('login', {
    state() {
        return {
            isLoggedIn: false
        }
    },
    actions: {
        updateLoginStatus(trueOrFalse = false) {
            this.isLoggedIn = trueOrFalse
        }
    }
})

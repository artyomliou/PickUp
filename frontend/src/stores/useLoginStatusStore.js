import { defineStore } from 'pinia';
export const useLoginStatusStore = defineStore('useLoginStatusId', {
    state: () => {
        return {
            isLogin: false,
            check: 0,
        }
    },
})
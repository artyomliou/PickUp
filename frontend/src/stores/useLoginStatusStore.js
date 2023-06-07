import { defineStore } from 'pinia';
import { isLoginStatus } from '../api';
export const useLoginStatusStore = defineStore('useLoginStatusId', {
    state: () => {
        return {
            isLogin: false,
        }
    },
    actions: {
        isLoginStatus
    },
})
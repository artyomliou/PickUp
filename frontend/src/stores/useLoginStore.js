import { defineStore } from 'pinia';
import { loginAction } from '../api/index';

// const loginStatusStore = useLoginStatusStore();
// const baseUrl = loginStatusStore.baseUrl;

export const useLoginStore = defineStore('useLoginId', {
    state: () => {
        return {
        }
    },
    actions: {
        loginAction
    },
})
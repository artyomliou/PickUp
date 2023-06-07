import { defineStore } from 'pinia';
import { logoutAction } from '../api';

export const useLogoutStore = defineStore('useLogoutId', {
    state: () => {
        return {}
    },
    actions: {
        logoutAction
    },
})
import { defineStore } from 'pinia'
import { isLoggedIn } from '../api'

// example https://pinia.vuejs.org/core-concepts/#option-stores

// pinia store
export const useLoginCheckStore = defineStore(
    'Login',
    {
        state: () => ({
            // as a data store
            isLogined: false
        }),
        getters: {
            // as the computed
        },
        actions: {
            async checkLoginStatus() {
                try {
                    this.isLogined = await isLoggedIn()
                } catch (error) {
                    console.error(error)
                }
            }
        }
    },
    true
)

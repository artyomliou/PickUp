import { defineStore } from 'pinia'
import { isLoggedIn } from '../api'

// example https://pinia.vuejs.org/core-concepts/#option-stores

// pinia store
export const useLoginStatusStore = defineStore(
    'Login',
    {
        state: () => ({
            isLogined: false
        }),
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

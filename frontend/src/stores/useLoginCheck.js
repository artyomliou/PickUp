import { defineStore } from 'pinia'

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
            // as the methods
            async fetchLoginApi() {
                try {
                    const response = await fetch('http://localhost:5000/api/is-logged-in', {
                        method: 'GET',
                        credentials: 'include'
                    })
                    const apiData = await response.json()
                    ;(this.isLogined = apiData['is-logged-in']),
                        console.log(
                            'action: ' + this.isLogined,
                            'apiData is-loggend-in: ' + apiData['is-logged-in']
                        )
                } catch (error) {
                    console.error(error)
                }
            }
        }
    },
    true
)

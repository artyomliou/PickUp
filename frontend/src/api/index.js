const baseUrl = 'http://localhost:5000'

function withBaseUrl(url) {
    return `${baseUrl}/api/${url}`
}

export async function isLoginStatus() {
    const url = withBaseUrl('is-logged-in')
    // console.log(url)
    const response = await fetch(url, {
        credentials: 'include'
    })
    const apiStatus = await response.json()
    // console.log(`apiStatus: ${!!apiStatus['is-logged-in']}`)
    return !!apiStatus['is-logged-in']
}

export async function loginAction(email, password) {
    const res = await fetch(withBaseUrl('login'), {
        method: 'POST',
        credentials: 'include',
        body: JSON.stringify({
            email: email,
            password: password
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    console.log(res)
    return res
}

export async function logoutAction() {
    const res = await fetch(withBaseUrl('logout'), {
        method: 'GET',
        credentials: 'include'
    })
    return res.ok
}

// Reactivity test
export function test() {
    const user = ''
    console.count(`test ${user}`)
}

/**
 * @returns {Promise<Store[]>}
 */
export async function fetchStoreList() {
    const res = await fetch(withBaseUrl('stores'))
    const json = await res.json()
    return json.stores
}

/**
 * @returns {Promise<Store>}
 */
export async function fetchStore(storeId) {
    const res = await fetch(withBaseUrl(`stores/${storeId}`))
    const json = await res.json()
    return json.store
}

/**
 * @returns {Promise<Product>}
 */
export async function fetchProduct(storeId, productId) {
    const res = await fetch(withBaseUrl(`stores/${storeId}/product/${productId}`))
    const json = await res.json()
    return json.product
}

/**
 * @returns {Promise<CartItem[]>}
 */
export async function fetchCartItemList(storeId) {
    const res = await fetch(withBaseUrl(`store/${storeId}/cart/items`), {
        credentials: 'include'
    })
    const json = await res.json()
    return json.items
}

/**
 *
 * @param {number} storeId
 * @param {number} productId
 * @param {number} amount
 * @param {SelectAnswer[]} selectAnswers
 * @param {CustomAnswer[]} customAnswers
 * @returns {Promise<boolean>}
 */
export async function createCartItem(storeId, productId, amount, selectAnswers, customAnswers) {
    const res = await fetch(withBaseUrl(`store/${storeId}/cart/items`), {
        credentials: 'include',
        method: 'POST',
        body: JSON.stringify({
            productId: Number.parseInt(productId, 10),
            amount: Number.parseInt(amount, 10),
            selectAnswers,
            customAnswers
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    if (!res.ok) {
        console.error(await res.json())
    }
    return res.ok
}

/**
 * @returns {Promise<Order>}
 */
export async function createOrder(storeId) {
    const res = await fetch(withBaseUrl(`orders`), {
        method: 'POST',
        credentials: 'include',
        body: JSON.stringify({
            storeId: Number.parseInt(storeId, 10)
        }),
        headers: {
            'Content-Type': 'application/json'
        }
    })
    const json = await res.json()
    return json.order
}

/**
 * @returns {Promise<Order>}
 */
export async function fetchOrder(orderId) {
    const res = await fetch(withBaseUrl(`orders/${orderId}`), {
        credentials: 'include'
    })
    const json = await res.json()
    return json.order
}

const baseurl = 'http://localhost:5000'

function withBaseurl(url) {
    return `${baseurl}${url}`
}

export async function isLoggedIn() {
    const response = await fetch(withBaseurl('/api/is-logged-in'), {
        credentials: 'include'
    })
    const apiData = await response.json()
    return !!apiData['is-logged-in']
}

export async function login(email, password) {
    return await fetch(withBaseurl('/api/login'), {
        method: 'POST',
        credentials: 'include',
        body: JSON.stringify({
            email: email,
            password: password
        })
    })
}

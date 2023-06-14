import './model';

const baseUrl = "http://localhost:5000";

function withBaseUrl(url){
    return `${baseUrl}/api/${url}`
};


export async function isLoginStatus(){
    const url = withBaseUrl('is-logged-in')
    // console.log(url)
    const response = await fetch(url, {
        credentials: 'include',
    });
    const apiStatus = await response.json();
    // console.log(`apiStatus: ${!!apiStatus['is-logged-in']}`)
    return !!apiStatus.isLoggedIn
}


export async function loginAction(email, password){
    const res = await fetch(withBaseUrl('login'), {
        method: 'POST',
        credentials: 'include',
        body: JSON.stringify({
            email: email,
            password: password
        }),
        headers: {
            "content-type": "application/json",
        }
    });
    console.log(!!res)
    return res.ok
}

export async function logoutAction(){
    const res = await fetch(withBaseUrl('logout'), {
        method: 'GET',
        credentials: 'include',
    });
    // console.log('--logoutaction...')
    // console.log(res)
    return res.ok
}

/**
 * @returns {Store[]}
 */
export async function storesListApi(){
    const res = await fetch(withBaseUrl('stores'), {
        method: 'GET',
        credentials: 'include',
    });
    const data = await res.json();
    // console.log(data.stores)
    return data.stores
}


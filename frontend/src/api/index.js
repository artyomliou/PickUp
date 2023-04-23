const baseUrl = "http://localhost:5000";

function withBaseUrl(url){
    return `${baseUrl}/api/${url}`
};

export async function isLoginStatus(){
    const url = withBaseUrl('is-logged-in')
    console.log(url)
    const response = await fetch(url, {
        credentials: 'include',
    });
    const apiStatus = await response.json();
    return !!apiStatus['is-logged-in']
}


export async function loginAction(email, password){
    const res = await fetch(withBaseUrl('login'), {
        method: 'POST',
        credentials: 'include',
        body: JSON.stringify({
            email: email,
            password: password
        })
    });
    console.log(res)
    return res
}

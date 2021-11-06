import React, {SyntheticEvent, useState} from 'react';
import { Redirect } from 'react-router-dom';
import googleLogo from "../assets/google_logo.png"
import vkLogo from "../assets/vk_logo.png"

const googleLogin = () => {
    window.open("http://localhost:8000/users/oauth/google", "_self");
}

const vkLogin = () => {
    window.open("http://localhost:8000/users/oauth/vk", "_self");
}

const Login = () => {

    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [redirect, setRedirect] = useState(false)

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault()

        await fetch('http://localhost:8000/users/signin', {
            method: 'POST',
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify({
                email,
                password
            })
        })
        setRedirect(true)
    }

    if (redirect)
    {
        return <Redirect to="/home"/>
    }

    return (
        <div className="loginPage">
                <h1 className="h3 mb-3 fw-normal">Please sign in via third party using OAuth</h1>
                <div className="oauthWrapper">
                <div className="oauthComponentContainer" onClick={googleLogin}>
                    <img src={googleLogo} alt="Google Icon" />
                    <p>Login With Google</p>
                </div>
                    <div className="oauthComponentContainer" onClick={vkLogin}>
                        <img src={vkLogo} alt="VK Icon" />
                        <p>Login With VK</p>
                    </div>
                </div>
            <h1 className="h3 mb-3 fw-normal">Or using local credentials</h1>
            <form onSubmit={submit}>
                <input type="email" className="form-control" placeholder="name@example.com"
                       onChange = {e => setEmail(e.target.value)}
                />
                <input type="password" className="form-control" placeholder="Password"
                       onChange = {e => setPassword(e.target.value)}
                />
                <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
            </form>
        </div>
    );
};

export default Login;
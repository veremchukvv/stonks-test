import React from 'react';
import googleLogo from "../assets/google_logo.png"
import vkLogo from "../assets/vk_logo.png"

const googleLogin = () => {
    window.open("https://o-auth-video-backend.herokuapp.com/auth/google", "_self");
}

const vkLogin = () => {
    window.open("https://o-auth-video-backend.herokuapp.com/auth/vk", "_self");
}

const Login = () => {
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
            <form>
                <input type="email" className="form-control" placeholder="name@example.com" />
                <input type="password" className="form-control" placeholder="Password" />
                <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
            </form>
        </div>
    );
};

export default Login;
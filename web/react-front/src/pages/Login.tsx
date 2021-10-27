import React from 'react';

const Login = () => {
    return (
        <div>
            <form>
                <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
                <input type="email" className="form-control" placeholder="name@example.com" />
                <input type="password" className="form-control" placeholder="Password" />
                <button className="w-100 btn btn-lg btn-primary" type="submit">Sign in</button>
            </form>
        </div>
    );
};

export default Login;
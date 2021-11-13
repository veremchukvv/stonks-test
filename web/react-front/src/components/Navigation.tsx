import React, {useContext} from 'react';
import {Link} from 'react-router-dom';
import {AuthContext} from "../context/authContext";

const Navigation = (props: {name: string, setName: (name: string) => void }) => {
    const auth = useContext(AuthContext)

    const logout = async () => {
        await fetch('http://localhost:8000/users/signout', {
            method: 'POST',
            headers: {'Content-Type':'application/json'},
            credentials: 'include',
            })
        props.setName('')
        }

    let menu

    if (!auth.isAuthenticated) {
        menu = (
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item active">
                    <Link to="/login" className="nav-link">Login</Link>
                </li>
                <li className="nav-item active">
                    <Link to="/register" className="nav-link">Register</Link>
                </li>
            </ul>
        )
    } else {
        menu = (
        <ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item active">
                <Link to="/profile" className="nav-link">{auth.userName}</Link>
            </li>
            <li className="nav-item active">
                <Link to="/" className="nav-link" onClick={logout}>Logout</Link>
            </li>
        </ul>
        )
    }

    return (
        <div>
            <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
                <div className="container-fluid">
                    <Link to="/" className="navbar-brand">Home</Link>
                    <div>
                        {menu}
                    </div>
                </div>
            </nav>
        </div>
    );
};

export default Navigation;
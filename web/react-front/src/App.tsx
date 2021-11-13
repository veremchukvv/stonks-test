import React, {useEffect, useState} from 'react';
import './App.css';
import Home from "./pages/Home"
import Login from "./pages/Login"
import Register from "./pages/Register"
import Profile from "./pages/Profile"
import Navigation from "./components/Navigation"
import {BrowserRouter, Route} from "react-router-dom";
import {AuthContext} from "./context/authContext";

function App() {
    const [name, setName] = useState('')
    const isAuthenticated = !!name
    const userName = name

    useEffect(() => {
        (
            async () => {
                const response = await fetch('http://localhost:8000/users/user', {
                    headers: {'Content-Type':'application/json'},
                    credentials: 'include',
                })
                const content = await response.json()

                setName(content.name)
            }
        )()
    }, [name]
        )

  return (
      <AuthContext.Provider value={{userName, isAuthenticated}}>
    <div className="App">
        <BrowserRouter>
            <Navigation name={name} setName={setName} />
      <main className="form-signin">
            <Route path="/" exact component={Home}/>
            <Route path="/login" component={() => <Login setName={setName}/>}/>
            <Route path="/register" component={Register}/>
            <Route path="/profile" component={Profile}/>
      </main>
        </BrowserRouter>
    </div>
      </AuthContext.Provider>
  );
}

export default App;

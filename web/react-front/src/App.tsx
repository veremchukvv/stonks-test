import React from 'react';
import './App.css';
import Home from "./pages/Home"
import Login from "./pages/Login"
import Register from "./pages/Register"
import Navigation from "./components/Navigation"
import {BrowserRouter, Route} from "react-router-dom";

function App() {
  return (
    <div className="App">
        <BrowserRouter>
            <Navigation />
      <main className="form-signin">
            <Route path="/" exact component={Home}/>
            <Route path="/login" component={Login}/>
            <Route path="/register" component={Register}/>

      </main>
        </BrowserRouter>
    </div>
  );
}

export default App;

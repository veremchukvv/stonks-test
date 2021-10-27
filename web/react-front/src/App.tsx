import React from 'react';
import './App.css';
import Login from "./pages/Login"
import Navigation from "./components/Navigation"

function App() {
  return (
    <div className="App">
        <Navigation />
      <main className="form-signin">
        <Login />
      </main>
    </div>
  );
}

export default App;

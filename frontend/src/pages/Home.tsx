import React from "react"
import Login from "../components/Login"
import logo from '../logo.svg';

const Home = () => {
    return (
    <div className="App">
        <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <h1 className="App-name">
        imageinit
        </h1>
        </header>
        <Login/>
    </div>
    )
}

export default Home;
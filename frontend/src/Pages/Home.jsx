import React from 'react'
import home from "../images/home.jpg"
import "../Styles/Home.css"
import { Link } from 'react-router-dom';

const Home = () => {
  return (
    <div className="home">
        <img src={home} alt="" />
    </div>
);
}

export default Home

import React, { useState } from 'react';
import '../Styles/Navbar.css';
import { Link } from 'react-router-dom';

const Navbar = () => {
    const [isOpen, setIsOpen] = useState(false);
    const toggleMenu = () => setIsOpen(!isOpen);

    return (
        <header className={`header ${isOpen ? "open" : ""}`} id='nav-menu'>
            <Link to="/" className="logo">Geo<span> Data</span></Link>
            <nav className={`navbar ${isOpen ? "showMenu" : ""}`}>
                <Link to="/" className="nav-link home">Home</Link>
                <Link to="/login" className="nav-link login">Login</Link>

                <div className="social-media-in">
                    <Link to="/" className="nav-link home">Home</Link>
                </div>
            </nav>
            <div className="social-media">
                <Link to="/" className="nav-link home">Home</Link>
            </div>
            <i className='bx bx-menu' id="menu-icon" onClick={toggleMenu}></i>
        </header>
    );
};

export default Navbar;

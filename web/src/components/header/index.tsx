import { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';
import './style.css';

export default function Header() {
  const location = useLocation();
  const [menuOpen, setMenuOpen] = useState(false);

  const isActive = (path: string) => {
    if (path === '/' && location.pathname === '/') return true;
    if (path === '/privacy' && location.pathname === '/privacy') return true;
    if (path === '/contact' && location.pathname === '/contact') return true;
    if (path === '/feedback' && location.pathname === '/feedback') return true;
    if (path === '/news' && location.pathname === '/news') return true;
    if (path === '/' && location.pathname.startsWith('/games/')) return true;
    return false;
  };

  const closeMenu = () => setMenuOpen(false);

  return (
    <header className="app-header">
      <div className="inner">
        <h1 className="logo" />

        <nav className="desktop-nav">
          <ul>
            <li>
              <Link to="/" className={isActive('/') ? 'active' : ''}>
                GAMES
              </Link>
            </li>
            <li>
              <Link to="/privacy" className={isActive('/privacy') ? 'active' : ''}>
                PRIVACY POLICY
              </Link>
            </li>
            <li>
              <Link to="/contact" className={isActive('/contact') ? 'active' : ''}>
                CONTACT
              </Link>
            </li>
            <li>
              <Link to="/feedback" className={isActive('/feedback') ? 'active' : ''}>
                FEEDBACK
              </Link>
            </li>
          </ul>
        </nav>

        <button
          className={`hamburger${menuOpen ? ' open' : ''}`}
          onClick={() => setMenuOpen(!menuOpen)}
          aria-label="Toggle menu"
          aria-expanded={menuOpen}
        >
          <span />
          <span />
          <span />
        </button>
      </div>

      <div className={`mobile-dropdown${menuOpen ? ' open' : ''}`}>
        <ul>
          <li>
            <Link to="/" className={isActive('/') ? 'active' : ''} onClick={closeMenu}>
              GAMES
            </Link>
          </li>
          <li>
            <Link to="/privacy" className={isActive('/privacy') ? 'active' : ''} onClick={closeMenu}>
              PRIVACY POLICY
            </Link>
          </li>
          <li>
            <Link to="/contact" className={isActive('/contact') ? 'active' : ''} onClick={closeMenu}>
              CONTACT
            </Link>
          </li>
          <li>
            <Link to="/feedback" className={isActive('/feedback') ? 'active' : ''} onClick={closeMenu}>
              FEEDBACK
            </Link>
          </li>
        </ul>
      </div>
    </header>
  );
}

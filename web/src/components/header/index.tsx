import { Link, useLocation } from 'react-router-dom';
import './style.css';

export default function Header() {
  const location = useLocation();

  const isActive = (path: string) => {
    if (path === '/' && location.pathname === '/') return true;
    if (path === '/privacy' && location.pathname === '/privacy') return true;
    if (path === '/contact' && location.pathname === '/contact') return true;
    if (path === '/news' && location.pathname === '/news') return true;
    if (path === '/' && location.pathname.startsWith('/games/')) return true;
    return false;
  };

  return (
    <header className="app-header">
      <div className="inner">
        <Link to="/" className="logo"></Link>

        <nav>
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
          </ul>
        </nav>
      </div>
    </header>
  );
}

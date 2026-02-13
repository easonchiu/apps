import { Link } from 'react-router-dom';
import './style.css';

export default function Footer() {
  return (
    <footer className="app-footer">
      <nav>
        <ul>
          <li>
            <Link to="/">GAMES</Link>
          </li>
          {/* <li>
            <Link to="/news">NEWS</Link>
          </li> */}
          <li>
            <Link to="/privacy">PRIVACY POLICY</Link>
          </li>
          <li>
            <Link to="/contact">CONTACT</Link>
          </li>
        </ul>
      </nav>
      <p>Copyright © 2023 YSGAME Studio. All rights reserved.</p>
    </footer>
  );
}

import { Link } from 'react-router-dom';
import './style.css';

export default function NotFound() {
  return (
    <div className="page-404">
      <main className="main">
        <p>
          404 - Page not found. <Link to="/">Go back home</Link>
        </p>
      </main>
    </div>
  );
}

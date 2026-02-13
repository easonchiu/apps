import { useParams, Navigate } from 'react-router-dom';
import { games } from '../../types/game';
import DownloadButton from '../../components/download-button';
import './style.css';

export default function GameDetail() {
  const { gameId } = useParams<{ gameId: string }>();
  const game = gameId ? games[gameId] : null;

  if (!game) {
    return <Navigate to="/404" replace />;
  }

  return (
    <div className="page-games">
      <main className="main">
        <div className="content">
          <div className="app-info">
            <img alt={game.name} src={game.icon} className="app-icon" />
            <div className="data">
              <h1>{game.name}</h1>
              <p>{game.description}</p>
            </div>
            <DownloadButton url={game.appStoreUrl} className="download" />
          </div>

          <div className="app-interface">
            <div className="list">
              {game.images.map((image, index) => (
                <img key={index} alt={game.name} src={image} />
              ))}
            </div>
          </div>

          <article>
            {game.content.map((paragraph, index) => (
              <p key={index}>{paragraph}</p>
            ))}
            <DownloadButton url={game.appStoreUrl} className="download" />
          </article>
        </div>
      </main>

      <DownloadButton url={game.appStoreUrl} className="download-bar" />
    </div>
  );
}

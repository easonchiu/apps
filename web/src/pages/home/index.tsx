import { Link } from 'react-router-dom';
import { games } from '../../types/game';
import DownloadButton from '../../components/download-button';
import './style.css';

export default function Home() {
  const gameList = Object.values(games);

  return (
    <div className="page-home">
      <main className="main">
        <div className="games">
          {gameList.map((game) => (
            <div key={game.id} className="game">
              <Link className="banner" to={`/games/${game.id}`}>
                <div className="list">
                  {game.images.map((image, index) => (
                    <img key={index} alt={game.name} src={image} />
                  ))}
                </div>
              </Link>
              <div className="info">
                <img alt={game.name} src={game.icon} className="app-icon" />
                <div className="data">
                  <h1>{game.name}</h1>
                  <p>{game.description}</p>
                </div>
                <DownloadButton url={game.appStoreUrl} className="download" />
              </div>
              <DownloadButton url={game.appStoreUrl} className="download-bar" />
            </div>
          ))}
        </div>
      </main>
    </div>
  );
}

import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Layout from './components/layout';
import Home from './pages/home';
import GameDetail from './pages/game-detail';
import Privacy from './pages/privacy';
import Contact from './pages/contact';
import NotFound from './pages/not-found';

function App() {
  return (
    <BrowserRouter>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/games/:gameId" element={<GameDetail />} />
          <Route path="/privacy" element={<Privacy />} />
          <Route path="/contact" element={<Contact />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </Layout>
    </BrowserRouter>
  );
}

export default App;

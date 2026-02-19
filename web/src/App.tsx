import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Layout from './components/layout';
import Home from './pages/home';
import GameDetail from './pages/game-detail';
import Privacy from './pages/privacy';
import Contact from './pages/contact';
import Feedback from './pages/feedback';
import NotFound from './pages/not-found';
import AdminLogin from './pages/admin/login';
import AdminFeedbacks from './pages/admin/feedbacks';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        {/* Admin routes — no Layout */}
        <Route path="/admin" element={<AdminLogin />} />
        <Route path="/admin/feedbacks" element={<AdminFeedbacks />} />

        {/* Main site routes — with Layout */}
        <Route path="/" element={<Layout><Home /></Layout>} />
        <Route path="/games/:gameId" element={<Layout><GameDetail /></Layout>} />
        <Route path="/privacy" element={<Layout><Privacy /></Layout>} />
        <Route path="/contact" element={<Layout><Contact /></Layout>} />
        <Route path="/feedback" element={<Layout><Feedback /></Layout>} />
        <Route path="*" element={<Layout><NotFound /></Layout>} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;

import { useState } from 'react';
import './style.css';

export default function RateUs() {
  const [email, setEmail] = useState('');
  const [content, setContent] = useState('');
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!email || !content) {
      setMessage('Please fill in all fields');
      return;
    }

    setLoading(true);
    setMessage('');

    try {
      const response = await fetch('/api/feedback', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, content }),
      });

      if (response.ok) {
        setMessage('Thank you for your feedback!');
        setEmail('');
        setContent('');
      } else {
        const data = await response.json();
        setMessage(data.error || 'Failed to submit feedback');
      }
    } catch (error) {
      setMessage('Network error. Please try again later.');
      console.error('Error:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="page-rate-us">
      <main className="main">
        <h1>Rate Us</h1>
        <p className="subtitle">We'd love to hear your feedback!</p>

        <div className="content">
          <form onSubmit={handleSubmit}>
            <div className="form-group">
              <label htmlFor="email">Email</label>
              <input
                type="email"
                id="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                placeholder="your@email.com"
                required
                disabled={loading}
              />
            </div>

            <div className="form-group">
              <label htmlFor="content">Your Feedback</label>
              <textarea
                id="content"
                value={content}
                onChange={(e) => setContent(e.target.value)}
                placeholder="Tell us what you think..."
                rows={8}
                required
                disabled={loading}
              />
            </div>

            {message && (
              <div className={`message ${message.includes('Thank you') ? 'success' : 'error'}`}>
                {message}
              </div>
            )}

            <button type="submit" disabled={loading}>
              {loading ? 'Submitting...' : 'Submit Feedback'}
            </button>
          </form>
        </div>
      </main>
    </div>
  );
}

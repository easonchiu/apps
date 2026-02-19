import { useState, useEffect, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';
import './style.css';

interface Feedback {
  id: string;
  email: string;
  content: string;
  app: string;
  status: string;
  created_at: string;
}

interface FeedbackResponse {
  data: Feedback[];
  total: number;
  page: number;
  pageSize: number;
}

const PAGE_SIZE = 20;

export default function AdminFeedbacks() {
  const [feedbacks, setFeedbacks] = useState<Feedback[]>([]);
  const [page, setPage] = useState(1);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [selectedIds, setSelectedIds] = useState<Set<string>>(new Set());
  const navigate = useNavigate();

  const totalPages = Math.max(1, Math.ceil(total / PAGE_SIZE));

  const isAllSelected = feedbacks.length > 0 && feedbacks.every((fb) => selectedIds.has(fb.id));

  const logout = useCallback(() => {
    localStorage.removeItem('admin_token');
    navigate('/admin', { replace: true });
  }, [navigate]);

  const fetchFeedbacks = useCallback(async (p: number) => {
    const token = localStorage.getItem('admin_token');
    if (!token) {
      navigate('/admin', { replace: true });
      return;
    }

    setLoading(true);
    setError('');

    try {
      const response = await fetch(
        `/wwwapi/admin/feedbacks?page=${p}&pageSize=${PAGE_SIZE}`,
        { headers: { Authorization: `Bearer ${token}` } }
      );

      if (response.status === 401) {
        logout();
        return;
      }

      if (!response.ok) {
        throw new Error('Failed to fetch feedbacks');
      }

      const data: FeedbackResponse = await response.json();
      setFeedbacks(data.data || []);
      setTotal(data.total);
    } catch {
      setError('Failed to load feedbacks. Please try again.');
    } finally {
      setLoading(false);
    }
  }, [navigate, logout]);

  useEffect(() => {
    const token = localStorage.getItem('admin_token');
    if (!token) {
      navigate('/admin', { replace: true });
      return;
    }
    fetchFeedbacks(page);
  }, [page, fetchFeedbacks, navigate]);

  // Clear selection when page changes
  useEffect(() => {
    setSelectedIds(new Set());
  }, [page]);

  const toggleSelect = (id: string) => {
    setSelectedIds((prev) => {
      const next = new Set(prev);
      if (next.has(id)) {
        next.delete(id);
      } else {
        next.add(id);
      }
      return next;
    });
  };

  const toggleSelectAll = () => {
    if (isAllSelected) {
      setSelectedIds(new Set());
    } else {
      setSelectedIds(new Set(feedbacks.map((fb) => fb.id)));
    }
  };

  const batchDelete = async () => {
    if (selectedIds.size === 0) return;
    if (!window.confirm(`Delete ${selectedIds.size} feedbacks?`)) return;

    const token = localStorage.getItem('admin_token');
    if (!token) { logout(); return; }

    try {
      const response = await fetch('/wwwapi/admin/feedbacks', {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ ids: Array.from(selectedIds) }),
      });

      if (response.status === 401) { logout(); return; }
      if (!response.ok) throw new Error('Failed to delete feedbacks');

      setSelectedIds(new Set());
      fetchFeedbacks(page);
    } catch {
      setError('Failed to delete feedbacks. Please try again.');
    }
  };

  const batchUpdateStatus = async (status: string) => {
    if (selectedIds.size === 0) return;

    const token = localStorage.getItem('admin_token');
    if (!token) { logout(); return; }

    try {
      const response = await fetch('/wwwapi/admin/feedbacks/status', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ ids: Array.from(selectedIds), status }),
      });

      if (response.status === 401) { logout(); return; }
      if (!response.ok) throw new Error('Failed to update status');

      setSelectedIds(new Set());
      fetchFeedbacks(page);
    } catch {
      setError('Failed to update status. Please try again.');
    }
  };

  const formatDate = (dateStr: string) => {
    try {
      const date = new Date(dateStr);
      return date.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
      });
    } catch {
      return dateStr;
    }
  };

  return (
    <div className="admin-feedbacks">
      <div className="admin-topbar">
        <h1>Feedback Management</h1>
        <button className="logout-btn" onClick={logout}>
          Logout
        </button>
      </div>

      {selectedIds.size > 0 && (
        <div className="batch-bar">
          <span>{selectedIds.size} selected</span>
          <div className="batch-actions">
            <button className="batch-btn pending" onClick={() => batchUpdateStatus('pending')}>
              Mark Pending
            </button>
            <button className="batch-btn processed" onClick={() => batchUpdateStatus('processed')}>
              Mark Processed
            </button>
            <button className="batch-btn delete" onClick={batchDelete}>
              Delete
            </button>
          </div>
        </div>
      )}

      <div className="table-wrap">
        {loading ? (
          <div className="loading-state">Loading...</div>
        ) : error ? (
          <div className="error-state">{error}</div>
        ) : feedbacks.length === 0 ? (
          <div className="empty-state">No feedbacks yet.</div>
        ) : (
          <>
            <div className="table-scroll">
              <table>
                <thead>
                  <tr>
                    <th className="col-check">
                      <div
                        className={`checkbox${isAllSelected ? ' checked' : ''}`}
                        onClick={toggleSelectAll}
                      />
                    </th>
                    <th className="col-email">Email</th>
                    <th className="col-app">App</th>
                    <th className="col-status">Status</th>
                    <th className="col-content">Content</th>
                    <th className="col-time">Time</th>
                  </tr>
                </thead>
                <tbody>
                  {feedbacks.map((fb) => (
                    <tr key={fb.id}>
                      <td className="col-check">
                        <div
                          className={`checkbox${selectedIds.has(fb.id) ? ' checked' : ''}`}
                          onClick={() => toggleSelect(fb.id)}
                        />
                      </td>
                      <td className="col-email">{fb.email}</td>
                      <td className="col-app">{fb.app || '—'}</td>
                      <td className="col-status">
                        <span className={`status-badge ${fb.status || 'pending'}`}>
                          {(fb.status || 'pending') === 'pending' ? 'Pending' : 'Processed'}
                        </span>
                      </td>
                      <td className="col-content">{fb.content}</td>
                      <td className="col-time">{formatDate(fb.created_at)}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>

            <div className="pagination">
              <button
                disabled={page <= 1}
                onClick={() => setPage((p) => p - 1)}
              >
                Previous
              </button>
              <span>
                Page {page} of {totalPages}
              </span>
              <button
                disabled={page >= totalPages}
                onClick={() => setPage((p) => p + 1)}
              >
                Next
              </button>
            </div>
          </>
        )}
      </div>
    </div>
  );
}

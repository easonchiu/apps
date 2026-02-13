import { type ReactNode } from 'react';
import Header from '../header';
import Footer from '../footer';
import './style.css';

interface LayoutProps {
  children: ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  return (
    <div className="app-layout">
      <Header />
      <div className="app-container">
        {children}
      </div>
      <Footer />
    </div>
  );
}

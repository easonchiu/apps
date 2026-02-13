import './style.css';

interface DownloadButtonProps {
  url: string;
  className?: string;
}

export default function DownloadButton({ url, className = '' }: DownloadButtonProps) {
  return (
    <a href={url} className={`download-button ${className}`} target="_blank" rel="noopener noreferrer">
      <span>Download on the</span>
      <strong>App Store</strong>
    </a>
  );
}

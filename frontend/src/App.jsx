import { useState, useEffect } from 'react';
import BookList from './components/BookList';
import AddBook from './components/AddBook';
import styles from './styles/App.module.css';

function App() {
  const [currentView, setCurrentView] = useState('list');
  const [version, setVersion] = useState('unknown');
  
  useEffect(() => {
    // Get version from environment variable
    const envVersion = import.meta.env.VITE_APP_VERSION || 'development';
    setVersion(envVersion);
  }, []);

  const handleBookAdded = () => {
    // Switch back to list view after adding a book
    setCurrentView('list');
  };

  return (
    <div className={styles.app}>
      <header className={styles.header}>
        <h1>AvaKart Bookstore</h1>
        <div className={styles.versionBadge}>Version: {version}</div>
      </header>
      <nav className={styles.nav}>
        <button
          className={currentView === 'list' ? styles.active : ''}
          onClick={() => setCurrentView('list')}
        >
          View Books
        </button>
        <button
          className={currentView === 'add' ? styles.active : ''}
          onClick={() => setCurrentView('add')}
        >
          Add New Book
        </button>
      </nav>

      {currentView === 'list' ? (
        <BookList />
      ) : (
        <AddBook onBookAdded={handleBookAdded} />
      )}
      
      <footer className={styles.footer}>
        <p>Running version: {version}</p>
      </footer>
    </div>
  );
}

export default App;

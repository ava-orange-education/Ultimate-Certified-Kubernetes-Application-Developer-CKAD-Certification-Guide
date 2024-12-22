import { useState } from 'react';
import BookList from './components/BookList';
import AddBook from './components/AddBook';
import styles from './styles/App.module.css';

function App() {
  const [currentView, setCurrentView] = useState('list');

  const handleBookAdded = () => {
    // Switch back to list view after adding a book
    setCurrentView('list');
  };

  return (
    <div className={styles.app}>
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
    </div>
  );
}

export default App;

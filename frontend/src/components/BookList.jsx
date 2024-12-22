import { useState, useEffect } from 'react';
import { api } from '../services/api';
import styles from '../styles/App.module.css';

const BookList = () => {
  const [books, setBooks] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        const data = await api.listBooks();
        setBooks(data);
        setError(null);
      } catch (err) {
        setError('Failed to fetch books. Please try again later.');
      } finally {
        setLoading(false);
      }
    };

    fetchBooks();
  }, []);

  const handlePurchase = async (bookId) => {
    try {
      // For demo purposes, using a hardcoded user ID
      await api.purchaseBook(bookId, 'user123');
      alert('Purchase successful!');
      // Refresh the book list to update quantities
      const updatedBooks = await api.listBooks();
      setBooks(updatedBooks);
    } catch (err) {
      alert('Failed to purchase book. Please try again.');
    }
  };

  if (loading) return <div className={styles.loading}>Loading books...</div>;
  if (error) return <div className={styles.error}>{error}</div>;

  return (
    <div className={styles.bookList}>
      <h1>Available Books</h1>
      <div className={styles.booksGrid}>
        {books.map((book) => (
          <div key={book.id} className={styles.bookCard}>
            <h2>{book.title}</h2>
            <p><strong>Author:</strong> {book.author}</p>
            <p><strong>Price:</strong> ${book.price}</p>
            <p><strong>Available:</strong> {book.quantity}</p>
            {book.description && <p>{book.description}</p>}
            <button 
              onClick={() => handlePurchase(book.id)}
              disabled={book.quantity < 1}
            >
              {book.quantity < 1 ? 'Out of Stock' : 'Purchase'}
            </button>
          </div>
        ))}
      </div>
    </div>
  );
};

export default BookList;

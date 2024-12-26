import { mockApi } from './mockData';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8081';

export const api = {
  listBooks: async () => {
    // return mockApi.listBooks();

    const response = await fetch(`${API_BASE_URL}/api/books/list`, {
      headers: {
        'Accept': 'application/json',
      },
    });
    if (!response.ok) {
      throw new Error('Failed to fetch books');
    }
    return response.json();
  },

  getBookDetails: async (bookId) => {
    // return mockApi.getBookDetails(bookId);

    const response = await fetch(`${API_BASE_URL}/api/books/details?id=${bookId}`, {
      headers: {
        'Accept': 'application/json',
      },
    });
    if (!response.ok) {
      throw new Error('Failed to fetch book details');
    }
    return response.json();
  },

  
  addBook: async (bookData) => {
    // return mockApi.addBook(bookData);

    const response = await fetch(`${API_BASE_URL}/api/books/add`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
      body: JSON.stringify(bookData),
    });
    if (!response.ok) {
      throw new Error('Failed to add book');
    }
    return response.json();
  },

  
  purchaseBook: async (bookId, userId) => {
    // return mockApi.purchaseBook(bookId);

    const response = await fetch(`${API_BASE_URL}/api/books/purchase`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
      body: JSON.stringify({
        book_id: bookId,
        quantity: 1,
      }),
    });
    if (!response.ok) {
      throw new Error('Failed to purchase book');
    }
    return response.json();
  },
};

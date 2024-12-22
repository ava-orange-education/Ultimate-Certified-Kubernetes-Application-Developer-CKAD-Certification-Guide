import { mockApi } from './mockData';

// Using mock API instead of real backend calls
export const api = {
  // Get all books
  listBooks: async () => {
    return mockApi.listBooks();
  },

  // Get book details
  getBookDetails: async (bookId) => {
    return mockApi.getBookDetails(bookId);
  },

  // Add a new book
  addBook: async (bookData) => {
    return mockApi.addBook(bookData);
  },

  // Purchase a book
  purchaseBook: async (bookId, userId) => {
    return mockApi.purchaseBook(bookId);
  },
};

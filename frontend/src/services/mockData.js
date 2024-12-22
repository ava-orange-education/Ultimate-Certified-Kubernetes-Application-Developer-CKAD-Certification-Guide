export const mockBooks = [
  {
    id: '1',
    title: 'The Kubernetes Book',
    author: 'Nigel Poulton',
    price: 29.99,
    quantity: 5,
    description: 'Updated for Kubernetes 1.21, this book is ideal for DevOps practitioners and container enthusiasts.',
    seller_id: 'seller123'
  },
  {
    id: '2',
    title: 'Docker Deep Dive',
    author: 'Nigel Poulton',
    price: 24.99,
    quantity: 3,
    description: 'Learn Docker from the ground up with hands-on examples and clear explanations.',
    seller_id: 'seller123'
  },
  {
    id: '3',
    title: 'Cloud Native DevOps with Kubernetes',
    author: 'John Arundel & Justin Domingus',
    price: 39.99,
    quantity: 2,
    description: 'Build, deploy, and scale modern applications in the cloud.',
    seller_id: 'seller123'
  }
];

// In-memory store for books
let books = [...mockBooks];

export const mockApi = {
  listBooks: async () => {
    return Promise.resolve([...books]);
  },

  getBookDetails: async (bookId) => {
    const book = books.find(b => b.id === bookId);
    if (!book) {
      return Promise.reject(new Error('Book not found'));
    }
    return Promise.resolve({...book});
  },

  addBook: async (bookData) => {
    const newBook = {
      ...bookData,
      id: String(books.length + 1)
    };
    books.push(newBook);
    return Promise.resolve(newBook);
  },

  purchaseBook: async (bookId) => {
    const bookIndex = books.findIndex(b => b.id === bookId);
    if (bookIndex === -1) {
      return Promise.reject(new Error('Book not found'));
    }
    if (books[bookIndex].quantity < 1) {
      return Promise.reject(new Error('Book out of stock'));
    }
    books[bookIndex] = {
      ...books[bookIndex],
      quantity: books[bookIndex].quantity - 1
    };
    return Promise.resolve({ success: true });
  }
};

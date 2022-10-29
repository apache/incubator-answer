import { RouterProvider } from 'react-router-dom';

import { routes, createBrowserRouter } from '@/router';

function App() {
  const router = createBrowserRouter(routes);
  return <RouterProvider router={router} />;
}

export default App;

import { BrowserRouter, Route, Routes } from "react-router-dom";

import Fiber from "./components/Fiber";
import NotFound from "./components/NotFound";
import React from "./components/React";

const App = () => (
  // Add basename to the <BrowserRouter basename="/web"> if you serve Single Page Application on "/web"
  <BrowserRouter>
    <Routes>
      <Route path="/" element={<Fiber />} />
      <Route path="/react" element={<React />} />
      <Route path="*" element={<NotFound />} />
    </Routes>
  </BrowserRouter>
);

export default App;

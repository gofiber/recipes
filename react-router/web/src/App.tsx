import { BrowserRouter, Route, Switch } from "react-router-dom";

import Fiber from "./components/Fiber";
import NotFound from "./components/NotFound";
import React from "./components/React";

const App = () => (
  <BrowserRouter>
    <Switch>
      <Route path="/" component={Fiber} exact />
      <Route path="/react" component={React} exact />
      <Route path="*" component={NotFound} />
    </Switch>
  </BrowserRouter>
);

export default App;

import { Link } from "react-router-dom";

import ReactLogo from "../assets/react-logo.svg";

const React = () => (
  <main className="application">
    <img src={ReactLogo} className="application-logo" alt="Logo of React" />

    <p>
      Edit <code>src/components/React.tsx</code> and save to reload.
    </p>

    <div className="application-links">
      <Link className="application-link" to="/">
        Go to Fiber page
      </Link>
      <Link className="application-link" to={{ pathname: "https://reactjs.org" }} target="_blank">
        Learn React, a JavaScript framework
      </Link>
    </div>
  </main>
);

export default React;

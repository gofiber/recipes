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
      <a className="application-link" href="https://reactjs.org" target="_blank" rel="noreferrer">
        Learn React, a JavaScript framework
      </a>
    </div>
  </main>
);

export default React;

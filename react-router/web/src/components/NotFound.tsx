import { Link } from "react-router-dom";

import FiberLogo from "../assets/fiber-logo.svg";
import ReactLogo from "../assets/react-logo.svg";

const NotFound = () => (
  <main className="application">
    <img src={FiberLogo} className="application-logo" alt="Logo of Fiber" />
    <img src={ReactLogo} className="application-logo" alt="Logo of React" />

    <p>Page not found! Let's go back home!</p>

    <div className="application-links">
      <Link className="application-link" to="/">
        Back home
      </Link>
    </div>
  </main>
);

export default NotFound;

import { Link } from "react-router-dom";

import FiberLogo from "../assets/fiber-logo.svg";

const Fiber = () => (
  <main className="application">
    <img src={FiberLogo} className="application-logo" alt="Logo of Fiber" />

    <p>
      Edit <code>src/components/Fiber.tsx</code> and save to reload.
    </p>

    <div className="application-links">
      <Link className="application-link" to="/react">
        Go to React page
      </Link>
      <a className="application-link" href="https://gofiber.io/" target="_blank" rel="noreferrer">
        Learn Fiber, a FastHTTP-based Go framework
      </a>
    </div>
  </main>
);

export default Fiber;

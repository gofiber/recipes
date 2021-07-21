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
      <Link className="application-link" to={{ pathname: "https://gofiber.io/" }} target="_blank">
        Learn Fiber, a FastHTTP-based Go framework
      </Link>
    </div>
  </main>
);

export default Fiber;

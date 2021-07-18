import "./styles/main.css";

import FiberLogo from "./assets/fiber-logo.svg";
import ReactLogo from "./assets/react-logo.svg";

function App() {
  return (
    <main className="application">
      <img src={FiberLogo} className="application-logo" alt="Logo of Fiber" />
      <img src={ReactLogo} className="application-logo" alt="Logo of React" />

      <p>
        Edit <code>src/App.tsx</code> and save to reload.
      </p>

      <div className="application-links">
        <a
          className="application-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React, a JavaScript framework
        </a>
        <a
          className="application-link"
          href="https://gofiber.io/"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn Fiber, a FastHTTP-based Go framework
        </a>
      </div>
    </main>
  );
}

export default App;

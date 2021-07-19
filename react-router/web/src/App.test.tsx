import { render, screen } from "@testing-library/react";
import App from "./App";

test("renders react text", () => {
  render(<App />);
  const linkElement = screen.getByText(/react page/i);
  expect(linkElement).toBeInTheDocument();
});

test("renders fiber text", () => {
  render(<App />);
  const linkElement = screen.getByText(/learn fiber/i);
  expect(linkElement).toBeInTheDocument();
});

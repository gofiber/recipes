import react from "@vitejs/plugin-react";
import { defineConfig } from "vitest/config";

export default defineConfig({
  plugins: [react()],
  // the Go server and the Dockerfile both read the frontend from web/build
  build: { outDir: "build" },
  test: { environment: "jsdom", globals: true, setupFiles: "./src/setupTests.ts" },
});

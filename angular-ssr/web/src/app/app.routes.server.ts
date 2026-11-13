import { RenderMode, ServerRoute } from '@angular/ssr';

// Render on every request, so the Fiber proxy has a live SSR server to talk to.
export const serverRoutes: ServerRoute[] = [
  {
    path: '**',
    renderMode: RenderMode.Server
  }
];

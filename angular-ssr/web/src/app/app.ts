import { isPlatformServer } from '@angular/common';
import { Component, PLATFORM_ID, inject } from '@angular/core';

@Component({
  selector: 'app-root',
  styleUrl: './app.css',
  templateUrl: './app.html',
})
export class App {
  // The HTML from the SSR server says "server"; after hydration the browser shows "browser".
  protected readonly renderedOn = isPlatformServer(inject(PLATFORM_ID)) ? 'server' : 'browser';
}

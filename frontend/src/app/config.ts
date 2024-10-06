import { ApplicationConfig } from '@angular/core';
import { provideRouter } from '@angular/router';
import { routes } from './routes';  // Import the routes

export const appConfig: ApplicationConfig = {
  providers: [
    provideRouter(routes),  // Use the routes defined in routes.ts
    // Other providers can be added here (e.g., HTTP interceptors, services)
  ]
};

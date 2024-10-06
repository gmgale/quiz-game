import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  standalone: true,  // Declare this component as standalone
  imports: [RouterModule]  // Import the RouterModule to use <router-outlet>
})
export class AppComponent {
  title = 'quiz-game';
}

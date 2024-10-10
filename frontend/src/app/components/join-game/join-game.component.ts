import { Component } from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms'; // FormsModule for ngModel two-way binding
import {DefaultService, Player} from "../../gen";

@Component({
  selector: 'app-join-game',
  templateUrl: './join-game.component.html',
  styleUrls: ['./join-game.component.css'],
  standalone: true,  // Declare this component as standalone
  imports: [CommonModule, FormsModule]  // Import FormsModule for form handling
})
export class JoinGameComponent {
  gameCode: string = '';
  gameID: string = '';
  playerName: string = '';
  errorMessage: string = '';

  constructor(
    private apiService: DefaultService,
    private router: Router,
    private route: ActivatedRoute)
  {}

  ngOnInit() {
    this.route.queryParams.subscribe(params => {
      this.gameCode = params['gameCode'] || '';
    });
  }

  joinGame() {
    if (!this.gameCode || !this.playerName) {
      this.errorMessage = 'Please enter both game code and your name.';
      return;
    }

    // Prepare the request body with both code and player name
    const requestBody = { code: this.gameCode, name: this.playerName };

    // Pass requestBody correctly in the API call
    this.apiService.gamesGameIdPlayersPost(this.gameCode, requestBody).subscribe({
      next: (player: Player) => {
        this.gameID = player.gameId as string;

        // Store player ID and game ID locally
        localStorage.setItem('playerId', player.id as string);
        localStorage.setItem('gameId', this.gameID);

        console.log('Player ID:', player);

        // Attempt to navigate to the game component
        this.router.navigate(['/game', player.gameId as string])
          .then(success => {
            if (!success) {
              // Handle unsuccessful navigation (e.g., route not found)
              console.error('Navigation failed: Route could not be activated.');
              // Optionally, display a user-friendly message or redirect to an error page
            }
          })
          .catch(err => {
            // Handle any errors that occur during navigation
            console.error('Navigation error:', err);
            // Optionally, display an error message to the user
          });
      },
      error: (err) => {
        // Handle errors from the API call
        console.error('API call error:', err);
        // Optionally, display an error message to the user
      }
    });
  }
}

import { Component } from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms'; // FormsModule for ngModel two-way binding
import { DefaultService} from "../../gen";

@Component({
  selector: 'app-join-game',
  templateUrl: './join-game.component.html',
  styleUrls: ['./join-game.component.css'],
  standalone: true,  // Declare this component as standalone
  imports: [CommonModule, FormsModule]  // Import FormsModule for form handling
})
export class JoinGameComponent {
  gameCode: string = '';
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
    this.apiService.gamesGameIdPlayersPost(this.gameCode, requestBody).subscribe(
      (player) => {
        // Store player ID and game ID locally
        localStorage.setItem('playerId', <string>player.id);
        localStorage.setItem('gameId', this.gameCode);

        // Navigate to the game component
        this.router.navigate(['/game', this.gameCode]);
      },
      (error) => {
        this.errorMessage = 'Failed to join the game. Please check the game code.';
      }
    );
  }
}

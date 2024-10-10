import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { NgIf } from "@angular/common";
import {DefaultService, GameSession} from '../../gen';

@Component({
  selector: 'app-create-game',
  templateUrl: './create-game.component.html',
  standalone: true,
  imports: [
    NgIf
  ],
  styleUrls: ['./create-game.component.css']
})

export class CreateGameComponent {
  gameId: string = '';
  gameCode: string = '';  // Add gameCode to store the code
  errorMessage: string = '';
  gameCreated: boolean = false;

  constructor(private apiService: DefaultService, private router: Router) { }

  createGame() {
    this.apiService.gamesPost().subscribe({
      next: (gameSession: GameSession) => {
        if (!gameSession.id || !gameSession.code) { // Check for both gameId and gameCode
          this.errorMessage = 'Failed to create a new game.';
          return;
        }
        this.gameId = gameSession.id;
        this.gameCode = gameSession.code;
        this.gameCreated = true;
      },
      error: (error: any) => {
        this.errorMessage = 'Failed to create a new game.';
        console.error('Create Game Error:', error);
      },
      complete: () => {
        console.log('Create game request completed.');
      }
    });
  }

  startGame() {
    this.apiService.gamesGameIdStartPost(this.gameId).subscribe({
      next: (response: any) => {
        // Navigate to the game component for the host
        this.router.navigate(['/game', this.gameId]);
      },
      error: (error: any) => {
        this.errorMessage = 'Failed to start the game.';
        console.error('Start Game Error:', error);
      },
      complete: () => {
        console.log('Start game request completed.');
      }
    });
  }

}

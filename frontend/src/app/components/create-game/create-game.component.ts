import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from '../../services/api.service';
import {NgIf} from "@angular/common";

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
  errorMessage: string = '';
  gameCreated: boolean = false;


  constructor(private apiService: ApiService, private router: Router) { }

  createGame() {
    this.apiService.createGame().subscribe(
      (gameSession) => {
        this.gameId = gameSession.id;
        this.gameCreated = true;
      },
      (error) => {
        this.errorMessage = 'Failed to create a new game.';
      }
    );
  }

  startGame() {
    this.apiService.startGame(this.gameId).subscribe(
      (response) => {
        // Navigate to the game component for the host
        this.router.navigate(['/game', this.gameId]);
      },
      (error) => {
        this.errorMessage = 'Failed to start the game.';
      }
    );
  }
}

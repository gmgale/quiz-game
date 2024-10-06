import { Routes } from '@angular/router';
import { JoinGameComponent } from './components/join-game/join-game.component';
import { GameComponent } from './components/game/game.component';
import { LeaderboardComponent } from './components/leaderboard/leaderboard.component';

export const routes: Routes = [
  { path: '', component: JoinGameComponent },
  { path: 'game/:gameId', component: GameComponent },
  { path: 'leaderboard/:gameId', component: LeaderboardComponent },
  { path: '**', redirectTo: '' }  // Wildcard route for unknown paths
];

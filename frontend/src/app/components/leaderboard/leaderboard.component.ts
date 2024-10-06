import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import {NgForOf} from "@angular/common";

@Component({
  selector: 'app-leaderboard',
  templateUrl: './leaderboard.component.html',
  standalone: true,
  imports: [
    NgForOf
  ],
  styleUrls: ['./leaderboard.component.css']
})
export class LeaderboardComponent implements OnInit {
  leaderboard: any[] = [];
  gameId: string | null = '';

  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.gameId = this.route.snapshot.paramMap.get('gameId');
    // Assuming the leaderboard data was stored in localStorage
    const leaderboardData = localStorage.getItem('leaderboard');
    if (leaderboardData) {
      this.leaderboard = JSON.parse(leaderboardData);
    } else {
      // Handle case when data is not available
      this.leaderboard = [];
    }
  }
}

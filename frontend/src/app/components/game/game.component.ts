import { Component, OnInit, OnDestroy } from '@angular/core';
import { CommonModule } from '@angular/common';
import { WebsocketService } from '../../services/websocket.service';
import { Subscription } from 'rxjs';
import { ApiService } from '../../services/api.service';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css'],
  standalone: true,  // Declare this component as standalone
  imports: [CommonModule, RouterModule]  // Import necessary Angular modules
})
export class GameComponent implements OnInit, OnDestroy {
  private wsSubscription: Subscription = new Subscription();
  currentQuestion: any = null;
  timeLeft: number = 0;
  timerInterval: any;
  gameId: string | null = '';
  playerId: string | null = '';
  showAnswerFeedback: boolean = false;
  answerFeedback: string = '';
  isAnswered: boolean = false;
  optionsDisabled: boolean = false;

  constructor(
    private websocketService: WebsocketService,
    private apiService: ApiService,
    private route: ActivatedRoute,
    private router: Router
  ) { }

  ngOnInit() {
    this.gameId = this.route.snapshot.paramMap.get('gameId');
    this.playerId = localStorage.getItem('playerId');

    if (!this.playerId) {
      this.router.navigate(['/']);
      return;
    }

    if (this.gameId) {
      this.wsSubscription = this.websocketService.connect(this.gameId).subscribe(msg => {
        if (msg.type === 'question') {
          this.currentQuestion = msg.data;
          this.timeLeft = this.currentQuestion.timeLimit;
          this.startTimer();
          this.isAnswered = false;
          this.showAnswerFeedback = false;
          this.optionsDisabled = false;
        } else if (msg.type === 'game_over') {
          localStorage.setItem('leaderboard', JSON.stringify(msg.data));
          this.router.navigate(['/leaderboard', this.gameId]);
        }
      });
    }
  }

  startTimer() {
    if (this.timerInterval) {
      clearInterval(this.timerInterval);
    }
    this.timerInterval = setInterval(() => {
      this.timeLeft--;
      if (this.timeLeft <= 0) {
        clearInterval(this.timerInterval);
        this.optionsDisabled = true;
      }
    }, 1000);
  }

  submitAnswer(optionIndex: number) {
    if (this.isAnswered || this.optionsDisabled) {
      return;
    }
    this.isAnswered = true;
    this.optionsDisabled = true;
    const responseTime = (this.currentQuestion.timeLimit - this.timeLeft) * 1000;

    const answer = {
      playerId: this.playerId,
      questionId: this.currentQuestion.id,
      selectedOption: optionIndex,
      responseTime: responseTime,
    };

    this.currentQuestion.selectedOption = optionIndex;

    if (this.gameId) {
      this.apiService.submitAnswer(this.gameId, answer).subscribe(
        response => {
          if (response.correct) {
            this.answerFeedback = `Correct! You scored ${response.scoreAwarded} points.`;
          } else {
            const correctOptionIndex = response.correctOption;
            const correctOptionText = this.currentQuestion.options[correctOptionIndex];
            this.answerFeedback = `Incorrect. The correct answer was: ${correctOptionText}`;
          }
          this.showAnswerFeedback = true;
        },
        error => {
          this.answerFeedback = 'Error submitting your answer.';
          this.showAnswerFeedback = true;
        }
      );
    }
  }

  ngOnDestroy() {
    if (this.wsSubscription) {
      this.wsSubscription.unsubscribe();
    }
    this.websocketService.close();
    if (this.timerInterval) {
      clearInterval(this.timerInterval);
    }
  }
}

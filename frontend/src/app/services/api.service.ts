import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private baseUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  createGame(): Observable<any> {
    return this.http.post<any>(`${this.baseUrl}/games`, {});
  }

  joinGame(gameId: string, playerName: string): Observable<any> {
    const body = { name: playerName };
    return this.http.post<any>(`${this.baseUrl}/games/${gameId}/players`, body);
  }

  submitAnswer(gameId: string, answer: any): Observable<any> {
    return this.http.post<any>(`${this.baseUrl}/games/${gameId}/answers`, answer);
  }

  getLeaderboard(gameId: string): Observable<any> {
    return this.http.get<any>(`${this.baseUrl}/games/${gameId}/leaderboard`);
  }

  startGame(gameId: string): Observable<any> {
    return this.http.post<any>(`${this.baseUrl}/games/${gameId}/start`, {});
  }
}

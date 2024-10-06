import { Injectable } from '@angular/core';
import { webSocket, WebSocketSubject, WebSocketSubjectConfig } from 'rxjs/webSocket';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class WebsocketService {
  private socket$: WebSocketSubject<any> | undefined;
  private config: WebSocketSubjectConfig<any> | undefined;

  connect(gameId: string): Observable<any> {
    if (!this.socket$ || this.socket$.closed) {
      console.log(`${environment.apiUrl}/ws/${gameId}`)
      this.config = {
        url: `${environment.apiUrl}/ws/${gameId}`,
        closeObserver: {
          next: () => {
            console.log('WebSocket closed');
            this.socket$ = undefined;
          }
        }
      };
      this.socket$ = webSocket(this.config);
    }
    return this.socket$;
  }

  sendMessage(msg: any) {
    this.socket$?.next(msg);
  }

  close() {
    this.socket$?.complete();
  }
}

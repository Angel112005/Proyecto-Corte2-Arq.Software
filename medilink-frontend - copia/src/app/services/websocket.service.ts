import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class WebSocketService {
  private socket!: WebSocket;
  private notificationsSubject = new Subject<any>();
  public notifications$ = this.notificationsSubject.asObservable();

  constructor() {
    this.connect();
  }

  private connect(): void {
    this.socket = new WebSocket('ws://localhost:8082/ws');

    this.socket.onopen = () => {
      console.log('🟢 Conectado al WebSocket');
    };

    this.socket.onmessage = (event) => {
      const message = JSON.parse(event.data);
      console.log('📩 Notificación recibida:', message);
      this.notificationsSubject.next(message);
    };

    this.socket.onclose = () => {
      console.warn('🔴 WebSocket desconectado, reintentando...');
      setTimeout(() => this.connect(), 5000);
    };
  }
}

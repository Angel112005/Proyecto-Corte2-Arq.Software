import { Injectable } from '@angular/core';
import { Observable, Subject, BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class WebSocketService {
  private socket!: WebSocket;
  private notificationsSubject = new Subject<any>(); // Observable para recibir notificaciones
  public notifications$ = this.notificationsSubject.asObservable();

  private connectionStatus = new BehaviorSubject<boolean>(false); // Estado de conexión
  public connectionStatus$ = this.connectionStatus.asObservable();

  private reconnectAttempts = 0; // Contador de intentos de reconexión
  private maxReconnectAttempts = 5; // Intentos máximos antes de detenerse
  private reconnectInterval = 2000; // Tiempo inicial de reconexión (2s)

  constructor() {
    this.connect();
  }

  private connect(): void {
    if (this.reconnectAttempts >= this.maxReconnectAttempts) {
      console.error('❌ WebSocket: Se alcanzó el límite de intentos de reconexión.');
      return;
    }

    this.socket = new WebSocket('ws://localhost:8082/ws'); // URL del WebSocket

    this.socket.onopen = () => {
      console.log('🟢 Conectado al WebSocket');
      this.connectionStatus.next(true);
      this.reconnectAttempts = 0; // Reiniciar intentos al conectarse
      this.startHeartbeat(); // Iniciar el "ping-pong"
    };

    this.socket.onmessage = (event) => {
      console.log('📩 Mensaje recibido del WebSocket:', event.data);
      try {
        const notification = JSON.parse(event.data);
        this.notificationsSubject.next(notification);
      } catch (error) {
        console.error('❌ Error al procesar el mensaje del WebSocket:', error);
      }
    };

    this.socket.onerror = (error) => {
      console.error('❌ Error en el WebSocket:', error);
    };

    this.socket.onclose = () => {
      console.warn(`🔴 WebSocket desconectado. Reintentando en ${this.reconnectInterval / 1000}s...`);
      this.connectionStatus.next(false);

      this.reconnectAttempts++;
      setTimeout(() => this.connect(), this.reconnectInterval);

      // Aumentar tiempo de espera para evitar reconexiones rápidas
      this.reconnectInterval = Math.min(this.reconnectInterval * 2, 30000); // Máximo 30s
    };
  }

  private startHeartbeat(): void {
    setInterval(() => {
      if (this.socket.readyState === WebSocket.OPEN) {
        this.socket.send(JSON.stringify({ type: 'ping' })); // Enviar ping al servidor
      }
    }, 5000); // Enviar cada 5 segundos
  }
}

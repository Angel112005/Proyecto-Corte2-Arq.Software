import { Component, OnInit } from '@angular/core';
import { WebSocketService } from '../services/websocket.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {
  notifications: any[] = [];

  constructor(private router: Router, private wsService: WebSocketService) {}

  navigateToAppointmentForm(): void {
    this.router.navigate(['/new']);
  }

  ngOnInit(): void {
    this.wsService.notifications$.subscribe((notification) => {
      console.log('📢 Nueva notificación:', notification);

      // Agregar la notificación con la propiedad dismissed
      const newNotification = { ...notification, dismissed: false };
      this.notifications.unshift(newNotification);

      // Hacer que la notificación desaparezca después de 5 segundos
      setTimeout(() => this.dismissNotification(newNotification), 9000);
    });
  }

  // Función para cerrar la notificación manualmente
  dismissNotification(notification: any) {
    notification.dismissed = true;
    setTimeout(() => {
      this.notifications = this.notifications.filter(n => n !== notification);
    }, 2000); // 1s extra para que termine la animación
  }
}

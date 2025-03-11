import { Injectable } from '@angular/core';
import { AppointmentApiService } from '../adapters/appointment-api.service';
import { Appointment } from '../entities/appointment-interface';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AppointmentService {
  constructor(private appointmentApi: AppointmentApiService) {}


  getAppointments(): Observable<Appointment[]> {
    return this.appointmentApi.getAppointments();
  }


  getAppointmentById(id: number): Observable<Appointment> {
    return this.appointmentApi.getAppointmentById(id);
  }


  createAppointment(appointment: Appointment): Observable<Appointment> {
    return this.appointmentApi.createAppointment(appointment);
  }


  updateAppointment(id: number, appointment: Appointment): Observable<Appointment> {
    return this.appointmentApi.updateAppointment(id, appointment);
  }


  deleteAppointment(id: number): Observable<void> {
    return this.appointmentApi.deleteAppointment(id);
  }
}

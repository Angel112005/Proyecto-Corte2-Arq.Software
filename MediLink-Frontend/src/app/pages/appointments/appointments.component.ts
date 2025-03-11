import { Component, OnInit } from '@angular/core';
import { AppointmentService } from '../../usecases/appointment.service';
import { Appointment } from '../../entities/appointment-interface';
import { Router } from '@angular/router';

@Component({
  selector: 'app-appointments',
  templateUrl: './appointments.component.html',
  styleUrl: './appointments.component.css'
})
export class AppointmentsComponent implements OnInit {
  appointments: Appointment[] = [];

  constructor(private appointmentService: AppointmentService, private router: Router) {}

  ngOnInit(): void {
    this.appointmentService.getAppointments().subscribe((data) => {
      this.appointments = data;
    });
  }

  editAppointment(id: number): void {
    console.log("ID de la cita enviada:", id)
    console.log(this.appointments)
    this.router.navigate(['/appointment/edit', id]); // Redirige al formulario de edición
  }

  deleteAppointment(id: number): void {
    this.appointmentService.deleteAppointment(id).subscribe(() => {
      this.appointments = this.appointments.filter((a) => a.id !== id);
    });
  }
}

import { DoctorService } from './../../../doctors/services/doctor.service';
import { Component, OnInit } from '@angular/core';
import { AppointmentService } from '../../services/appointment.service';
import { Appointment } from '../../../../core/models/appointment-interface';
import { Doctor } from '../../../../core/models/doctor-interface';
// import { DoctorService } from '../../../doctors/services/doctor.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-appointment-list',
  templateUrl: './appointment-list.component.html',
  styleUrl: './appointment-list.component.css'
})
export class AppointmentListComponent implements OnInit {
  appointments: Appointment[] = [];
  doctors: Doctor[] = [];


  constructor(private appointmentService: AppointmentService, private router: Router, private doctorService: DoctorService) {}

  ngOnInit(): void {
    this.loadAppointments();

    this.doctorService.getDoctors().subscribe({
      next: (data) => {
        this.doctors = data;
        console.log('📋 Doctores cargados:', this.doctors);
      },
      error: (err) => {
        console.error('❌ Error al obtener los doctores:', err);
      },
    });
  }

  getDoctorNameById(id: number): string {
    const doctor = this.doctors.find(d => d.id === id);
    return doctor ? doctor.name : 'Desconocido';
  }


  private loadAppointments(): void {
    this.appointmentService.getAppointments().subscribe({
      next: (data) => {
        this.appointments = data;
        console.log('📌 Citas obtenidas:', this.appointments);
      },
      error: (err) => {
        console.error('❌ Error al obtener las citas:', err);
      },
    });
  }

  editAppointment(id: number): void {
    console.log('📝 Editar cita ID:', id);
    this.router.navigate(['/appointments/edit', id]); // Redirige al formulario de edición
  }

  deleteAppointment(id: number): void {
    this.appointmentService.deleteAppointment(id).subscribe({
      next: () => {
        this.appointments = this.appointments.filter((a) => a.id !== id);
        console.log('🗑️ Cita eliminada con éxito');
      },
      error: (err) => {
        console.error('❌ Error al eliminar la cita:', err);
      },
    });
  }
}

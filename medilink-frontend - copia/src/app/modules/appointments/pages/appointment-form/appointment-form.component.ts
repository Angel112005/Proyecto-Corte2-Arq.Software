import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AppointmentService } from '../../services/appointment.service';
import { DoctorService } from '../../../doctors/services/doctor.service';
import { Appointment } from '../../../../core/models/appointment-interface';
import { Doctor } from '../../../../core/models/doctor-interface';
import Swal from 'sweetalert2';


@Component({
  selector: 'app-appointment-form',
  templateUrl: './appointment-form.component.html',
  styleUrl: './appointment-form.component.css'
})
export class AppointmentFormComponent implements OnInit {
  appointment: Appointment = { id: 0, patient_name: '', doctor_id: 0, appointment_date: '', status: 'agendado' };
  doctors: Doctor[] = [];
  isEdit: boolean = false;

  constructor(
    private appointmentService: AppointmentService,
    private doctorService: DoctorService,
    private route: ActivatedRoute,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.doctorService.getDoctors().subscribe({
      next: (data) => {
        this.doctors = data;
      },
      error: (err) => {
        console.error('Error al cargar los doctores:', err);
      },
    });

    const id = this.route.snapshot.params['id'];
    if (id) {
      this.isEdit = true;
      this.appointmentService.getAppointmentById(id).subscribe({
        next: (data) => {
          this.appointment = data;
          this.appointment.doctor_id = Number(this.appointment.doctor_id);
          this.appointment.status = 'reagendado';
              // 🔧 Ajuste del formato de fecha para <input type="datetime-local">
          if (this.appointment.appointment_date.includes(' ')) {
            this.appointment.appointment_date = this.appointment.appointment_date.replace(' ', 'T');
          }
        },
        error: (err) => {
          console.error('Error al cargar la cita:', err);
        },
      });
    }
  }

  submitForm(): void {
    this.appointment.doctor_id = Number(this.appointment.doctor_id);

    if (this.isEdit) {
      this.appointmentService.updateAppointment(this.appointment.id, this.appointment).subscribe({
        next: () => {
          this.showSuccessAlert('Cita actualizada con éxito.');
          this.router.navigate(['/appointments']);
        },
        error: (err) => {
          this.showErrorAlert('Error al actualizar la cita.');
          console.error('Error al actualizar la cita:', err);
        },
      });
    } else {
      this.appointmentService.createAppointment(this.appointment).subscribe({
        next: () => {
          this.showSuccessAlert('Cita creada con éxito.');
          this.router.navigate(['']);
        },
        error: (err) => {
          this.showErrorAlert('Error al crear la cita.');
          console.error('Error al crear la cita:', err);
        },
      });
    }
  }

  showSuccessAlert(message: string): void {
    Swal.fire({
      icon: 'success',
      title: '¡Éxito!',
      text: message,
      showConfirmButton: false,
      timer: 1000
    });
  }

  showErrorAlert(message: string): void {
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: message,
      confirmButtonColor: '#d33',
    });
  }
}

import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AppointmentService } from '../../usecases/appointment.service';
import { DoctorService } from '../../usecases/doctor.service';
import { Appointment } from '../../entities/appointment-interface';
import { Doctor } from '../../entities/doctor-interface';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-appointment-form',
  templateUrl: './appointment-form.component.html',
  styleUrl: './appointment-form.component.css'
})
export class AppointmentFormComponent implements OnInit {
  appointment: Appointment = {
    id: 0,
    patient_name: '',
    doctor_id: 0,
    appointment_date: '',
    status: 'agendado', // Estado inicial
  }; // Modelo de la cita
  doctors: Doctor[] = []; // Lista de doctores
  isEdit: boolean = false; // Estado para diferenciar entre creación y edición

  constructor(
    private appointmentService: AppointmentService,
    private doctorService: DoctorService, // Importamos el servicio de doctores
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
      this.isEdit = true; // Modo edición si existe un ID
      this.appointmentService.getAppointmentById(id).subscribe({
        next: (data) => {
          this.appointment = data;
          this.appointment.doctor_id = Number(this.appointment.doctor_id); // Convertir a número
          this.appointment.status = 'reagendado'; // Estado automático al editar
        },
        error: (err) => {
          console.error('Error al cargar la cita:', err);
        },
      });
    }
  }

  submitForm(): void {
    this.appointment.doctor_id = Number(this.appointment.doctor_id); // Convertir a número antes de enviarlo

    if (this.isEdit) {
      this.appointmentService.updateAppointment(this.appointment.id, this.appointment).subscribe({
        next: () => {
          this.showSuccessAlert('Cita actualizada con éxito.');
          this.router.navigate(['/appointments']); // Redirigir a la lista de citas
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
          this.router.navigate(['/']); // Redirigir a la página principal
        },
        error: (err) => {
          this.showErrorAlert('Error al crear la cita.');
          console.error('Error al crear la cita:', err);
        },
      });
    }
  }

  // Mostrar alerta de éxito con SweetAlert2
  showSuccessAlert(message: string): void {
    Swal.fire({
      icon: 'success',
      title: '¡Éxito!',
      text: message,
      showConfirmButton: false,
      timer: 1000 // Se cierra automáticamente después de 2.5 segundos
    });
  }

  // Mostrar alerta de error con SweetAlert2
  showErrorAlert(message: string): void {
    Swal.fire({
      icon: 'error',
      title: 'Oops...',
      text: message,
      confirmButtonColor: '#d33',
    });
  }



  // submitForm(): void {
  //   this.appointment.doctor_id = Number(this.appointment.doctor_id); // Convertir a número antes de enviarlo

  //   if (this.isEdit) {
  //     // Actualizar cita existente
  //     this.appointmentService.updateAppointment(this.appointment.id, this.appointment).subscribe({
  //       next: () => {
  //         console.log('Cita actualizada con éxito.');
  //         console.log(this.appointment)
  //         this.router.navigate(['/appointments']); // Redirigir a la lista de citas
  //       },
  //       error: (err) => {
  //         console.error('Error al actualizar la cita:', err);
  //       },
  //     });
  //   } else {
  //     // Crear nueva cita
  //     this.appointmentService.createAppointment(this.appointment).subscribe({
  //       next: () => {
  //         console.log('Cita creada con éxito.');
  //         console.log(this.appointment)
  //         this.router.navigate(['']); // Redirigir a la lista de citas
  //       },
  //       error: (err) => {
  //         console.error('Error al crear la cita:', err);
  //       },
  //     });
  //   }
  // }



}

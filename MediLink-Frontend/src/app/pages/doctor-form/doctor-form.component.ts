import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { DoctorService } from '../../usecases/doctor.service';
import { Doctor } from '../../entities/doctor-interface';
import Swal from 'sweetalert2';


@Component({
  selector: 'app-doctor-form',
  templateUrl: './doctor-form.component.html',
  styleUrl: './doctor-form.component.css'
})
export class DoctorFormComponent implements OnInit {
  doctor: Doctor = { id: 0, name: '', specialty: '', experience: 0 }; // Modelo inicial
  isEdit: boolean = false; // Indica si estamos en modo edición

  constructor(
    private doctorService: DoctorService,
    private route: ActivatedRoute,
    private router: Router
  ) {}

  ngOnInit(): void {
    // Verificar si estamos editando (ID en la ruta)
    const id = this.route.snapshot.params['id'];
    if (id) {
      this.isEdit = true;
      this.doctorService.getDoctorById(id).subscribe({
        next: (data) => {
          this.doctor = data; // Cargar los datos del doctor existente
        },
        error: (err) => {
          console.error('Error al cargar el doctor:', err);
        },
      });
    }
  }

  // submitForm(): void {
  //   if (this.isEdit) {
  //     // Actualizar doctor existente
  //     this.doctorService.updateDoctor(this.doctor.id, this.doctor).subscribe({
  //       next: () => {
  //         console.log('Doctor actualizado con éxito');
  //         this.router.navigate(['/doctors']); // Redirigir a la lista de doctores
  //       },
  //       error: (err) => {
  //         console.error('Error al actualizar el doctor:', err);
  //       },
  //     });
  //   } else {
  //     // Crear nuevo doctor
  //     this.doctorService.createDoctor(this.doctor).subscribe({
  //       next: () => {
  //         console.log('Doctor creado con éxito');
  //         this.router.navigate(['/doctors']); // Redirigir a la lista de doctores
  //       },
  //       error: (err) => {
  //         console.error('Error al crear el doctor:', err);
  //       },
  //     });
  //   }
  // }

  submitForm(): void {
    if (this.isEdit) {
      // Actualizar doctor existente
      this.doctorService.updateDoctor(this.doctor.id, this.doctor).subscribe({
        next: () => {
          this.showSuccessAlert('Doctor actualizado con éxito.');
          this.router.navigate(['/doctors']); // Redirigir a la lista de doctores
        },
        error: (err) => {
          this.showErrorAlert('Error al actualizar el doctor.');
          console.error('Error al actualizar el doctor:', err);
        },
      });
    } else {
      // Crear nuevo doctor
      this.doctorService.createDoctor(this.doctor).subscribe({
        next: () => {
          this.showSuccessAlert('Doctor creado con éxito.');
          this.router.navigate(['/doctors']); // Redirigir a la lista de doctores
        },
        error: (err) => {
          this.showErrorAlert('Error al crear el doctor.');
          console.error('Error al crear el doctor:', err);
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
      timer: 2500 // Se cierra automáticamente después de 2.5 segundos
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

}

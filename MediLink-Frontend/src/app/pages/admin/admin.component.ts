import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DoctorService } from '../../usecases/doctor.service';
import { Doctor } from '../../entities/doctor-interface';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrl: './admin.component.css'
})
export class AdminComponent implements OnInit {
  doctors: Doctor[] = [];

  constructor(private router: Router, private doctorService: DoctorService) {}

  ngOnInit(): void {
    // Cargar la lista de doctores
    this.doctorService.getDoctors().subscribe({
      next: (data) => {
        this.doctors = data;
      },
      error: (err) => {
        console.error('Error al cargar doctores:', err);
      },
    });
  }

  // Método para navegar al formulario de nuevo doctor
  navigateToNewDoctor(): void {
    this.router.navigate(['/doctors-new']);
  }

  // Método para editar un doctor existente
  editDoctor(id: number): void {
    this.router.navigate(['/doctor/edit', id]);
  }

  // Método para eliminar un doctor
  deleteDoctor(id: number): void {
    this.doctorService.deleteDoctor(id).subscribe({
      next: () => {
        this.doctors = this.doctors.filter((doctor) => doctor.id !== id);
        console.log('Doctor eliminado con éxito.');
      },
      error: (err) => {
        console.error('Error al eliminar doctor:', err);
      },
    });
  }
}

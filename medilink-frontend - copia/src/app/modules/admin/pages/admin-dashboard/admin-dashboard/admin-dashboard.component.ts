import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Doctor } from './../../../../../core/models/doctor-interface';
import { DoctorService } from '../../../../doctors/services/doctor.service';
// import { Doctor } from '../../../../core/models/doctor-interface';

@Component({
  selector: 'app-admin-dashboard',
  templateUrl: './admin-dashboard.component.html',
  styleUrl: './admin-dashboard.component.css'
})
export class AdminDashboardComponent implements OnInit {
  doctors: Doctor[] = [];

  constructor(private router: Router, private doctorService: DoctorService) {}

  ngOnInit(): void {
    this.doctorService.getDoctors().subscribe({
      next: (data) => {
        this.doctors = data;
        console.log('📌 Doctores obtenidos en Admin Dashboard:', this.doctors); // ✅ Verifica si los datos llegan
      },
      error: (err) => {
        console.error('❌ Error al cargar doctores en Admin Dashboard:', err);
      },
    });
  }



  // Método para navegar al formulario de nuevo doctor
  navigateToNewDoctor(): void {
    this.router.navigate(['/doctors/new']);
  }

  // Método para editar un doctor existente
  editDoctor(id: number): void {
    this.router.navigate(['/doctors/edit', id]);
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

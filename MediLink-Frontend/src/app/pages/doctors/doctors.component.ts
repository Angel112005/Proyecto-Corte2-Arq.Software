import { Component, OnInit } from '@angular/core';
import { DoctorService } from '../../usecases/doctor.service';
import { Doctor } from '../../entities/doctor-interface';

@Component({
  selector: 'app-doctors',
  templateUrl: './doctors.component.html',
  styleUrl: './doctors.component.css'
})
export class DoctorsComponent implements OnInit {
  doctors: Doctor[] = []; // Lista de doctores

  constructor(private doctorService: DoctorService) {}

  ngOnInit(): void {
    // Llamar al servicio para obtener los doctores
    this.doctorService.getDoctors().subscribe({
      next: (data) => {
        this.doctors = data; // Asignar datos a la lista de doctores
        console.log('Doctores obtenidos:', this.doctors); // Mostrar en consola
      },
      error: (err) => {
        console.error('Error al obtener los doctores:', err); // Mostrar errores
      },
    });
  }
}

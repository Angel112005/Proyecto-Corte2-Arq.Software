import { Injectable } from '@angular/core';
import { DoctorApiService } from '../adapters/doctor-api.service';
import { Doctor } from '../entities/doctor-interface';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DoctorService {
  constructor(private doctorApi: DoctorApiService) {}


  getDoctors(): Observable<Doctor[]> {
    return this.doctorApi.getDoctors();
  }


  getDoctorById(id: number): Observable<Doctor> {
    return this.doctorApi.getDoctorById(id);
  }


  createDoctor(doctor: Doctor): Observable<Doctor> {
    return this.doctorApi.createDoctor(doctor);
  }


  updateDoctor(id: number, doctor: Doctor): Observable<Doctor> {
    return this.doctorApi.updateDoctor(id, doctor);
  }

  
  deleteDoctor(id: number): Observable<void> {
    return this.doctorApi.deleteDoctor(id);
  }
}

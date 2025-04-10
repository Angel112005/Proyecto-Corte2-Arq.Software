import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { DoctorService } from '../../services/doctor.service';
import { Doctor } from '../../../../core/models/doctor-interface';

@Component({
  selector: 'app-doctor-form',
  templateUrl: './doctor-form.component.html',
  styleUrl: './doctor-form.component.css'
})
export class DoctorFormComponent implements OnInit {
  doctor: Doctor = { id: 0, name: '', specialty: '', experience: 0 };
  isEdit: boolean = false;

  constructor(
    private doctorService: DoctorService,
    private route: ActivatedRoute,
    private router: Router
  ) {}

  ngOnInit(): void {
    const id = this.route.snapshot.params['id'];
    if (id) {
      this.isEdit = true;
      this.doctorService.getDoctorById(id).subscribe((data) => {
        this.doctor = data;
      });
    }
  }

  submitForm(): void {
    if (this.isEdit) {
      this.doctorService.updateDoctor(this.doctor.id, this.doctor).subscribe(() => {
        this.router.navigate(['/doctors']);
      });
    } else {
      this.doctorService.createDoctor(this.doctor).subscribe(() => {
        this.router.navigate(['/doctors']);
      });
    }
  }
}

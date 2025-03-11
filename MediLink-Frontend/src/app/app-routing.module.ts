import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { DoctorsComponent } from './pages/doctors/doctors.component';
import { AppointmentsComponent } from './pages/appointments/appointments.component';
import { AdminComponent } from './pages/admin/admin.component';
import { AppointmentFormComponent } from './pages/appointment-form/appointment-form.component';
import { DoctorFormComponent } from './pages/doctor-form/doctor-form.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'doctors', component: DoctorsComponent },
  { path: '', redirectTo: '/doctors', pathMatch: 'full' },
  { path: 'appointments', component: AppointmentsComponent },
  { path: 'admin', component: AdminComponent },
  { path: 'appointment-form', component: AppointmentFormComponent },
  { path: 'doctors-new', component: DoctorFormComponent },
  { path: 'doctor/edit/:id', component: DoctorFormComponent },
  { path: 'appointment/edit/:id', component: AppointmentFormComponent }, 

];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

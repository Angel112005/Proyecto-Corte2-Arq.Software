import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DoctorTableComponent } from './components/doctor-table/doctor-table.component';
import { DoctorFormComponent } from './pages/doctor-form/doctor-form.component';
import { DoctorListComponent } from './pages/doctor-list/doctor-list.component';
import { DoctorsRoutingModule } from './doctors-routing.module';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    DoctorTableComponent,
    DoctorFormComponent,
    DoctorListComponent
  ],
  imports: [
    CommonModule,
    DoctorsRoutingModule,
    FormsModule
  ]
})
export class DoctorsModule { }

import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
// import { AppointmentTableComponent } from './components/appointment-table/appointment-table.component';
import { AppointmentFormComponent } from './pages/appointment-form/appointment-form.component';
import { AppointmentListComponent } from './pages/appointment-list/appointment-list.component';
import { AppointmentsRoutingModule } from './appointments-routing.module';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    // AppointmentTableComponent,
    AppointmentFormComponent,
    AppointmentListComponent
  ],
  imports: [
    CommonModule,
    AppointmentsRoutingModule,
    FormsModule
  ]
})
export class AppointmentsModule { }

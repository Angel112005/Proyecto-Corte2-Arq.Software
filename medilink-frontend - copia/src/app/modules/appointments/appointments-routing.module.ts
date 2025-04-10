import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

// Componentes del módulo de citas
import { AppointmentListComponent } from './pages/appointment-list/appointment-list.component';
import { AppointmentFormComponent } from './pages/appointment-form/appointment-form.component';

const routes: Routes = [
  { path: '', component: AppointmentListComponent }, // Lista de citas
  { path: 'new', component: AppointmentFormComponent }, // Nueva cita
  { path: 'edit/:id', component: AppointmentFormComponent } // Editar cita
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AppointmentsRoutingModule { }

import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

// Componentes del módulo de doctores
import { DoctorListComponent } from './pages/doctor-list/doctor-list.component';
import { DoctorFormComponent } from './pages/doctor-form/doctor-form.component';

const routes: Routes = [
  { path: '', component: DoctorListComponent }, // Ver lista de doctores
  { path: 'new', component: DoctorFormComponent }, // Agregar un doctor
  { path: 'edit/:id', component: DoctorFormComponent } // Editar un doctor
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DoctorsRoutingModule { }

import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

// Componentes de la página principal
import { HomeComponent } from './home/home.component';

const routes: Routes = [
  { path: '', component: HomeComponent }, // Página de inicio
  { path: 'admin', loadChildren: () => import('./modules/admin/admin.module').then(m => m.AdminModule) }, // Admin
  { path: 'doctors', loadChildren: () => import('./modules/doctors/doctors.module').then(m => m.DoctorsModule) }, // Doctores
  { path: 'appointments', loadChildren: () => import('./modules/appointments/appointments.module').then(m => m.AppointmentsModule) }, // Citas
  { path: '**', redirectTo: '' } // Redirigir cualquier ruta desconocida a la home
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
